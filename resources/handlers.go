package resources

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
	"github.com/KvalitetsIT/kih-telecare-exporter/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type MeasurementResponse struct {
	Patient           measurement.PatientResult         `json:"patient"`
	Measurement       measurement.Measurement           `json:"measurement"`
	StoredMeasurement repository.MeasurementExportState `json:"storedMeasurement"`
}

// measurementHandler retrieves measurement by id and returns the patient and measurement
func measurementHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "measurement")
	logger.Debugf("Requesting information for id: %s", id)

	mes, err := repo.FindMeasurement(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			render.Render(w, r, &RestResponse{HTTPStatusCode: http.StatusNotFound, StatusText: fmt.Sprintf("Meaurement %s not found", id)}) // nolint
			return
		}
		logger.Error("Error running export ", err)
		render.Render(w, r, &RestResponse{HTTPStatusCode: http.StatusInternalServerError, StatusText: err.Error()}) // nolint
		return
	}
	logger.Debugf("Got %v", mes)

	measurement, err := api.FetchMeasurement(mes.Measurement)
	if err != nil {
		logger.Error("Error running export ", err)
		render.Render(w, r, &RestResponse{HTTPStatusCode: http.StatusInternalServerError, StatusText: err.Error()}) // nolint
		return
	}
	res := MeasurementResponse{}
	res.StoredMeasurement = mes
	res.Measurement = measurement

	patient, err := api.FetchPatient(res.Measurement.Links.Patient)
	if err != nil {
		logger.Error("Error running export ", err)
		render.Render(w, r, &RestResponse{HTTPStatusCode: http.StatusInternalServerError, StatusText: err.Error()}) // nolint
		return
	}

	res.Patient = patient
	render.JSON(w, r, res)
}
