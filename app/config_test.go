package app

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func TestGetDatabaseURL(t *testing.T) {
	app := Config{}

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	app.Logger = log
	app.Location = "Europe/Copenhagen"
	db := DatabaseConfig{
		Hostname: "localhost",
		Username: "test",
		Password: "test",
		Database: "db",
	}
	app.Database = db

	fmt.Printf("C %+v", app)

	dburl, err := app.CreateDatabaseURL()
	if err != nil {
		t.Errorf("Caught error %v", err)
	}

	fmt.Println("DB ", dburl)
	if !strings.Contains(dburl, "%2F") {
		t.Errorf("Location  not properly escaped - url is %v", dburl)
	}
}

func TestSetPortDatabaseURL(t *testing.T) {
	app := Config{}

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	app.Logger = log
	app.Location = "Europe/Copenhagen"
	db := DatabaseConfig{
		Hostname: "localhost",
		Username: "test",
		Password: "test",
		Database: "db",
		Port:     "3307",
	}
	app.Database = db

	fmt.Printf("C %+v", app)

	dburl, err := app.CreateDatabaseURL()
	if err != nil {
		t.Errorf("Caught error %v", err)
	}

	fmt.Println("DB ", dburl)
	if !strings.Contains(dburl, "%2F") {
		t.Errorf("Location  not properly escaped - url is %v", dburl)
	}
	if !strings.Contains(dburl, ":3307)/") {
		t.Errorf("Port not properly configured - url is %v", dburl)
	}
}

func TestPackageBasedLoggers(t *testing.T) {
	logMap := viper.GetStringMap("logging")

	for k, v := range logMap {
		fmt.Println("k", k, " v ", v)
	}
}
