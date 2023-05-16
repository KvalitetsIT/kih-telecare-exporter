/// Copyright OpenTeleHealth A/s (C) 2019-2020

package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "exporter",
	Short: "OTH KIH export application",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Usage(); err != nil {
			logrus.Fatalf("Error executing commmand %v", err)
		}
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type empty struct{}

var configFile string

//var appConfig *app.Config
var log *logrus.Logger

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "exporter", "", "config file (default is exporter.yaml)")
}

func initConfig() {
	viper.SetConfigName("exporter")
	viper.AddConfigPath("/app")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.Debug(fmt.Sprintf("unable to read config: %v", err))
	}

	// Default for when reports is to be started from
	viper.SetDefault("PORT", 8360)
	viper.SetDefault("PROXY_SCHEME", "http")
	viper.SetDefault("loglevel", "warn")
	viper.SetDefault("PROXY_HOST", "localhost")
	viper.SetDefault("PROXY_PORT", 443)
	viper.SetDefault("OT_ENV", "production")
	viper.SetDefault("EXPORT_TYPE", "kih")
	viper.SetDefault("location", "Europe/Copenhagen")
	viper.SetDefault("export.kih.version", 1)
	viper.SetDefault("export.retrydays", 15)
	viper.SetDefault("export.start", "2019-06-01")
}
