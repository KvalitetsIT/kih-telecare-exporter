package oioxds

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"

	"bitbucket.org/opentelehealth/exporter/app"
	"bitbucket.org/opentelehealth/exporter/backend/kih/exporttypes"
	"bitbucket.org/opentelehealth/exporter/backend/kih/shared"
	"bitbucket.org/opentelehealth/exporter/internal"
	"bitbucket.org/opentelehealth/exporter/measurement"
	"bitbucket.org/opentelehealth/exporter/repository"

	"github.com/akyoto/cache"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Initialize the OIO XDS exporter backend
func InitExporter(appConfig *app.Config, api measurement.MeasurementApi) OioXdsExporter {
	pkg := app.GetPackage(reflect.TypeOf(OioXdsExporter{}).PkgPath())
	log = app.NewLogger(appConfig.GetLoggerLevel(pkg))
	log.Debug("OIO XDS ", pkg, " -  loglevel", appConfig.GetLoggerLevel(pkg))

	config = appConfig

	exportURL := appConfig.Export.OIOXDSExport.XdsGenerator.URL
	healthCheckURL := appConfig.Export.OIOXDSExport.XdsGenerator.HealthCheck
	c := cache.New(1 * time.Hour)

	log.Info("Export URL: ", exportURL, " - health check URL: ", healthCheckURL)

	var httpClient http.Client
	if appConfig.Export.OIOXDSExport.SkipSslVerify {
		log.Debug("Setting TLS verify to true")
		httpClient = http.Client{}
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	} else {
		httpClient = http.Client{}
	}

	// Setup XDS generator
	log.Debugf("Is XDS generator setup? %s health: %s", appConfig.Export.OIOXDSExport.XdsGenerator.URL, appConfig.Export.OIOXDSExport.XdsGenerator.HealthCheck)
	// Remember to setup the logger
	shared.Init(appConfig)

	exporterBackend := OioXdsExporter{c: c, api: api, config: appConfig}
	exporterBackend.healthCheckURL = config.Export.OIOXDSExport.XdsGenerator.HealthCheck
	exporterBackend.exportURL = config.Export.OIOXDSExport.XdsGenerator.URL
	exporterBackend.exportedTypes = exporttypes.GetOioXdsExportTypes()
	return exporterBackend

}

// Returns the exported types handled by this exporter
func (exprt OioXdsExporter) GetExportTypes() map[string]exporttypes.MeasurementType {
	return exprt.exportedTypes
}

// Checks whether a measurement should be exported
func (exprt OioXdsExporter) ShouldExport(m measurement.Measurement) bool {
	measurementtype, ok := exprt.exportedTypes[m.Type]

	if !ok {
		return false
	} else {
		return measurementtype.IsToBeExported()
	}
}

// Perform health check of required backends (kihdb and sosiserver)
func (exprt OioXdsExporter) CheckHealth() error {
	log.Debugf("Performing health check against %s", exprt.healthCheckURL)

	if err := internal.PerformHealthCheck(exprt.client, http.MethodGet, http.StatusOK, exprt.healthCheckURL); err != nil {
		log.Errorf("Received error %v", err)
		return errors.Wrap(err, "Error testing OIOXDS generator health")
	}
	return nil
}

// Export the measurement
func (exprt OioXdsExporter) ExportMeasurement(s string) (string, error) {
	log.Debug("Exporting measurement - ", exprt.exportURL)

	// Convert
	xdsRequest, err := http.NewRequest(http.MethodPost, exprt.config.Export.OIOXDSExport.XdsGenerator.URL, strings.NewReader(s))
	if err != nil {
		return "", errors.Wrap(err, "Error creating HTTP request to XDS generator")
	}
	xdsRequest.Header.Add("Content-Type", "application/json")

	resp, err := exprt.client.Do(xdsRequest)
	if err != nil {
		return "", errors.Wrap(err, "Error submitting request to XDS generator")
	}
	defer resp.Body.Close()

	log.Debugf("Received: %v", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Error readings XDS generator reply")
	}
	defer func() {
		resp.Body.Close()
	}()

	log.Debugf("Got %s", string(body))

	if resp.StatusCode > 400 {
		var errResponse XdsErrorResponse

		if err := json.Unmarshal(body, &errResponse); err != nil {
			return "", errors.Wrap(err, "Error parsing body from XdsGenerator")
		}

		log.Warnf("Status: %v", resp.Status)
		return resp.Status, fmt.Errorf("Server said %v", errResponse.Message)
	}

	return "", nil
}

