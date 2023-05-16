package backend

import (
	"fmt"
	"reflect"
	"time"

	"bitbucket.org/opentelehealth/exporter/app"
	"bitbucket.org/opentelehealth/exporter/backend/kih/exporttypes"
	"bitbucket.org/opentelehealth/exporter/backend/oioxds"
	"bitbucket.org/opentelehealth/exporter/measurement"
	"bitbucket.org/opentelehealth/exporter/repository"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var api measurement.MeasurementApi
var repo repository.Repository
var cfg *app.Config
var log *logrus.Logger

const OIOXDS_BACKEND = "oioxds"

func InitExporter(config *app.Config, measurementApi measurement.MeasurementApi, repos repository.Repository) (Exporter, error) {
	cfg = config
	api = measurementApi
	pkg := app.GetPackage(reflect.TypeOf(ExportResult{}).PkgPath())
	wantedLevel := config.GetLoggerLevel(pkg)
	log = app.NewLogger(wantedLevel)
	repo = repos
	exporter := exporterImpl{}
	log.Debug("Type: ", config.Export.Backend)

	switch config.Export.Backend {
	case OIOXDS_BACKEND:
		log.Debug("Setting up OIOXDS export ")
		oioxdsBackend := oioxds.InitExporter(config, api)
		exporter.exporter = oioxdsBackend
	default:
		log.Warnf("Unsupported backend - %s", config.Export.Backend)
		return &exporter, fmt.Errorf("Unsupported backend")
	}

	return &exporter, nil
}

type ExportBackend interface {
	ConvertMeasurement(m measurement.Measurement, mr repository.MeasurementExportState) (string, error)
	ExportMeasurement(string) (string, error)
	ShouldExport(m measurement.Measurement) bool
	GetExportTypes() map[string]exporttypes.MeasurementType
	CheckHealth() error
}

type exporterImpl struct {
	exporter ExportBackend
}

func MeasurementToMeasurementType(measurement measurement.Measurement) repository.MeasurementExportState {
	m := repository.MeasurementExportState{}
	m.Measurement = measurement.Links.Measurement
	m.Patient = measurement.Links.Patient
	return m
}

func (e exporterImpl) ShouldExport(m measurement.Measurement) bool {
	return e.exporter.ShouldExport(m)
}

func (e exporterImpl) CheckHealth() error {
	return e.exporter.CheckHealth()
}

// Loops over measurements and calculates if retry time is up
func (e exporterImpl) MarkPermanentFailed() error {
	hoursToWait := cfg.Export.DaysToRetry * 24
	log.Debug("Hours to stay in temp failure ", hoursToWait)
	measurements, err := repo.FindMeasurementsByStatus(repository.TEMP_FAILURE)

	if err != nil {
		log.Debugf("Trace %+v", err)
		return errors.Wrap(err, "Error search temporarily failed measurements")
	}

	log.Debug("Found ", len(measurements), " temp failed meassages")
	for _, v := range measurements {
		hours_parked := int(time.Since(v.CreatedAt.Time).Hours())
		if hours_parked > 24*cfg.Export.DaysToRetry {
			log.Debug("Marked temp failed for ", v, " temp failed for ", hours_parked, " hours")
			v.Status = repository.FAILED

			_, err := repo.UpdateMeasurement(v)
			if err != nil {
				log.Errorf("Error updating repository - %+v", err)
				log.Debugf("Trace %+v", err)
			}
		}
	}

	return nil
}

// ExportMeasurement takes a measurement and exports it and updates the repository with the new state
func (e exporterImpl) ExportMeasurement(othMeasurement measurement.Measurement, exportState repository.MeasurementExportState) (ExportResult, error) {
	startTime := time.Now()

	result := ExportResult{Success: true}
	var err error

	var localMeasurement measurement.Measurement

	if len(othMeasurement.Links.Measurement) == 0 {
		log.Debug("Measurement not - set - refresh")
		localMeasurement, err = api.FetchMeasurement(exportState.Measurement)
		if err != nil {
			log.Debugf("Trace %+v", err)
			return ExportResult{Success: false}, errors.Wrap(err, "Error refreshing measurementt")
		}
	} else {
		localMeasurement = othMeasurement
	}

	log.Debug("Using measurement ", localMeasurement, " patient ", localMeasurement.Links.Patient)

	exportState, err = repo.FindOrCreateMeasurement(exportState)
	if err != nil {
		result.Success = false
		log.Debugf("Trace %+v", err)

		return result, errors.Wrap(err, "Error exporting measurement")
	}

	if nil != e.exporter {
		res, err := e.exporter.ConvertMeasurement(localMeasurement, exportState)

		if err != nil {
			errmsg := fmt.Sprintf("Error converting measurement - id %s - %v", exportState.ID, err)
			log.Errorf(errmsg)
			log.Debugf("Trace %+v", errmsg)

			exportState.Status = repository.TEMP_FAILURE

			m, errdb := repo.UpdateMeasurement(exportState)
			if errdb != nil {
				log.Errorf("Error updating measurement %+v - %s", errdb, m)
			}

			result.Success = false
			result.Measurement = m
			return result, errors.Wrap(err, errmsg)
		}
		conversionTime := time.Now()
		log.Debug("M: ", exportState.ID.String(), " conversion=", conversionTime.Sub(startTime))
		// Do SOSI'fication

		// Send message
		log.Debug("Exporting ", exportState)
		_, err = e.exporter.ExportMeasurement(res)
		if err != nil {
			errmsg := fmt.Sprintf("Error exporting message- id %s - %s", exportState.ID, err)
			log.Errorf(errmsg)
			log.Debugf("Trace %+v", errmsg)

			exportState.Status = repository.TEMP_FAILURE

			m, errdb := repo.UpdateMeasurement(exportState)
			if errdb != nil {
				log.Error("Error updating measurment - ", errdb, m)
				log.Debugf("Trace %+v", errdb)
			}

			result.Measurement = m
			result.Success = false
			return result, errors.Wrap(err, errmsg)
		}

		log.Debug("M: ", exportState.ID.String(), " exportedtime=", time.Since(conversionTime))
		exportState.Status = repository.COMPLETED
		log.Debug("Setting ", exportState, " as completed")

		exportState, err = repo.UpdateMeasurement(exportState)
		if err != nil {
			log.Error("Error updating measurment")
			log.Debugf("Trace %+v", err)
		}

	} else {
		log.Error("Exporter is not set")
		exportState.Status = repository.TEMP_FAILURE

		exportState, err = repo.UpdateMeasurement(exportState)
		if err != nil {
			log.Error("Error updating measurment - ", err, exportState)
			log.Debugf("Trace %+v", err)
		}

		result.Success = false
	}
	result.Measurement = exportState

	return result, nil
}

func (e exporterImpl) ExportMeasurements() ([]ExportResult, error) {
	start := time.Now()
	startTime, _ := repo.StartExport()

	log.Debug("Using start time:", startTime.Lastrun.Format(time.RFC3339))

	exports := []ExportResult{}

	rejected := 0
	exported := 0
	failed := 0
	iteration := 0
	handled := 0
	var err error
	res := measurement.MeasurementResponse{}
	res.Total = cfg.ClinicianConfig.BatchSize + 1 // make sure we at least run onces

	// Handle pagination
	for i := 0; res.Offset+cfg.ClinicianConfig.BatchSize < res.Total; i++ {
		log.Debug("Off", res.Offset, " batch", cfg.ClinicianConfig.BatchSize, " total ", res.Total)
		res, err = api.FetchMeasurements(startTime.Lastrun, i*cfg.ClinicianConfig.BatchSize)
		if err != nil {
			return exports, err
		}
		for _, measurement := range res.Results {
			m := MeasurementToMeasurementType(measurement)

			m, err := repo.FindOrCreateMeasurement(m)
			if err != nil {
				log.Errorf("Error searching measurements - %+v", err)
				log.Debugf("Trace %+v", err)

				return exports, errors.Wrap(err, "Error getting measurement from DB ")
			}

			switch m.Status {
			case repository.COMPLETED:
				log.Debug("M, ", m, " is already completed")
				handled++
				continue
			case repository.NO_EXPORT:
				log.Debug("M, ", m, " is flagged as no-export")
				handled++
				continue
			case repository.FAILED:
				log.Debug("M, ", m, " is already flaggged failed")
				handled++
				continue
			default:
				export, ex, fai, re, _ := e.HandleMeasurement(measurement, m)
				exports = append(exports, export)
				rejected += re
				exported += ex
				failed += fai

			}
		}
		iteration++
	}

	if failed > 0 {
		startTime.Status = repository.FAILED
	} else {
		startTime.Status = repository.COMPLETED
	}
	if err := repo.UpdateExport(startTime); err != nil {
		return exports, errors.Wrap(err, "Error updating export")
	}

	log.Info(
		fmt.Sprintf("type=export uuid=%s completed=%s starttime=%s iterations=%d tt=%d total=%d exported=%d rejected=%d failed=%d",
			startTime.Id.String(), time.Now().Format(time.RFC3339),
			startTime.Lastrun.Format(time.RFC3339),
			iteration, time.Since(start).Milliseconds(),
			exported+failed+rejected, exported, rejected, failed))

	return exports, nil
}

// Handle by measurement
// - Checks if it should be exported
//   - If not, mark as NO_EXPORT ( status = 5 in DB) return 1 for rejected
//
// - Tries to export it
//   - If, mark as completed
//   - Retry if marked as temporarily failed
//
// - Update state in db
// - return result, and values indicating if exported,failed, or rejected
func (e exporterImpl) HandleMeasurement(othMeasurement measurement.Measurement, exportState repository.MeasurementExportState) (ExportResult, int, int, int, error) {
	var export ExportResult
	var err error
	failed := 0
	exported := 0
	rejected := 0
	startTime := time.Now()

	if e.exporter.ShouldExport(othMeasurement) {
		log.Debug("Handling measuremnt - ", exportState)

		if exportState.Status != repository.COMPLETED && exportState.Status != repository.NO_EXPORT {
			export, err = e.ExportMeasurement(othMeasurement, exportState)
			if err != nil {
				log.Error("Error exporting measurement")
				log.Debugf("Trace %+v", err)

				failed++
			} else {
				exportState.Status = repository.COMPLETED
				_, err := repo.UpdateMeasurement(exportState)
				if err != nil {
					log.Error("Error updating repository - ", exportState, " - ", err)
					log.Debugf("Trace %+v", err)
				}
				exported++
			}
		}
	} else {
		exportState.Status = repository.NO_EXPORT
		export.Success = false
		exportState, err = repo.UpdateMeasurement(exportState)
		if err != nil {
			return export, exported, failed, rejected, errors.Wrap(err, "Error exporting measurement")
		}
		rejected++
		log.Debug("Noexport uuid=", exportState.ID.String(), " status=",
			repository.StatusToText(exportState.Status), " type=", othMeasurement.Type)
		export.Measurement = exportState

	}
	log.Info(fmt.Sprintf("type=measurement uuid=%s status=%s exporttotal=%d ms", exportState.ID.String(),
		repository.StatusToText(exportState.Status), time.Since(startTime).Milliseconds()))

	return export, exported, failed, rejected, nil
}
