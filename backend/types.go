package backend

import (
	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
	"github.com/KvalitetsIT/kih-telecare-exporter/repository"
)

type ExportResult struct {
	Success     bool
	Measurement repository.MeasurementExportState
}

// Exporter interface to denote
type Exporter interface {
	ExportMeasurements() ([]ExportResult, error)
	ShouldExport(m measurement.Measurement) bool
	HandleMeasurement(measurement measurement.Measurement, m repository.MeasurementExportState) (ExportResult, int, int, int, error)
	ExportMeasurement(measurement measurement.Measurement, m repository.MeasurementExportState) (ExportResult, error)
	MarkPermanentFailed() error
	CheckHealth() error
}
