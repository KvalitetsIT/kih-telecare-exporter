package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"bitbucket.org/opentelehealth/exporter/app"
	othtest "bitbucket.org/opentelehealth/exporter/internal/testutil"
	"bitbucket.org/opentelehealth/exporter/measurement"
	"bitbucket.org/opentelehealth/exporter/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var application *app.Config

type mockApi struct {
	measurements measurement.MeasurementResponse
}

func (ma mockApi) FetchMeasurements(since time.Time, offset int) (measurement.MeasurementResponse, error) {
	fmt.Println("OFFE", offset, ma.measurements.Total)
	return ma.measurements, nil
}

func (ma mockApi) FetchMeasurement(mea string) (measurement.Measurement, error) {
	return ma.measurements.Results[0], nil
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

	person_file, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s", filename))
	if err != nil {
		return patient, errors.Wrap(err, "Error reading input file")
	}

	if err := json.Unmarshal(person_file, &patient); err != nil {
		return patient, errors.Wrap(err, "Error reading input file")
	}

	return patient, nil
}

func measurementFromFile(f string) (measurement.Measurement, error) {
	var mm measurement.Measurement
	weight_file, err := ioutil.ReadFile(fmt.Sprintf("kih/testdata/%s", f))
	if err != nil {
		return mm, errors.Wrap(err, "Error reading input file")
	}

	if err := json.Unmarshal(weight_file, &mm); err != nil {
		return mm, errors.Wrap(err, "Error reading input file")
	}

	return mm, nil
}

func measurementsFromFile(f string) (measurement.MeasurementResponse, error) {
	var response measurement.MeasurementResponse
	weight_file, err := ioutil.ReadFile(fmt.Sprintf("kih/testdata/%s", f))
	if err != nil {
		return response, errors.Wrap(err, "Error reading input file")
	}

	if err := json.Unmarshal(weight_file, &response); err != nil {
		return response, errors.Wrap(err, "Error reading input file")
	}

	return response, nil
}

const created_by = "unit testing framework"

const sqliteDSN = "file:test.db?cache=shared&mode=memory"

