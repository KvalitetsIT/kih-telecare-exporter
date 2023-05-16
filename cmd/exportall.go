package cmd

import (
	"fmt"
	"reflect"
	"time"

	"bitbucket.org/opentelehealth/exporter/app"
	"bitbucket.org/opentelehealth/exporter/backend"
	"bitbucket.org/opentelehealth/exporter/measurement"
	"bitbucket.org/opentelehealth/exporter/repository"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(exportCmd)
	// Default for when reports is to be started from
	viper.SetDefault("clinician.batchsize", 100)
}

var exportCmd = &cobra.Command{
	Use:   "exportall",
	Short: "Starts export of all old measurements",
	Run: func(cmd *cobra.Command, args []string) {
		application, err := app.InitConfig()
		if err != nil {
			logrus.Fatal("Error initializing exporter ", err)
		}

		pkg := app.GetPackage(reflect.TypeOf(empty{}).PkgPath())
		log = app.NewLogger(application.GetLoggerLevel(pkg))

		dbstr, err := application.CreateDatabaseURL()
		log.Debug("DB ", dbstr)
		if err != nil {
			log.Fatal("Error parsing db url: ", err)
		}

		conn, err := sqlx.Open("mysql", dbstr)
		if err != nil {
			panic(err)
		}
		conn.SetMaxOpenConns(10)

		api, _ := measurement.InitMeasurementApi(application)
		log.Warn("Measurements API ", api)
		repo, err := repository.InitRepository(application, conn)
		defer func() { repo.Close() }()
		if err != nil {
			log.Fatal("Error initializing exporter ", err)
		}

		exprtr, err := backend.InitExporter(application, api, repo)
		if err != nil {
			log.Fatal("Error creating exporter", err)
		}

		res, err := reExportAll(application, api, repo, exprtr)
		if err != nil {
			log.Fatal("Error running exportall", err)
		}
		log.Info("Exported - ", len(res), " measurements")

	},
}

func reExportAll(application *app.Config, api measurement.MeasurementApi, repo repository.Repository, e backend.Exporter) ([]backend.ExportResult, error) {
	start := time.Now()
	var err error
	startTime := repository.RunStatus{}
	startTime.Lastrun, err = time.Parse("2006-01-02", application.Export.StartDate)
	if err != nil {
		log.Fatal("Error parsing time", err)
	}

	log.Info("Using start time:", startTime.Lastrun.Format(time.RFC3339))
	exports := []backend.ExportResult{}

	rejected := 0
	exported := 0
	failed := 0
	iteration := 0
	handled := 0

	log.Info("Start time: ", application.Export.StartDate)

	res := measurement.MeasurementResponse{}
	res.Total = application.ClinicianConfig.BatchSize + 1 // make sure we at least run onces

	// Handle pagination
	for i := 0; res.Offset+application.ClinicianConfig.BatchSize < res.Total; i++ {
		log.Debug("Off", res.Offset, " batch", application.ClinicianConfig.BatchSize, " total ", res.Total)
		res, err = api.FetchMeasurements(startTime.Lastrun, i*application.ClinicianConfig.BatchSize)
		if err != nil {
			return exports, err
		}
		for _, measurement := range res.Results {
			m := backend.MeasurementToMeasurementType(measurement)

			m, err := repo.FindOrCreateMeasurement(m)
			if err != nil {
				log.Errorf("Error searching measurements - %+v", err)
				log.Debugf("Trace %+v", err)

				return exports, errors.Wrap(err, "Error getting measurement from DB ")
			}

			switch m.Status {
			case repository.COMPLETED:
				log.Info("M, ", m, " is already completed")
				handled++
				continue
			case repository.NO_EXPORT:
				log.Info("M, ", m, " is flagged as no-export")
				handled++
				continue
			case repository.FAILED:
				log.Debug("M, ", m, " is already flaggged failed")
				handled++
				continue
			default:
				export, ex, fai, re, _ := e.HandleMeasurement(measurement, m)
				exports = append(exports, export)
				rejected += re
				exported += ex
				failed += fai

			}
		}
		iteration++
	}

	if failed > 0 {
		startTime.Status = repository.FAILED
	} else {
		startTime.Status = repository.COMPLETED
	}

	startTime.CreatedAt.Scan(start)
	startTime.Lastrun = start
	startTime.UpdatedAt.Scan(time.Now())

	run, _ := repo.StartExport()

	run.CreatedAt.Scan(start)
	run.Lastrun = start
	run.UpdatedAt.Scan(time.Now())
	run.Status = repository.COMPLETED
	if err := repo.UpdateExport(run); err != nil {
		log.Println("Error ", err)
	}

	fmt.Println("Storted run", run)
	log.Info(
		fmt.Sprintf("type=exportall uuid=%s completed=%s starttime=%s iterations=%d tt=%d total=%d exported=%d rejected=%d failed=%d wasexported=%d",
			startTime.Id.String(), time.Now().Format(time.RFC3339),
			startTime.Lastrun.Format(time.RFC3339),
			iteration, time.Since(start).Milliseconds(),
			exported+failed+rejected, exported, rejected, failed, handled))

	return exports, nil
}
