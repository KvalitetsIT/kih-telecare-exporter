package oioxds

import (
	"encoding/xml"
	"net/http"
	"time"

	"bitbucket.org/opentelehealth/exporter/app"
	"bitbucket.org/opentelehealth/exporter/backend/kih/exporttypes"
	"bitbucket.org/opentelehealth/exporter/backend/kih/shared"
	"bitbucket.org/opentelehealth/exporter/measurement"
	"github.com/akyoto/cache"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Name struct {
	PersonGivenName  string `json:"personGivenName,omitempty"`
	PersonMiddleName string `json:"personMiddleName,omitempty"`
	PersonSurName    string `json:"personSurName,omitempty"`
}

type Email struct {
	EmailAddressIdentifier string `json:"emailAddressIdentifier,omitempty"`
	EmailAddressUse        string `json:"emailAddressUse,omitempty"`
}

type Address struct {
	StreetName               string `json:"streetName,omitempty"`
	StreetBuildingIdentifier string `json:"streetBuildingIdentifier,omitempty"`
	PostCodeIdentifier       string `json:"postCodeIdentifier,omitempty"`
	MunicipalityName         string `json:"municipalityName,omitempty"`
}

type Phone struct {
	PhoneNumberIdentifier string `json:"phoneNumberIdentifier,omitempty"`
	PhoneNumberUse        string `json:"phoneNumberUse,omitempty"`
}

type Citizen struct {
	PersonCivilRegistrationIdentifier string   `json:"personCivilRegistrationIdentifier"`
	Person                            *Name    `json:"personNameStructure,omitempty"`
	Email                             *Email   `json:"emailAddress,omitempty"`
	Address                           *Address `json:"addressPostal,omitempty"`
	Phone                             *Phone   `json:"phoneNumberSubscriber,omitempty"`
}

type XdsGeneratorRequest struct {
	DocumentUuid             uuid.UUID                  `json:"DocumentUUID,omitempty"`
	SelfMonitoringCollection []SelfMonitoringCollection `json:"SelfMonitoringCollection"`
}

type SelfMonitoringSamples struct {
	SelfMonitoringSample SelfMonitoredSample `json:"SelfMonitoringSample"`
}

type SelfMonitoredSample struct {
	CreatedByText     string                            `json:"CreatedByText"`
	LaboratoryReports []shared.LaboratoryReportExtended `json:"LaboratoryReports"`
}

type SelfMonitoringCollection struct {
	Citizen               Citizen                 `json:"Citizen"`
	SelfMonitoringSamples []SelfMonitoringSamples `json:"SelfMonitoringSamples"`
}

type OioXdsExporter struct {
	c              *cache.Cache
	client         http.Client
	skipSOSI       bool
	config         *app.Config
	healthCheckURL string
	exportURL      string
	exportedTypes  map[string]exporttypes.MeasurementType
	api            measurement.MeasurementApi
}

var log *logrus.Logger

var config *app.Config

const (
	MEASURED_BY_PATIENT                 = "Patient målt"
	POT_CODE                            = "POT"
	NATIONAL_SAMPLE_IDENTIFIER          = "9999999999"
	UNSPECIFIED                         = "unspecified"
	CLINICIAN                           = "clinician"
	MEASUREMENT_TRANSFER_TYPE           = "automatic"
	MEASUREMENT_LOCATION_TYPE           = "home"
	MEAUREMENT_SCHEDULED_TYPE           = "scheduled"
	MEAUREMENT_UNSCHEDULED_TYPE         = "unscheduled"
	RESULT_ENCODING_NUMERIC             = "numeric"
	RESULT_ENCODING_ALPHANUMERIC        = "alphanumeric"
	MEASUREMENT_TRANSFERED_BY_HCPROF    = "typedbyhcprof"
	MEASUREMENT_TRANSFERED_BY_TYPED     = "typed"
	MEASUREMENT_TRANSFERED_BY_AUTOMATIC = "automatic"
)

type XdsResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"s,attr"`
	A       string   `xml:"a,attr"`
	Header  struct {
		Text   string `xml:",chardata"`
		Action struct {
			Text           string `xml:",chardata"`
			MustUnderstand string `xml:"mustUnderstand,attr"`
		} `xml:"Action"`
		RelatesTo string `xml:"RelatesTo"`
	} `xml:"Header"`
	Body struct {
		Text             string `xml:",chardata"`
		RegistryResponse struct {
			Text           string `xml:",chardata"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Status         string `xml:"status,attr"`
			Rs             string `xml:"rs,attr"`
			Xsi            string `xml:"xsi,attr"`
		} `xml:"RegistryResponse"`
	} `xml:"Body"`
}

type XdsErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
	Message   string    `json:"message"`
	Path      string    `json:"path"`
}
