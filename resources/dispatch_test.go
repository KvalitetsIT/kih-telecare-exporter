package resources

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/KvalitetsIT/kih-telecare-exporter/app"
	"github.com/KvalitetsIT/kih-telecare-exporter/backend"
	"github.com/KvalitetsIT/kih-telecare-exporter/internal"
	"github.com/KvalitetsIT/kih-telecare-exporter/internal/testutil"
	othtest "github.com/KvalitetsIT/kih-telecare-exporter/internal/testutil"
	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
	"github.com/KvalitetsIT/kih-telecare-exporter/repository"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var appConfig *app.Config

const sqliteDSN = "file:test-dispatch.db?cache=shared&mode=memory"

func init() {
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
}

type failedRepositoryMock struct{}

func (rp failedRepositoryMock) StartExport() (repository.RunStatus, error) {
	return repository.RunStatus{}, nil
}
func (rp failedRepositoryMock) UpdateExport(lr repository.RunStatus) error { return nil }

// Returns stats. Returns total numbed of measurements, failed messaurements, temporarily failed and rejected measusmrents
func (rp failedRepositoryMock) GetTotals() (int, int, int, int) { return 0, 0, 0, 0 }
func (rp failedRepositoryMock) GetRuns() (time.Time, int, int, int, int) {
	return time.Now(), 0, 0, 0, 0
}
func (rp failedRepositoryMock) FindOrCreateMeasurement(m repository.MeasurementExportState) (repository.MeasurementExportState, error) {
	return repository.MeasurementExportState{}, nil
}
func (rp failedRepositoryMock) UpdateMeasurement(m repository.MeasurementExportState) (repository.MeasurementExportState, error) {
	return repository.MeasurementExportState{}, nil

}
func (rp failedRepositoryMock) FindMeasurement(id string) (repository.MeasurementExportState, error) {
	return repository.MeasurementExportState{}, nil

}
func (rp failedRepositoryMock) FindMeasurements() ([]repository.MeasurementExportState, error) {
	return []repository.MeasurementExportState{}, nil

}
func (rp failedRepositoryMock) FindMeasurementsByStatus(status int) ([]repository.MeasurementExportState, error) {
	return []repository.MeasurementExportState{}, nil
}
func (rp failedRepositoryMock) CheckRepository() error {
	return fmt.Errorf("Its and error")
}
func (rp failedRepositoryMock) Close() error { return nil }

type exportMock struct {
	results  []backend.ExportResult
	mustFail bool
}

func (em exportMock) ShouldExport(m measurement.Measurement) bool {
	return true
}

func (em exportMock) MarkPermanentFailed() error {
	return nil
}
func (em exportMock) CheckHealth() error {
	return nil
}

func (em exportMock) ExportMeasurements() ([]backend.ExportResult, error) {
	if em.mustFail {
		return em.results, fmt.Errorf("Triggered failure - ")
	}

	return em.results, nil
}
func (em exportMock) ExportMeasurement(measurement measurement.Measurement, m repository.MeasurementExportState) (backend.ExportResult, error) {
	return backend.ExportResult{}, nil
}

func (em exportMock) HandleMeasurement(measurement measurement.Measurement, m repository.MeasurementExportState) (backend.ExportResult, int, int, int, error) {
	return backend.ExportResult{}, 0, 0, 0, nil
}

func TestMain(m *testing.M) {
	var err error
	fmt.Println("------- running test ------")

	logger = logrus.New()
	logger.Level = logrus.WarnLevel
	viper.SetDefault("loglevel", "warn")

	appConfig, err = app.InitConfig()
	if err != nil {
		logger.Fatal("Error creating config", err)
	}

	appConfig.Database.Database = "exporter"
	appConfig.Database.Username = "opentele"
	appConfig.Database.Password = "opentele"
	appConfig.Database.Hostname = "localhost"

	code := m.Run()
	os.Exit(code)
}

