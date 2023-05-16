package app

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Get log level. Handles package logging or default loglevel
func (c Config) GetLoggerLevel(pkg string) string {
	lvl, ok := c.logging[pkg]
	if ok {
		logrus.Debug("Returning - ", lvl)
		return lvl
	} else {
		logrus.Debug("Returning - ", c.Level)
		return c.Level
	}
}

// Helper function for DB connections
func (c Config) CreateDatabaseURL() (string, error) {
	if len(c.Database.Hostname) == 0 ||
		len(c.Database.Username) == 0 ||
		len(c.Database.Password) == 0 ||
		len(c.Database.Database) == 0 {
		return "", fmt.Errorf("Database parameters is missing")
	}
	var dbURL strings.Builder

	if c.Database.Port == 3306 || c.Database.Port == 0 {
		c.Logger.Debug("Assuming standard DB port")
		dbURL.WriteString(fmt.Sprintf("%s:%s@tcp(%s)/%s", c.Database.Username, c.Database.Password, c.Database.Hostname, c.Database.Database))
	} else {
		dbURL.WriteString(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Database.Username, c.Database.Password, c.Database.Hostname, c.Database.Port, c.Database.Database))
	}

	// Append location
	dbURL.WriteString("?charset=utf8&parseTime=True")
	// Multistatementsd
	dbURL.WriteString("&multiStatements=true")
	dbURL.WriteString("&loc=")
	dbURL.WriteString(url.PathEscape(c.Location))

	return dbURL.String(), nil
}

func GetPackage(input string) string {
	pkg := strings.ReplaceAll(input, "bitbucket.org/opentelehealth/exporter/", "")
	logrus.Debug("Returning ", pkg)
	return pkg
}

func NewLogger(lvl string) *logrus.Logger {
	logger := logrus.New()
	//logger.SetReportCaller(true)
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		logrus.Fatal("Error parsing level ", err)
	}
	logger.Level = level
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})
	return logger

}

// Read in config
func InitConfig() (*Config, error) {

	var config Config
	config.Level = "info" // Set default level to help tests
	mapstructure.Decode(viper.AllSettings(), &config)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&config); err != nil {
		logrus.Errorf("unable to decode into struct, %v", err)
	}

	logMap := viper.GetStringMapString("logging")
	logConfig := make(map[string]string)

	for k, v := range logMap {
		logConfig[k] = v
	}

	config.logging = logConfig
	logger := NewLogger(config.GetLoggerLevel("app"))

	config.Logger = logger

	return &config, nil
}