func TestMain(m *testing.M) {
	var err error
	fmt.Println("------- running test ------")
	log = logrus.New()
	log.SetLevel(logrus.WarnLevel)
	viper.SetDefault("loglevel", "warn")
	viper.SetDefault("export.start", "2019-06-01")
	application, err = app.InitConfig()
	if err != nil {
		log.Fatal("Error creating config", err)
	}

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

func setupTestDatabase() (*sql.DB, *sqlx.DB, repository.Repository, error) {
	db, conn, err := othtest.SetupTestSQLDatabase(sqliteDSN)

	if db == nil {
		log.Fatal("DB is nil!!")
	} else {
		log.Info("DB IS NOT NIL")
	}

	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	// defer conn.Close()

	if err := othtest.PrepareDatabase(db); err != nil {
		log.Fatalf("Error preparing database - %v", err)
	}

	sqlx.NewDb(db, "mysql")

	repo, err = repository.InitRepository(application, conn)
	if err != nil {
		log.Fatal("Error connection to database", err)
	}

	return db, conn, repo, nil
}

// func prepareDatabase() error {
// 	createQueryMeasurements := `
//     DROP TABLE IF EXISTS measurements;
//     CREATE TABLE IF NOT EXISTS measurements (
//   id TEXT UNIQUE NOT NULL,
//   measurement TEXT UNIQUE NOT NULL PRIMARY KEY,
//   status int,
//   backend_status int,
//   backend_reply text,
//   created_at datetime,
//   updated_at datetime);`

// 	res, err := db.Exec(createQueryMeasurements)
// 	if err != nil {
// 		log.Fatalf("Error bootstrapping db - %v", err)
// 	}
// 	log.Debug("Res: ", res)

// 	log.Debug("Creating runstatus")
// 	createQueryRUnstatus := `
// DROP TABLE IF EXISTS runstatus;
// CREATE TABLE IF NOT EXISTS runstatus (
//   id text UNIQUE NOT NULL PRIMARY KEY,
//   lastrun datetime,
//   status int,
//   created_at datetime);`

// 	res, err = db.Exec(createQueryRUnstatus)
// 	if err != nil {
// 		log.Fatalf("Error bootstrapping db - %v", err)
// 	}
// 	log.Debug("Res: ", res)

// 	return nil
// }

func TestExportMeasurements(t *testing.T) {
	tests := []struct {
		name             string
		fails            bool
		measurementsFile string
	}{
		{name: "All successeds", fails: false, measurementsFile: "server_response.json"},
		{name: "All fails", fails: true, measurementsFile: "server_response.json"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Preparing database")
			db, conn, repo, err := setupTestDatabase()
			if err != nil {
				t.Fatal("Error setting up DB")
			}
			defer func() {
				log.Info("Closing db connections")
				repo.Close()
				conn.Close()
				db.Close()
			}()

			fmt.Println("Done setting up db")
			application, err := app.InitConfig()
			if err != nil {
				t.Errorf("error instantiating %+v", err)
			}
			application.Logger = log
			application.Export.Backend = "oioxds"
			mockApi := mockApi{}
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, _ := ioutil.ReadAll(r.Body)
				log.Debug("Server received: ", string(body))

				if tt.fails {
					fmt.Println("Failing export")
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("500 - Something bad happened!")) // nolint
				} else {
					fmt.Println("Succeding export")
					fmt.Fprintln(w, "Hello, client")
				}

				// w.Header().
			}))
			defer ts.Close()
			application.Export.OIOXDSExport.SkipSslVerify = true
			application.Export.OIOXDSExport.XdsGenerator.URL = ts.URL

			mm, err := measurementsFromFile(tt.measurementsFile)
			if err != nil {
				t.Errorf("Error reading measurement from file - %v", err)
			}
			mockApi.measurements = mm
			mockApi.measurements.Total = len(mm.Results)
			application.ClinicianConfig.BatchSize = len(mm.Results) + 400

			// mockRepo := repositoryMock{measurements: make(map[string]repository.MeasurementType)}

			exprtr, err := InitExporter(application, mockApi, repo)
			if err != nil {
				t.Errorf("error instantiating %+v - message: %s", err, err.Error())
			}

			a, err := exprtr.ExportMeasurements()
			if tt.fails {
				for _, v := range a {
					if v.Success {
						t.Errorf("Error - measurment should be flagged as failed  %+v", v)
					}
				}
			} else {
				if err != nil {
					t.Errorf("error exporting measurement %+v", err)
				}
			}

			if len(a) != len(mm.Results) {
				t.Errorf("Number of exports does not match - got: %v, expected: %v", len(a), len(mm.Results))
			}
		})
	}
}

