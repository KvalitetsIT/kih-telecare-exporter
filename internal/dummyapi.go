package internal

import (
	"time"

	"bitbucket.org/opentelehealth/exporter/measurement"
)

type TestInjectorApi struct {
	Patient measurement.PatientResult
}

func (r TestInjectorApi) FetchMeasurements(since time.Time, offset int) (measurement.MeasurementResponse, error) {
	return measurement.MeasurementResponse{}, nil
}
func (r TestInjectorApi) FetchMeasurement(m string) (measurement.Measurement, error) {
	return measurement.Measurement{}, nil
}
func (r TestInjectorApi) FetchPatient(person string) (measurement.PatientResult, error) {
	return r.Patient, nil
}
