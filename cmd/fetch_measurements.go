package cmd

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"time"

	"github.com/KvalitetsIT/kih-telecare-exporter/app"
	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
	"github.com/KvalitetsIT/kih-telecare-exporter/repository"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var fetchOutFile string

func init() {
	fetchMeasurementsCmd.Flags().StringVar(&fetchOutFile, "out-file", "measurements.json", "File to write measurements to")
	rootCmd.AddCommand(fetchMeasurementsCmd)
}

var fetchMeasurementsCmd = &cobra.Command{
	Use:   "fetch-measurements",
	Short: "Fetches measurements from the clinician API and saves them to a file.",
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

		log.Info("Fetching all measurements...")

		var allMeasurements []measurement.Measurement
		res := measurement.MeasurementResponse{}
		res.Total = application.ClinicianConfig.BatchSize + 1 // make sure we at least run once

		// Set start time to 1970 so we get all measurements, not just recent ones.
		fetchTime := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)

		for i := 0; res.Offset+application.ClinicianConfig.BatchSize < res.Total; i++ {
			log.Debug("Offset ", res.Offset, " batch ", application.ClinicianConfig.BatchSize, " total ", res.Total)
			response, err := api.FetchMeasurements(fetchTime, i*application.ClinicianConfig.BatchSize)
			if err != nil {
				log.Fatal("Error fetching measurements", err)
			}
			allMeasurements = append(allMeasurements, response.Results...)
			res = response
		}

		log.Infof("Fetched %d measurements", len(allMeasurements))

		data, err := json.MarshalIndent(allMeasurements, "", "  ")
		if err != nil {
			log.Fatal("Error marshalling measurements", err)
		}

		err = ioutil.WriteFile(fetchOutFile, data, 0644)
		if err != nil {
			log.Fatal("Error writing measurements to file", err)
		}

		log.Infof("Successfully wrote measurements to %s", fetchOutFile)
	},
}