func TestExportMeasurement(t *testing.T) {
	tests := []struct {
		name            string
		filepath        string
		mustFail        bool
		numberOfResults int
		succces         bool
		uipacs          []string
		results         []string
		units           []string
	}{
		{"Weight", "weight.json", false, 1, true, []string{"NPU03804"}, []string{"84.9"}, []string{"kg"}},
		{"Bloodpressure", "blood_pressure.json", false, 2, true, []string{"DNK05472", "DNK05473"}, []string{"130", "80"}, []string{"mmHg"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Preparing database")
			db, conn, repo, err := setupTestDatabase()
			if err != nil {
				t.Fatal("Error setting up DB")
			}
			defer func() {
				log.Info("Closing db connections")
				repo.Close()
				conn.Close()
				db.Close()
			}()

			fmt.Println("Done setting up db")

			application, err := app.InitConfig()
			if err != nil {
				t.Errorf("error instantiating %+v", err)
			}
			application.Logger = log
			application.Export.Backend = "oioxds"
			mockApi := mockApi{}
			// mockRepo := repositoryMock{measurements: measurements}

			var soapRequest string
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, _ := ioutil.ReadAll(r.Body)
				soapRequest = string(body)
				log.Debug("Server received: ", soapRequest)
				if tt.mustFail {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("500 - Something bad happened!")) // nolint
				} else {
					// w.Header().
					fmt.Fprintln(w, "Hello, client")
				}
			}))
			defer ts.Close()

			application.Export.OIOXDSExport.XdsGenerator.URL = ts.URL
			exprtr, err := InitExporter(application, mockApi, repo)
			if err != nil {
				t.Errorf("error instantiating %+v - message: %s", err, err.Error())
			}

			mm, err := measurementFromFile(tt.filepath)
			if err != nil {
				t.Errorf("Error reading measurement from file - %v", err)
			}

			rm := MeasurementToMeasurementType(mm)

			rm, err = repo.FindOrCreateMeasurement(rm)
			if err != nil {
				t.Errorf("Error getting measurement from repository - %v", err)
			}

			// rm := repository.MeasurementType{Measurement: mm.Links.Measurement, Status: 0}
			// rm.ID = uuid.New()
			fmt.Println("MM", mm)
			fmt.Println("RM", rm)

			a, err := exprtr.ExportMeasurement(mm, rm)
			if err != nil {
				t.Errorf("error exporting measurement %+v", err)
			}
			fmt.Println("Got: ", a)

			if !a.Success {
				t.Error("Export failed")
			}

			// Get the update version

			var mmDbRes []repository.MeasurementExportState
			err = conn.Select(&mmDbRes, "SELECT id,measurement,patient,status,created_at,updated_at from measurements where measurement=$1", mm.Links.Measurement)
			if err != nil {
				t.Errorf("Unexpected error %v", err)

			}
			mmDb := mmDbRes[0]
			fmt.Println("I3 ", mmDb)
			if !strings.Contains(soapRequest, mmDb.ID.String()) {
				t.Error("UUID not found in request - looking for ", mmDb.ID.String())
			}

			// res := a.Results.(repository.MeasurementType)
			if mmDb.Status != repository.COMPLETED {
				t.Errorf("Status should be clompeted - but is %s", repository.StatusToText(mmDb.Status))
			}

			//if res.Status != repository.COMPLETED {
			// 	t.Error("Status should be COMPLETE - is ", repository.StatusToText(res.Status))
			// }
			log.Debug(a)
		})
	}
}

func TestHandleTemporaryFailedMeasurements(t *testing.T) {
	fmt.Println("Preparing database")
	db, conn, repo, err := setupTestDatabase()
	if err != nil {
		t.Fatal("Error setting up DB")
	}
	defer func() {
		fmt.Println("Closing db connections")
		repo.Close()
		conn.Close()
		db.Close()
	}()

	mApi, err := othtest.InitMockApi()
	if err != nil {
		t.Errorf("Error creating router %v", err)
	}

	tofail, _ := mApi.FetchMeasurements(time.Now().AddDate(-1, 0, 0), 50)
	setTempFailed := 0
	for i, v := range tofail.Results {
		m := repository.MeasurementExportState{}
		setTempFailed++
		m.Measurement = v.Links.Measurement
		m.Status = repository.TEMP_FAILURE
		m.CreatedAt.Time = time.Now().Add(time.Duration(-i) * time.Hour * 24)
		fmt.Println(i, "Status: ", repository.StatusToText(m.Status), " - C: ", m.CreatedAt.Time, " is Zero? ", m.CreatedAt.Time.IsZero())
		m.UpdatedAt.Time = time.Now()
		m, err = repo.FindOrCreateMeasurement(m)
		if err != nil {
			t.Errorf("Error storing measurements %v", err)
		}
		log.Debug("M: ", m)
	}

	fmt.Println("Done setting up db")

	application, err := app.InitConfig()
	if err != nil {
		t.Errorf("error instantiating %+v", err)
	}
	application.Logger = log
	application.Export.DaysToRetry = 7
	application.Export.Backend = "oioxds"

	mockApi := mockApi{}

	exprtr, err := InitExporter(application, mockApi, repo)
	if err != nil {
		t.Errorf("error instantiating %+v - message: %s", err, err.Error())
	}

	if err := exprtr.MarkPermanentFailed(); err != nil {
		t.Error("Error handling temp failed ", err)
	}

	failed, err := repo.FindMeasurementsByStatus(repository.FAILED)

	if err != nil {
		t.Errorf("Error querying repo %v", err)
	}

	log.Debug("Comparing ", setTempFailed, " ", len(failed))
	if len(failed) != setTempFailed-(application.Export.DaysToRetry+1) { // Substracted 1 day per measuremnet in test set - per day.
		t.Errorf(fmt.Sprintf("Expected %d - got %d ", setTempFailed, len(failed)))

	}
}