func setupTestDatabase() (*sql.DB, *sqlx.DB, repository.Repository, error) {
	db, conn, err := othtest.SetupTestSQLDatabase(sqliteDSN)

	if db == nil {
		log.Error("DB is nil!!")
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

	repo, err = repository.InitRepository(appConfig, conn)
	if err != nil {
		log.Fatal("Error connection to database", err)
	}

	return db, conn, repo, nil
}

func TestRootResource(t *testing.T) {
	xprtr := exportMock{}
	api := internal.TestInjectorApi{}

	version := "test version"
	environ := "unit test"

	appConfig.Version = version
	appConfig.Level = "warn"
	appConfig.Environment = environ

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

	router, err := InitRouter(appConfig, repo, api, xprtr)

	if err != nil {
		t.Errorf("Error creating router %v", err)
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Error creating http request %v", err)
	}

	router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Status code should be 200 but is %d", rr.Code)
	}

	var root RootResource

	if err := json.Unmarshal(rr.Body.Bytes(), &root); err != nil {
		t.Errorf("Error unmarshalling node %v", err)
	}

	if len(root.APIVersion) == 0 || version != root.APIVersion {
		t.Error("API version is not set")
	}
	if len(root.Environment) == 0 || environ != root.Environment {
		t.Error("environment is not set correctly")
	}
}

func TestHealthStatusResource(t *testing.T) {
	xprtr := exportMock{}
	api := internal.TestInjectorApi{}
	appConfig, _ := app.InitConfig()

	version := "test version"
	environ := "unit test"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		log.Debug("Server received: ", string(body))

		fmt.Println("Succeding export")
		fmt.Fprintln(w, "Hello, client")

		// w.Header().
	}))
	defer ts.Close()

	appConfig.Version = version
	appConfig.Environment = environ

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
	router, err := InitRouter(appConfig, repo, api, xprtr)

	if err != nil {
		t.Errorf("Error creating router %v", err)
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Errorf("Error creating http request %v", err)
	}

	router.ServeHTTP(rr, req)
	fmt.Println("rr.body", rr.Body.String())

	if rr.Code != 200 {
		t.Errorf("Status code should be 200 but is %d", rr.Code)
	}

	var reply healthCheckResponse

	fmt.Println("B: ", rr.Body.String())
	if err := json.Unmarshal(rr.Body.Bytes(), &reply); err != nil {
		t.Errorf("Error unmarshalling node %v", err)
	}

	if len(reply.Errors) > 0 {
		t.Error("Errors from health check - should not happen")
	}

	if len(reply.APIVersion) == 0 || version != reply.APIVersion {
		t.Error("API version is not set")
	}
}
func TestFailedHealthStatusResource(t *testing.T) {
	xprtr := exportMock{}
	appConfig, _ := app.InitConfig()
	api := internal.TestInjectorApi{}

	version := "test version"
	environ := "unit test"

	appConfig.Version = version
	appConfig.Environment = environ

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
	failRepo := failedRepositoryMock{}

	router, err := InitRouter(appConfig, failRepo, api, xprtr)

	if err != nil {
		t.Errorf("Error creating router %v", err)
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Errorf("Error creating http request %v", err)
	}

	router.ServeHTTP(rr, req)
	fmt.Println("rr.body", rr.Body.String())

	if rr.Code != 503 {
		t.Errorf("Status code should be 503 but is %d", rr.Code)
	}

	var reply healthCheckResponse

	fmt.Println("B: ", rr.Body.String())
	if err := json.Unmarshal(rr.Body.Bytes(), &reply); err != nil {
		t.Errorf("Error unmarshalling node %v", err)
	}

	if len(reply.Errors) == 0 {
		t.Error("No errors found - should be triggered")
	}

}

func TestStatusResource(t *testing.T) {
	xprtr := exportMock{}

	appConfig, _ := app.InitConfig()
	api := internal.TestInjectorApi{}

	version := "test version"
	environ := "unit test"

	appConfig.Version = version
	appConfig.Environment = environ

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
	router, err := InitRouter(appConfig, repo, api, xprtr)

	if err != nil {
		t.Errorf("Error creating router %v", err)
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Errorf("Error creating http request %v", err)
	}

	router.ServeHTTP(rr, req)
	fmt.Println("rr.body", rr.Body.String())

	if rr.Code != 200 {
		t.Errorf("Status code should be 200 but is %d", rr.Code)
	}

	var reply exportOverview

	fmt.Println("B: ", rr.Body.String())
	if err := json.Unmarshal(rr.Body.Bytes(), &reply); err != nil {
		t.Errorf("Error unmarshalling node %v", err)
	}

	// if reply.TotalMeasurements != totals {
	// 	t.Errorf("Totals returned %d expected %d", reply.TotalMeasurements, totals)
	// }
	// if reply.RejectedMeasurements != rejected {
	// 	t.Errorf("Rejected returned %d expected %d", reply.RejectedMeasurements, rejected)
	// }
	// if reply.FailedMeasurements != failed {
	// 	t.Errorf("Failed returned %d expected %d", reply.FailedMeasurements, failed)
	// }
	// if reply.TempFailedMeasurements != tempfailed {
	// 	t.Errorf("Temp failed returned %d expected %d", reply.TempFailedMeasurements, tempfailed)
	// }

}

