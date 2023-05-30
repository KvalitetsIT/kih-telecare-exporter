package measurement

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/KvalitetsIT/kih-telecare-exporter/app"
	"github.com/sirupsen/logrus"
)

// MeasurementApi interface incapsulates exporters requirements against OTH measurements services
type MeasurementApi interface {
	// FetchMeasurements takes a timestamp from with the retrieve measurements
	FetchMeasurements(since time.Time, offset int) (MeasurementResponse, error)
	FetchMeasurement(measurement string) (Measurement, error)
	FetchPatient(person string) (PatientResult, error)
}

var (
	config *app.Config
	log    *logrus.Logger
	client http.Client
	token  string
)

func InitMeasurementApi(appConfig *app.Config) (MeasurementApi, error) {
	pkg := app.GetPackage(reflect.TypeOf(Measurement{}).PkgPath())
	log = app.NewLogger(appConfig.GetLoggerLevel(pkg))

	config = appConfig
	client = http.Client{}

	tokenString := fmt.Sprintf("%s:%s", config.Authentication.Key, config.Authentication.Secret)
	token = base64.StdEncoding.EncodeToString([]byte(tokenString))

	log.Debug(fmt.Sprintf("Setting up clinician API for %s", config.ClinicianConfig.URL))

	var api MeasurementApi

	impl := clinicianApi{}
	impl.batchSize = config.ClinicianConfig.BatchSize
	impl.key = config.Authentication.Key
	impl.secret = config.Authentication.Secret
	impl.apiUrl = config.ClinicianConfig.URL

	api = impl
	return api, nil
}
