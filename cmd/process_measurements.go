package cmd

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"time"

	"github.com/KvalitetsIT/kih-telecare-exporter/app"
	"github.com/KvalitetsIT/kih-telecare-exporter/backend"
	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
	"github.com/KvalitetsIT/kih-telecare-exporter/repository"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var processInFile string

func init() {
	processMeasurementsCmd.Flags().StringVar(&processInFile, "in-file", "measurements.json", "File to read measurements from")
	rootCmd.AddCommand(processMeasurementsCmd)
}

var processMeasurementsCmd = &cobra.Command{
	Use:   "process-measurements",
	Short: "Processes measurements from a file and exports them.",
	Run: func(cmd *cobra.Command, args []string) {
		application, err := app.InitConfig()
		if err != nil {
			logrus.Fatal("Error initializing exporter ", err)
		}

		pkg := app.GetPackage(reflect.TypeOf(empty{}).PkgPath())
		log = app.NewLogger(application.GetLoggerLevel(pkg))

		dbstr, err := application.CreateDatabaseURL()
		if err != nil {
			log.Fatal("Error parsing db url: ", err)
		}

		conn, err := sqlx.Open("mysql", dbstr)
		if err != nil {
			panic(err)
		}
		conn.SetMaxOpenConns(10)
		defer conn.Close()

		repo, err := repository.InitRepository(application, conn)
		if err != nil {
			log.Fatal("Error initializing repository ", err)
		}
		defer repo.Close()

		api, err := measurement.InitMeasurementApi(application)
		if err != nil {
			log.Fatal("Error initializing measurement api", err)
		}

		exprtr, err := backend.InitExporter(application, api, repo)
		if err != nil {
			log.Fatal("Error creating exporter", err)
		}

		log.Infof("Reading measurements from %s", processInFile)

		fileContent, err := ioutil.ReadFile(processInFile)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", processInFile, err)
		}

		var measurements []measurement.Measurement
		err = json.Unmarshal(fileContent, &measurements)
		if err != nil {
			log.Fatalf("Error unmarshalling measurements from file: %v", err)
		}

		log.Infof("Processing %d measurements", len(measurements))

		start := time.Now()
		runStatus, err := repo.StartExport()
		if err != nil {
			log.Fatalf("Error starting export in repository: %v", err)
		}

		var exports []backend.ExportResult
		rejected := 0
		exported := 0
		failed := 0

		for _, measurement := range measurements {
			m := backend.MeasurementToMeasurementType(measurement)

			m, err := repo.FindOrCreateMeasurement(m)
			if err != nil {
				log.Errorf("Error searching measurements - %+v", err)
				failed++
				continue
			}

			switch m.Status {
			case repository.COMPLETED:
				log.Debug("M, ", m, " is already completed")
				continue
			case repository.NO_EXPORT:
				log.Debug("M, ", m, " is flagged as no-export")
				continue
			case repository.FAILED:
				log.Debug("M, ", m, " is already flaggged failed")
				continue
			default:
				export, ex, fai, re, _ := exprtr.HandleMeasurement(measurement, m)
				exports = append(exports, export)
				rejected += re
				exported += ex
				failed += fai
			}
		}

		if failed > 0 {
			runStatus.Status = repository.FAILED
		} else {
			runStatus.Status = repository.COMPLETED
		}
		if err := repo.UpdateExport(runStatus); err != nil {
			log.Fatalf("Error updating export status: %v", err)
		}

		log.Infof(
			"type=export-from-file uuid=%s completed=%s starttime=%s tt=%d total=%d exported=%d rejected=%d failed=%d",
			runStatus.Id.String(), time.Now().Format(time.RFC3339),
			runStatus.Lastrun.Format(time.RFC3339),
			time.Since(start).Milliseconds(),
			exported+failed+rejected, exported, rejected, failed)
	},
}