func TestExportResource(t *testing.T) {
	xprtr := exportMock{}

	// rejected := 13
	// totals := 237
	// failed := 144
	// tempfailed := 13

	// rm.totals = totals
	// rm.rejected = rejected
	// rm.tempfailed = tempfailed
	// rm.failed = failed

	results := []backend.ExportResult{}
	results = append(results, backend.ExportResult{Success: true})
	results = append(results, backend.ExportResult{Success: false})

	xprtr.results = results

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
	version := "test version"
	environ := "unit test"

	appConfig.Version = version
	appConfig.Environment = environ
	api := internal.TestInjectorApi{}

	router, err := InitRouter(appConfig, repo, api, xprtr)

	if err != nil {
		t.Errorf("Error creating router %v", err)
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/export", nil)
	if err != nil {
		t.Errorf("Error creating http request %v", err)
	}

	router.ServeHTTP(rr, req)
	fmt.Println("rr.body", rr.Body.String())

	if rr.Code != 200 {
		t.Errorf("Status code should be 200 but is %d", rr.Code)
	}

	var reply []struct {
		Success     bool
		Measurement struct {
			ID          uuid.UUID `json:"id,omitempty"`
			Measurement string    `json:"measurement,omitempty"`
			Patient     string    `json:"patient,omitempty"`
			Status      string    `json:"status,omitempty"`
			CreatedAt   time.Time `json:"created_at,omitempty"`
			UpdatedAt   time.Time `json:"updated_at,omitempty"`
		}
	}

	if err := json.Unmarshal(rr.Body.Bytes(), &reply); err != nil {
		t.Errorf("Error unmarshalling node %v", err)
	}

	if len(reply) != 2 {
		t.Errorf("Results should be 0 is %d", len(reply))
	}

	if !reply[0].Success {
		t.Errorf("Result should be true is %v", reply[0].Success)
	}

	if reply[1].Success {
		t.Errorf("Result should be false is %v", reply[0].Success)
	}
}

func TestFailingExportResource(t *testing.T) {
	xprtr := exportMock{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		log.Debug("Server received: ", string(body))

		fmt.Println("Succeding export")
		fmt.Fprintln(w, "Hello, client")

		// w.Header().
	}))
	defer ts.Close()

	xprtr.mustFail = true
	appConfig, _ := app.InitConfig()
	version := "test version"
	environ := "unit test"

	appConfig.Version = version
	appConfig.Environment = environ
	api := internal.TestInjectorApi{}

	router, err := InitRouter(appConfig, repo, api, xprtr)

	if err != nil {
		t.Errorf("Error creating router %v", err)
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/export", nil)
	if err != nil {
		t.Errorf("Error creating http request %v", err)
	}

	router.ServeHTTP(rr, req)
	fmt.Println("rr.body", rr.Body.String())

	if rr.Code != 500 {
		t.Errorf("Status code should be 200 but is %d", rr.Code)
	}

	var reply RestResponse

	fmt.Println("B: ", rr.Body.String())
	if err := json.Unmarshal(rr.Body.Bytes(), &reply); err != nil {
		t.Errorf("Error unmarshalling node %v", err)
	}
	if len(reply.StatusText) == 0 {
		t.Errorf("Lengtht of status should be larger than 0 - is 0")
	}

}

func TestDevEnvironmentHandling(t *testing.T) {
	xprtr := exportMock{}
	app, _ := app.InitConfig()

	version := "test version"
	environ := "dev"

	app.Version = version
	app.Environment = environ

	api := internal.TestInjectorApi{}
	router, err := InitRouter(app, repo, api, xprtr)

	if err != nil {
		t.Errorf("Error creating router %v", err)
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Error creating http request %v", err)
	}

	router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Status code should be 200 but is %d", rr.Code)
	}

	var root RootResource

	if err := json.Unmarshal(rr.Body.Bytes(), &root); err != nil {
		t.Errorf("Error unmarshalling node %v", err)
	}

	if len(root.APIVersion) == 0 || version != root.APIVersion {
		t.Error("API version is not set")
	}
	if len(root.Environment) == 0 || environ != root.Environment {
		t.Error("environment is not set correctly")
	}

	if !strings.Contains(root.Links.Export, "localhost") {
		t.Errorf("Localhost should be used as host - found %v", root)
	}
}