// ConvertMeasurement is converts the internal Measuremnet struct into the format for the specific backend. Returns a string
func (exprt OioXdsExporter) ConvertMeasurement(m measurement.Measurement, mr repository.MeasurementExportState) (string, error) {
	startTime := time.Now()
	log.Debug("Starting conversion of ", m)

	s := SelfMonitoredSample{}
	s.CreatedByText = config.Export.CreatedBy

	var reports []shared.LaboratoryReportExtended

	reports, err := shared.ReportFromMeasurement(exprt.exportedTypes, m, mr)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("Error parseing measurement - %v", err))
	}

	s.LaboratoryReports = reports

	var patient measurement.PatientResult
	p, found := exprt.c.Get(mr.Patient)
	if !found {
		log.Debug("Fetching patient data")
		var err error
		patient, err = exprt.api.FetchPatient(mr.Patient)
		if err != nil {
			log.Errorf("Error retrieving patient information - %v", err)
			return "", errors.Wrap(err, "Error retriving information")
		}

		log.Debug("Add information to Cache")
		exprt.c.Set(mr.Patient, patient, 1*time.Hour)
	} else {
		log.Debug("Found patient in cache")
		patient = p.(measurement.PatientResult)
	}

	//s.CreatedByText = config.Export.KIHExport.CreatedBy

	//var reports []shared.LaboratoryReportExtended

	xdsGeneratorRequest, err := convertXdsGeneratorRequest(mr.ID, s, patient)
	if err != nil {
		return "", errors.Wrap(err, "Error creating XDS generator request")
	}

	log.Debug("type=conversion uuid= ", mr.ID.String(), " tt=", time.Since(startTime), " done")

	return string(xdsGeneratorRequest), nil
}

// handles mapping ot OTH patient to KIH Citizen
func mapCitizenToPatient(citizen *Citizen, patient measurement.PatientResult) {
	if len(patient.FirstName) > 0 {
		name := &Name{}
		name.PersonGivenName = patient.FirstName
		if len(patient.LastName) > 0 {
			name.PersonSurName = patient.LastName
		}

		citizen.Person = name
	}

	if len(patient.Address) > 0 {
		address := &Address{
			StreetName:         patient.Address,
			PostCodeIdentifier: patient.PostalCode,
			MunicipalityName:   patient.City,
		}
		citizen.Address = address
	}

	if len(patient.MobilePhone) > 0 {
		phone := &Phone{
			PhoneNumberIdentifier: patient.MobilePhone,
			PhoneNumberUse:        "W",
		}

		citizen.Phone = phone
	}
}

// converts to XDS generator format and converts to []byte for posting to backend
func convertXdsGeneratorRequest(uuid uuid.UUID, s SelfMonitoredSample, patient measurement.PatientResult) ([]byte, error) {
	xdsGeneratorRequest := XdsGeneratorRequest{}
	xdsGeneratorRequest.DocumentUuid = uuid

	log.Debugf("Person: %+v", patient)

	collectionList := SelfMonitoringCollection{}
	citizen := &collectionList.Citizen
	citizen.PersonCivilRegistrationIdentifier = patient.UniqueID

	mapCitizenToPatient(citizen, patient)

	mySample := SelfMonitoringSamples{SelfMonitoringSample: s}
	mySamples := []SelfMonitoringSamples{mySample}
	collectionList.SelfMonitoringSamples = mySamples

	xdsGeneratorRequest.SelfMonitoringCollection = []SelfMonitoringCollection{collectionList}

	jsonBody, err := json.Marshal(xdsGeneratorRequest)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Error marshalling OIO Request")
	}

	log.Debugf("Sending: \n%s - bytes %d", string(jsonBody), len(string(jsonBody)))

	return jsonBody, nil
}
