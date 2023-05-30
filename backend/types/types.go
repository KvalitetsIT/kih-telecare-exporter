package types

import "github.com/KvalitetsIT/kih-telecare-exporter/repository"

type ExportResult struct {
	Success         bool
	Measurement     repository.MeasurementExportState
	ServiceResponse string
}
