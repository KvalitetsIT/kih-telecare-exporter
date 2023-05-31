package repository

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/KvalitetsIT/kih-telecare-exporter/app"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const sqliteDSN = "file:repositorytest.db?cache=shared&mode=memory"

var application *app.Config

func setupTestDatabase() (*sql.DB, *sqlx.DB, Repository, error) {
	// db, conn, err := testutil.SetupTestSQLDatabase(sqliteDSN)

	var repo Repository
	var conn *sqlx.DB
	db, err := sql.Open("sqlite3", sqliteDSN)
	if err != nil {
		log.Fatal(err)
	}

	if db == nil {
		log.Fatal("DB is nil!!")
	} else {
		log.Info("DB IS NOT NIL")
	}

	// defer db.Close()
	// defer conn.Close()

	// if err := testutil.PrepareDatabase(db); err != nil {
	// 	log.Fatalf("Error preparing database - %v", err)
	// }

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
		return db, conn, repo, fmt.Errorf("DB instance is nil")
	}

	_, err = db.Exec(createQueryMeasurements)
	if err != nil {
		return db, conn, repo, errors.Wrap(err, "Error bootstrapping db / measurements")
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
		return db, conn, repo, errors.Wrap(err, "Error bootstrapping db / runstatus")
	}

	conn = sqlx.NewDb(db, "mysql")

	repo, err = InitRepository(application, conn)
	if err != nil {
		log.Fatal("Error connection to database", err)
	}

	return db, conn, repo, nil
}

func TestMain(m *testing.M) {
	var err error
	fmt.Println("------- running test ------")
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
	viper.SetDefault("loglevel", "debug")
	viper.SetDefault("export.start", "2019-06-01")
	application, err = app.InitConfig()
	if err != nil {
		log.Fatal("Error creating config", err)
	}

	log.Level = logrus.DebugLevel
	application.Database.Database = "exporter"
	application.Database.Username = "opentele"
	application.Database.Password = "opentele"
	application.Database.Hostname = "localhost"

	application.Export.Backend = "oioxds"
	application.ClinicianConfig.BatchSize = 400

	// os.Remove("./foo.db")

	//sqliteDSN := "file:test.db"

	fmt.Println("Setup repo")
	code := m.Run()
	os.Exit(code)
}

func TestStatusToText(t *testing.T) {

	var tests = []struct {
		status   int
		expected string
	}{
		{0, "INITIAL"},
		{1, "CREATED"},
		{2, "COMPLETED"},
		{3, "TEMP_FAILURE"},
		{4, "FAILED"},
		{5, "NO_EXPORT"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Test status %d", tt.status), func(t *testing.T) {
			if tt.expected != StatusToText(tt.status) {
				t.Errorf("got %q, want %q", StatusToText(tt.status), tt.expected)
			}
		})
	}
}

func TestFindOrCreateMeasurement(t *testing.T) {
	db, conn, repo, err := setupTestDatabase()
	if err != nil {
		t.Errorf("Error setting up db %+v", err)
	}
	defer func() {
		repo.Close()
		conn.Close()
		db.Close()
		fmt.Println("Resources closed")
	}()

	var count int
	res := conn.QueryRow("SELECT count(*) from measurements")
	if err := res.Scan(&count); err != nil {
		t.Errorf("Error reading results - %+v", err)
	}
	fmt.Println("Count ", count)

	if count != 0 {
		t.Errorf("Database count is %d, expect 0", count)
	}
	mes := "/a/measurement"
	m := MeasurementExportState{Measurement: mes}

	m, err = repo.FindOrCreateMeasurement(m)
	if err != nil {
		t.Errorf("Error retrieving measurements %v", err)
	}
	if m.Status != 0 {
		t.Errorf("Wrong status")
	}
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("Errors encountered %v", err)
	// }
	res = conn.QueryRow("SELECT count(*) from measurements")
	if err := res.Scan(&count); err != nil {
		t.Errorf("Error reading results - %+v", err)
	}
	fmt.Println("Count ", count)
	if count != 1 {
		t.Errorf("Database count is %d, expect 0", count)
	}

	if m.ID.URN() == "" {
		t.Errorf("M ID %s is invalid %s", m.ID, m.ID.URN())
	}

	if m.CreatedAt.Time.IsZero() {
		t.Errorf("Created Time stamp not set")
	}
	if m.UpdatedAt.Time.IsZero() {
		t.Errorf("Updated Time stamp not set")
	}
	if m.Status != INITIAL {
		t.Errorf("Status is %s should be %s", StatusToText(m.Status), StatusToText(INITIAL))
	}
}

