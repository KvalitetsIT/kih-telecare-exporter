package testutil

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func SetupTestSQLDatabase(dsn string) (*sql.DB, *sqlx.DB, error) {
	var db *sql.DB
	var conn *sqlx.DB
	var err error

	fmt.Println("Setting up SQLLite test DB")
	db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return db, conn, errors.Wrap(err, "Error opening SQLite DB")

	}

	conn = sqlx.NewDb(db, "sqlite3")

	fmt.Println("Setup SQL Lite")
	return db, conn, nil
}

func PrepareDatabase(db *sql.DB) error {
	createQueryMeasurements := `
    DROP TABLE IF EXISTS measurements;
    CREATE TABLE IF NOT EXISTS measurements (
  id TEXT UNIQUE NOT NULL,
  measurement TEXT UNIQUE NOT NULL PRIMARY KEY,
  patient TEXT,
  status int,
  backend_status int,
  backend_reply text,
  created_at datetime,
  updated_at datetime);`

	if db == nil {
		return fmt.Errorf("DB instance is nil")
	}

	_, err := db.Exec(createQueryMeasurements)
	if err != nil {
		return errors.Wrap(err, "Error bootstrapping db / measurements")
	}

	createQueryRUnstatus := `
DROP TABLE IF EXISTS runstatus;
CREATE TABLE IF NOT EXISTS runstatus (
  id text UNIQUE NOT NULL PRIMARY KEY,
  lastrun datetime,
  status int,
  created_at datetime,
  updated_at datetime);`

	_, err = db.Exec(createQueryRUnstatus)
	if err != nil {
		return errors.Wrap(err, "Error bootstrapping db / runstatus")
	}

	return nil
}

func MeasurementFromFile(f string) (measurement.Measurement, error) {
	var mm measurement.Measurement
	weight_file, err := ioutil.ReadFile(f)
	if err != nil {
		return mm, errors.Wrap(err, "Error reading input file")
	}

	if err := json.Unmarshal(weight_file, &mm); err != nil {
		return mm, errors.Wrap(err, "Error reading input file")
	}

	return mm, nil
}

func MeasurementsFromFile(f string) (measurement.MeasurementResponse, error) {
	var response measurement.MeasurementResponse
	weight_file, err := ioutil.ReadFile(f)
	if err != nil {
		return response, errors.Wrap(err, "Error reading input file")
	}

	if err := json.Unmarshal(weight_file, &response); err != nil {
		return response, errors.Wrap(err, "Error reading input file")
	}

	return response, nil
}

type mockApi struct {
	measurements  measurement.MeasurementResponse
	masurementMap map[string]int
}

// CheckHealth implements measurement.MeasurementApi
func (mockApi) CheckHealth() error {
	return nil
}

func (ma mockApi) FetchMeasurements(since time.Time, offset int) (measurement.MeasurementResponse, error) {
	return ma.measurements, nil
}

func (ma mockApi) FetchMeasurement(mea string) (measurement.Measurement, error) {
	fmt.Println("Fetching measurement - ", mea)
	index, ok := ma.masurementMap[mea]
	fmt.Println("Got ", index)

	if !ok {
		fmt.Println("Did not find measurement")
		return measurement.Measurement{}, fmt.Errorf("No measurement found for %s", mea)
	}

	return ma.measurements.Results[index], nil
}

func InitMockApi() (measurement.MeasurementApi, error) {
	ma := mockApi{}

	mm, err := MeasurementsFromFile("../backend/kih/testdata/server_response.json")
	if err != nil {
		dir, _ := os.Getwd()
		return ma, errors.Wrap(err, fmt.Sprintf("Error getting test data - %s", dir))
	}
	fmt.Println("# of measuremnts ", len(mm.Results))
	ma.measurements = mm
	ma.masurementMap = make(map[string]int)

	for k, v := range ma.measurements.Results {
		ma.masurementMap[v.Links.Measurement] = k
	}

	return ma, nil
}

func (ma mockApi) FetchPatient(person string) (measurement.PatientResult, error) {
	fmt.Println("Person: ", person)
	var filename string
	var patient measurement.PatientResult

	switch person {
	case "http://clinician:8080/clinician/api/patients/13":
		filename = "person_13.json"
	case "http://clinician:8080/clinician/api/patients/14":
		filename = "person_14.json"
	default:
		return patient, fmt.Errorf("Person not found")
	}

	person_file, err := ioutil.ReadFile(fmt.Sprintf("../backend/testdata/%s", filename))
	if err != nil {
		return patient, errors.Wrap(err, "Error reading input file")
	}

	if err := json.Unmarshal(person_file, &patient); err != nil {
		return patient, errors.Wrap(err, "Error reading input file")
	}

	return patient, nil
}
