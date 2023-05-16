package app

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// Proxy Config settings matching OTH guidelines
type ProxyConfig struct {
	Scheme string `mapstructure:"scheme"`
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
}

func (p ProxyConfig) String() string {
	return fmt.Sprintf("%s://%s:%d", p.Scheme, p.Host, p.Port)
}

// IdP keys
type Authentication struct {
	Key    string `mapstructure:"key"`
	Secret string `mapstructure:"secret"`
}

// Application Config
type Config struct {
	// Application Logger configured
	Logger          *logrus.Logger
	logging         map[string]string
	Level           string         `mapstructure:"loglevel"`
	Environment     string         `mapstructure:"environment"`
	Location        string         `mapstructure:"location"`
	Version         string         `mapstructure:"version"`
	Port            int            `mapstructure:"port"`
	AppConfig       AppConfig      `mapstructure:"app"`
	Database        DatabaseConfig `mapstructure:"database"`
	Proxy           ProxyConfig    `mapstructure:"proxy"`
	Export          ExportConfig   `mapstructure:"export"`
	ClinicianConfig ClincianConfig `mapstructure:"clinician"`
	Authentication  Authentication `mapstructure:"authentication"`
}

func (c Config) String() string {
	var output strings.Builder

	output.WriteString(fmt.Sprintf("\t- environment: %s\n", c.Environment))
	output.WriteString(fmt.Sprintf("\t- version: %s\n", c.Version))
	output.WriteString(fmt.Sprintf("\t- proxy: %s\n", c.Proxy))
	output.WriteString(fmt.Sprintf("\t- export: %s\n", c.Export))
	output.WriteString(fmt.Sprintf("\t- authenticationKey: %s\n", c.Authentication.Key))

	return output.String()
}

// Own settings
type AppConfig struct {
	URL  string `mapstructure:"url"`
	Port int    `mapstructure:"port"`
}

// configure linan endpoint
type ClincianConfig struct {
	BatchSize int    `mapstructure:"batchsize"`
	URL       string `mapstructure:"url"`
}

// Export backends
type ExportConfig struct {
	StartDate         string       `mapstructure:"start"`
	Backend           string       `mapstructure:"backend"`
	CreatedBy         string       `mapstructure:"created_by"`
	DaysToRetry       int          `mapstructure:"retrydays"`
	NoDeviceWhiteList bool         `mapstructure:"nodevicewhitelist"`
	OIOXDSExport      OIOXDSConfig `mapstructure:"oioxds"`
}

// Returns endpoint depending on configuration
func (e ExportConfig) GetExportEndpoint() string {
	switch e.Backend {
	case "oioxds":
		return e.OIOXDSExport.XdsGenerator.URL
	default:
		return "Unknown"
	}
}

func (e ExportConfig) String() string {
	return fmt.Sprintf("%s - OIOXDS: %s", e.Backend, e.OIOXDSExport)
}

// Setting up Sosi for DGWS
type SosiConfig struct {
	URL             string `mapstructure:"url"`
	DumpSosiRequest bool   `mapstructure:"dumpRequest"`
}

type OIOXDSConfig struct {
	SkipSslVerify bool      `mapstructure:"skipSSLVerify"`
	XdsGenerator  XdsConfig `mapstructure:"xdsgenerator"`
}

func (o OIOXDSConfig) String() string {
	return fmt.Sprintf("XDS Generator: %v", o.XdsGenerator)
}

type XdsConfig struct {
	URL         string `mapstructur:"url"`
	HealthCheck string `mapstructur:"healthcheck"`
}

// Local database
type DatabaseConfig struct {
	Hostname string `mapstructure:"hostname"`
	Username string `mapstructure:"username"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
