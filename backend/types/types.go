package types

import "bitbucket.org/opentelehealth/exporter/repository"

type ExportResult struct {
	Success         bool
	Measurement     repository.MeasurementExportState
	ServiceResponse string
}