func TestCheckRepository(t *testing.T) {
	db, conn, repo, err := setupTestDatabase()
	if err != nil {
		t.Errorf("Error setting up db %+v", err)
	}
	defer func() {
		repo.Close()
		conn.Close()
		db.Close()
		fmt.Println("Resources closed")
	}()

	if err := repo.CheckRepository(); err != nil {
		t.Errorf("Check failed %v", err)
	}

	if err := conn.Close(); err != nil {
		t.Errorf("Error closing connection %v", err)
	}
	if err := repo.CheckRepository(); err == nil {
		t.Errorf("Check succeeded - should fail %v", err)
	}

	db, conn, repo, err = setupTestDatabase()
	if err != nil {
		t.Errorf("Error setting up db %+v", err)
	}
	if err := conn.Close(); err != nil {
		t.Errorf("Error closing connection %v", err)
	}
	if err := db.Close(); err != nil {
		t.Errorf("Error closing connection %v", err)
	}

	if err := repo.CheckRepository(); err == nil {
		t.Errorf("Check succeeded - should fail %v", err)
	}

	db, conn, repo, err = setupTestDatabase()
	if err != nil {
		t.Errorf("Error setting up db %+v", err)
	}
	if err := conn.Close(); err != nil {
		t.Errorf("Error closing connection %v", err)
	}
	if err := db.Close(); err != nil {
		t.Errorf("Error closing connection %v", err)
	}
	if err := repo.Close(); err != nil {
		t.Errorf("Error closing connection %v", err)
	}

	if err := repo.CheckRepository(); err == nil {
		t.Errorf("Check succeeded - should fail %v", err)
	}
}

func TestUpdateMeasurement(t *testing.T) {
	db, conn, repo, err := setupTestDatabase()
	if err != nil {
		t.Errorf("Error setting up db %+v", err)
	}
	defer func() {
		repo.Close()
		conn.Close()
		db.Close()
		fmt.Println("Resources closed")
	}()

	var count int
	res := conn.QueryRow("SELECT count(*) from measurements")
	if err := res.Scan(&count); err != nil {
		t.Errorf("Error reading results - %+v", err)
	}
	fmt.Println("Count ", count)

	if count != 0 {
		t.Errorf("Database count is %d, expect 0", count)
	}
	mes := "/a/measurement"
	m := MeasurementExportState{Measurement: mes, Patient: "mypatient"}

	m, err = repo.FindOrCreateMeasurement(m)
	if err != nil {
		t.Errorf("Error retrieving measurements %v", err)
	}
	if m.Status != 0 {
		t.Errorf("Wrong status")
	}
	if m.CreatedAt.Time.IsZero() {
		t.Errorf("Created Time stamp not set")
	}
	if m.UpdatedAt.Time.IsZero() {
		t.Errorf("Updated Time stamp not set")
	}
	if m.Status != INITIAL {
		t.Errorf("Status is %s should be %s", StatusToText(m.Status), StatusToText(INITIAL))
	}

	m.Status = TEMP_FAILURE

	m, err = repo.UpdateMeasurement(m)
	if err != nil {
		t.Errorf("Error updating measurement %v", err)
	}

	m2, err := repo.FindOrCreateMeasurement(MeasurementExportState{Measurement: mes})
	if err != nil {
		t.Errorf("Error retrieving measurements %v", err)
	}
	fmt.Println("Received ", m2, " ", m2.UpdatedAt.Time)
	if m2.Status != m.Status {
		t.Errorf("Status is %s - should be %s", StatusToText(m2.Status), StatusToText(m.Status))

	}

	if m2.ID != m.ID {
		t.Errorf("IDS different - %s <-> %s", m2.ID.String(), m.ID.String())
	}
	if !m2.UpdatedAt.Valid {
		t.Errorf("Updated time stamp not correct %v", m2.UpdatedAt.Valid)
	}

	if m2.UpdatedAt.Time.Format(time.RFC3339) != m.UpdatedAt.Time.Format(time.RFC3339) {
		t.Errorf("Updated time are different - %s<>%s", m2.UpdatedAt.Time.Format(time.RFC3339), m.UpdatedAt.Time.Format(time.RFC3339))
	}
}
