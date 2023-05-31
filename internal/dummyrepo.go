package internal

import (
	"time"

	"github.com/KvalitetsIT/kih-telecare-exporter/repository"
)

// Setups a dummy repo

type DummyRepo struct{}

func (r DummyRepo) StartExport() (repository.RunStatus, error) {
	return repository.RunStatus{}, nil
}
func (r DummyRepo) UpdateExport(lr repository.RunStatus) error {
	return nil
}

func (r DummyRepo) GetTotals() (int, int, int, int) {
	return 0, 0, 0, 0
}
func (r DummyRepo) GetRuns() (time.Time, int, int, int, int) {
	return time.Now(), 0, 0, 0, 0
}
func (r DummyRepo) FindOrCreateMeasurement(m repository.MeasurementExportState) (repository.MeasurementExportState, error) {
	return repository.MeasurementExportState{}, nil
}
func (r DummyRepo) UpdateMeasurement(m repository.MeasurementExportState) (repository.MeasurementExportState, error) {
	return repository.MeasurementExportState{}, nil
}
func (r DummyRepo) FindMeasurement() repository.MeasurementExportState {
	return repository.MeasurementExportState{}
}
func (r DummyRepo) FindMeasurements() ([]repository.MeasurementExportState, error) {
	return []repository.MeasurementExportState{}, nil

}
func (r DummyRepo) FindMeasurementsByStatus(status int) ([]repository.MeasurementExportState, error) {
	return []repository.MeasurementExportState{}, nil

}
func (r DummyRepo) CheckRepository() error { return nil }
func (r DummyRepo) Close() error           { return nil }
