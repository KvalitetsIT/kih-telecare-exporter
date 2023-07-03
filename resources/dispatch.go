package resources

import (
	"fmt"
	"net/http"
	"time"

	"github.com/KvalitetsIT/kih-telecare-exporter/app"
	"github.com/KvalitetsIT/kih-telecare-exporter/backend"
	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
	"github.com/KvalitetsIT/kih-telecare-exporter/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

//var repoConfig *repository.Config

var logger *logrus.Logger
var config *app.Config
var exprtr backend.Exporter
var repo repository.Repository
var api measurement.MeasurementApi

type RootResource struct {
	APIVersion  string `json:"apiVersion"`
	Environment string `json:"environment"`
	Links       struct {
		Apidoc      string `json:"apidoc,omitempty"`
		Categories  string `json:"categories,omitempty"`
		Measurement string `json:"measurement,omitempty"`
		Contents    string `json:"contents,omitempty"`
		Export      string `json:"export,omitempty"`
		Failed      string `json:"failed,omitempty"`
		Health      string `json:"health,omitempty"`
		Status      string `json:"status,omitempty"`
		Schemas     string `json:"schemas,omitempty"`
		Self        string `json:"self"`
	} `json:"links"`
}

var rootResource RootResource

func rootHandler(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, rootResource)
}

type RestResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *RestResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

type healthCheckError struct {
	Resource string
	Error    string
}

type exportOverview struct {
	Measurements struct {
		TotalMeasurements      int
		TempFailedMeasurements int
		RejectedMeasurements   int
		FailedMeasurements     int
	}
	LastRun struct {
		TimeStamp string
		Status    string
	}
	Runs struct {
		Total       int
		Failed      int
		Successfull int
	}
	Source struct {
		Endpoint string
		LastSuccesfullPing string
	}
	Destination struct {
		Type     string
		Endpoint string
	}
}

func setupRootResource() RootResource {
	root := RootResource{}
	var host string
	if "dev" == config.Environment {
		host = "http://localhost:8360"
	} else {
		host = fmt.Sprintf("%s://%s", config.Proxy.Scheme, config.Proxy.Host)
	}

	root.APIVersion = config.Version
	root.Environment = config.Environment
	//root.Links.Apidoc = fmt.Sprintf("%s/apidoc/", host)
	root.Links.Health = fmt.Sprintf("%s/health", host)
	root.Links.Measurement = fmt.Sprintf("%s/measurement", host)
	root.Links.Self = fmt.Sprintf("%s/", host)
	root.Links.Export = fmt.Sprintf("%s/export", host)
	root.Links.Failed = fmt.Sprintf("%s/failed", host)
	root.Links.Status = fmt.Sprintf("%s/status", host)
	return root
}

func InitRouter(appConfig *app.Config, rp repository.Repository, ap measurement.MeasurementApi, ex backend.Exporter) (*chi.Mux, error) {
	config = appConfig
	logger = config.Logger
	exprtr = ex
	repo = rp
	api = ap
	// this is

	//var err error
	// if repoConfig, err = repository.InitConfig(appConfig); err != nil {
	// 	logger.Error("Error creating repository config", err)
	// 	os.Exit(1)
	// }

	rootResource = setupRootResource()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	// r.Use(repoCheck)

	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", rootHandler)
	r.Get("/health", healthcheckHandler)
	r.Get("/status", statusHandler)
	r.Get("/export", exportHandler)
	r.Get("/failed", failedHandler)
	r.Get("/measurement/{measurement}", measurementHandler)

	return r, nil
}
