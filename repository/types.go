package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	INITIAL      = 0
	CREATED      = 1
	COMPLETED    = 2
	TEMP_FAILURE = 3
	FAILED       = 4
	NO_EXPORT    = 5
)

func StatusToText(s int) string {

	name := ""
	switch s {
	case INITIAL:
		name = "INITIAL"
	case CREATED:
		name = "CREATED"
	case COMPLETED:
		name = "COMPLETED"
	case TEMP_FAILURE:
		name = "TEMP_FAILURE"
	case FAILED:
		name = "FAILED"
	case NO_EXPORT:
		name = "NO_EXPORT"
	}
	return name
}

type MeasurementExportState struct {
	ID            uuid.UUID      `json:"id"`
	Measurement   string         `json:"measurement"`
	Patient       string         `json:"patient"`
	Status        int            `json:"status"`
	BackendStatus sql.NullInt32  `json:"-" db:"-"`
	BackendValue  sql.NullString `json:"-" db:"-"`
	CreatedAt     sql.NullTime   `json:"created_at" db:"created_at"`
	UpdatedAt     sql.NullTime   `json:"updated_at" db:"updated_at"`
}

func (m MeasurementExportState) String() string {
	return fmt.Sprintf("ID: %s - Status: %s - measurement: %s", m.ID, StatusToText(m.Status), m.Measurement)
}

func (m MeasurementExportState) MarshalJSON() ([]byte, error) {

	values := struct {
		ID            uuid.UUID      `json:"id,omitempty"`
		Measurement   string         `json:"measurement,omitempty"`
		Patient       string         `json:"patient,omitempty"`
		Status        string         `json:"status,omitempty"`
		BackendStatus sql.NullInt32  `json:"-"`
		BackendValue  sql.NullString `json:"-"`
		CreatedAt     time.Time      `json:"created_at,omitempty"`
		UpdatedAt     time.Time      `json:"updated_at,omitempty"`
	}{
		ID:          m.ID,
		Measurement: m.Measurement,
		Patient:     m.Patient,
		Status:      StatusToText(m.Status),
		CreatedAt:   m.CreatedAt.Time,
		UpdatedAt:   m.UpdatedAt.Time,
	}

	return json.Marshal(values)
}

type Repository interface {
	StartExport() (RunStatus, error)
	UpdateExport(lr RunStatus) error
	// Returns stats. Returns total numbed of measurements, failed messaurements, temporarily failed and rejected measusmrents
	GetTotals() (int, int, int, int)
	GetRuns() (time.Time, int, int, int, int)
	FindOrCreateMeasurement(m MeasurementExportState) (MeasurementExportState, error)
	UpdateMeasurement(m MeasurementExportState) (MeasurementExportState, error)
	FindMeasurement(id string) (MeasurementExportState, error)
	FindMeasurements() ([]MeasurementExportState, error)
	FindMeasurementsByStatus(status int) ([]MeasurementExportState, error)
	CheckRepository() error
	Close() error
}

type RunStatus struct {
	Id        uuid.UUID    `db:"id"`
	Lastrun   time.Time    `db:"lastrun"`
	Status    int          `db:"status"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (rs RunStatus) String() string {
	return fmt.Sprintf("[%s] State: %s - created: %s - lastrun: %s", rs.Id.String(), StatusToText(rs.Status), rs.CreatedAt.Time.Format(time.RFC3339), rs.Lastrun.Format(time.RFC3339))
}

type Measurement struct {
}
