package resources

import (
	"net/http"
	"time"

	"github.com/KvalitetsIT/kih-telecare-exporter/backend"
	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
	"github.com/KvalitetsIT/kih-telecare-exporter/repository"
	"github.com/go-chi/render"
)

var lastSuccesfullSourcePing time.Time
var lastFailedSourcePing time.Time

var lastSuccesfullDestinatiomPing time.Time
var lastFailedDestinatiomPing time.Time

var lastSuccesfullDBPing time.Time
var lastFailedDBPing time.Time
var serviceStarted = time.Now()

func statusHandler(w http.ResponseWriter, r *http.Request) {
	checkSource()
	checkDestination()
	checkDb()

	overview := exportOverview{}

	total, failed, tempfailed, rejects := repo.GetTotals()
	overview.Measurements.TotalMeasurements = total
	overview.Measurements.TempFailedMeasurements = tempfailed
	overview.Measurements.FailedMeasurements = failed
	overview.Measurements.RejectedMeasurements = rejects

	lasttime, laststatus, totalruns, successfullruns, failedruns := repo.GetRuns()
	overview.LastRun.TimeStamp = lasttime.Format(time.RFC3339)
	overview.LastRun.Status = repository.StatusToText(laststatus)
	overview.Runs.Total = totalruns
	overview.Runs.Failed = failedruns
	overview.Runs.Successfull = successfullruns

	overview.Source.Endpoint = config.ClinicianConfig.URL
	overview.Source.LastSuccesfullPing = lastSuccesfullSourcePing.Format(time.RFC3339)
	overview.Source.LastFailedPing = lastFailedSourcePing.Format(time.RFC3339)
	overview.Destination.Type = config.Export.Backend
	overview.Destination.Endpoint = config.Export.GetExportEndpoint()
	overview.Destination.LastSuccesfullPing = lastSuccesfullDestinatiomPing.Format(time.RFC3339)
	overview.Destination.LastFailedPing = lastFailedDestinatiomPing.Format(time.RFC3339)
	overview.DB.LastSuccesfullPing = lastSuccesfullDBPing.Format(time.RFC3339)
	overview.DB.LastFailedPing = lastFailedDBPing.Format(time.RFC3339)

	overview.Service.Started = serviceStarted.Format(time.RFC3339)
	render.JSON(w, r, overview)
}

func checkSource() {
	if err := api.CheckHealth(); err != nil {
		lastFailedSourcePing = time.Now()
	} else {
		lastSuccesfullSourcePing = time.Now()
	}

}

func checkDb() {
	if err := repo.CheckRepository(); err != nil {
		lastFailedDBPing = time.Now()
	} else {
		lastSuccesfullDBPing = time.Now()
	}
}

func checkDestination() {
	if err := exprtr.CheckHealth(); err != nil {
		lastFailedDestinatiomPing = time.Now()
	} else {
		lastSuccesfullDestinatiomPing = time.Now()
	}
}

func exportHandler(w http.ResponseWriter, r *http.Request) {
	res, err := exprtr.ExportMeasurements()
	if err != nil {
		logger.Error("Error running export ", err)
		render.Render(w, r, &RestResponse{HTTPStatusCode: http.StatusInternalServerError, StatusText: err.Error()}) // nolint
		return
	}
	render.JSON(w, r, res)
}

func failedHandler(w http.ResponseWriter, r *http.Request) {
	type noResultsToRender struct {
		Status string
	}

	measurements, err := repo.FindMeasurementsByStatus(repository.TEMP_FAILURE)
	if err != nil {
		logger.Error("Error talking with repository", err)

	}
	logger.Debug("Got number of failed measurements ", len(measurements))

	var results []backend.ExportResult

	for _, m := range measurements {
		logger.Debug("Processing - ", m)
		res, err := exprtr.ExportMeasurement(measurement.Measurement{}, m)
		if err != nil {
			logger.Error("Error exporting measurement")
		}
		results = append(results, res)
	}

	if err := exprtr.MarkPermanentFailed(); err != nil {
		logger.Error("Error handling ageing temp. failed measurements - ", err)
	}

	if len(measurements) == 0 {
		noResults := noResultsToRender{Status: "no measurements to export"}
		render.JSON(w, r, noResults)
	} else {
		render.JSON(w, r, results)
	}
}