func TestFailedMeasurementsEndpoint(t *testing.T) {
	xprtr := exportMock{}

	// rejected := 13
	// totals := 237
	// failed := 144
	// tempfailed := 13

	// rm.totals = totals
	// rm.rejected = rejected
	// rm.tempfailed = tempfailed
	// rm.failed = failed

	results := []backend.ExportResult{}
	results = append(results, backend.ExportResult{Success: true})
	results = append(results, backend.ExportResult{Success: false})

	xprtr.results = results

	app, _ := app.InitConfig()
	app.Logger = log
	version := "test version"
	environ := "unit test"

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

	app.Version = version
	app.Environment = environ

	// for i := 1; i <= 20; i++ {
	// 	m := repository.MeasurementType{}
	// 	m.Measurement = fmt.Sprintf("http://clinician/measurement/%d", i)
	// 	m.Status = repository.TEMP_FAILURE
	// 	m.CreatedAt = time.Now().Add(time.Duration(-i) * time.Hour * 24)
	// 	m.UpdatedAt = time.Now()
	// 	repo.FindOrCreateMeasurement(m)
	// }

	mApi, err := testutil.InitMockApi()
	if err != nil {
		t.Errorf("Error creating router %v", err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		log.Debug("Server received: ", string(body))

		fmt.Println("Succeding export")
		fmt.Fprintln(w, "Hello, client")

		// w.Header().
	}))
	defer ts.Close()

	viper.SetDefault("loglevel", "warn")

	app.Export.Backend = "oioxds"
	app.Export.OIOXDSExport.XdsGenerator.URL = ts.URL

	exportr, err := backend.InitExporter(app, mApi, repo)
	if err != nil {
		t.Errorf("Error creating exporter %v", err)
	}

	// set up failed items
	tofail, _ := mApi.FetchMeasurements(time.Now().AddDate(-1, 0, 0), 400)
	setFailed := 0
	for i, v := range tofail.Results {
		if exportr.ShouldExport(v) {
			m := repository.MeasurementExportState{}
			setFailed++
			m.Measurement = v.Links.Measurement
			m.Patient = v.Links.Patient
			m.Status = repository.TEMP_FAILURE
			m.CreatedAt.Time = time.Now().Add(time.Duration(-i) * time.Hour * 24)
			m.UpdatedAt.Time = time.Now()
			var err error
			if m, err = repo.FindOrCreateMeasurement(m); err != nil {
				t.Errorf("Error setting up test %v", err)
			}
		}
	}

	api := internal.TestInjectorApi{}
	router, err := InitRouter(app, repo, api, exportr)

	if err != nil {
		t.Errorf("Error creating router %v", err)
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/failed", nil)
	if err != nil {
		t.Errorf("Error creating http request %v", err)
	}

	router.ServeHTTP(rr, req)
	var failedResults []struct {
		Success     bool
		Measurement struct {
			ID          uuid.UUID `json:"id,omitempty"`
			Measurement string    `json:"measurement,omitempty"`
			Patient     string    `json:"patient,omitempty"`
			Status      string    `json:"status,omitempty"`
			CreatedAt   time.Time `json:"created_at,omitempty"`
			UpdatedAt   time.Time `json:"updated_at,omitempty"`
		}
	}

	if err := json.Unmarshal(rr.Body.Bytes(), &failedResults); err != nil {
		t.Errorf("Error unmarshalling results")
	}

	if len(failedResults) != setFailed {
		t.Errorf("Error - got %d - expected %d", len(failedResults), setFailed)
	}

	for _, v := range failedResults {
		log.Info(reflect.ValueOf(v.Measurement))
		if !v.Success {
			t.Errorf("Failed measurements - expected successful export")
		}

		iStatus := v.Measurement.Status
		if v.Measurement.Status != repository.StatusToText(repository.COMPLETED) {
			t.Errorf("Messaurement is in in-correct state - %s", iStatus)
		}
	}
}
