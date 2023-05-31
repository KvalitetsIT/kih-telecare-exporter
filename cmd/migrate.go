package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/KvalitetsIT/kih-telecare-exporter/app"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Perform database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting migrations")
		application, err := app.InitConfig()
		if err != nil {
			logrus.Fatal("Error initializing exporter ", err)
		}

		log = application.Logger
		log.Info("starting")
		dburl, err := application.CreateDatabaseURL()
		if err != nil {
			log.Fatalln("Error retrieving database credentials", err)
		}

		db, err := sql.Open("mysql", dburl)
		if err != nil {
			log.Fatal("Error connection to database ", err)
		}
		if err := db.Ping(); err != nil {
			log.Fatal("Error communicating with database ", err)
		} else {
			log.Debug("DB pinged successfully")
		}

		driver, _ := mysql.WithInstance(db, &mysql.Config{})

		log.Debug("D", driver)
		m, err := migrate.NewWithDatabaseInstance(
			"file://migrations",
			"mysql",
			driver,
		)
		if err != nil {
			log.Fatal("Error getting migrations", err)
		}

		if len(args) == 1 && args[0] == "drop" {
			log.Info("Dropping database")
			if err := m.Drop(); err != nil {
				log.Fatal("Error dropping database ", err)
			}
			os.Exit(0)
		}
		if len(args) == 1 && args[0] == "down" {
			log.Info("Rolling back database")
			if err := m.Down(); err != nil {
				log.Fatal("Error dropping database ", err)
			}
			os.Exit(0)
		}

		log.Debug("M", m)
		if err := m.Up(); err != nil && err.Error() != "no change" {
			log.Fatal("error migrating database", err)
		} else {
			log.Info("Migrated database")
		}

	},
}
