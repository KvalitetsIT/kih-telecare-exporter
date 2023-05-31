package cmd

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/KvalitetsIT/kih-telecare-exporter/app"
	"github.com/KvalitetsIT/kih-telecare-exporter/backend"
	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
	"github.com/KvalitetsIT/kih-telecare-exporter/repository"
	"github.com/KvalitetsIT/kih-telecare-exporter/resources"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)
	// Default for when reports is to be started from
	viper.SetDefault("PORT", 8360)
	viper.SetDefault("PROXY_SCHEME", "http")
	viper.SetDefault("PROXY_HOST", "localhost")
	viper.SetDefault("PROXY_PORT", 443)
	viper.SetDefault("OT_ENV", "production")
	viper.SetDefault("export.retrydays", 7)
	viper.SetDefault("clinician.batchsize", 1000)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the KIH Export web server",
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
		repo, err := repository.InitRepository(application, conn)
		defer func() {
			repo.Close()
		}()

		if err != nil {
			log.Fatal("Error initializing exporter ", err)
		}

		exprtr, err := backend.InitExporter(application, api, repo)
		if err != nil {
			log.Fatal("Error creating exporter", err)
		}

		r, _ := resources.InitRouter(application, repo, api, exprtr)
		// Start router
		log.Info("starting http endpoint")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", application.Port), r); err != nil {
			log.Fatal("Error creating HTTP endpoint ", err)
		}
	},
}
