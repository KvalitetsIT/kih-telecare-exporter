package measurement

import (
	"fmt"
	"strings"
	"time"
)

type PatientResult struct {
	CreatedDate   string      `json:"createdDate"`
	UniqueID      string      `json:"uniqueId"`
	Username      string      `json:"username"`
	FirstName     string      `json:"firstName"`
	LastName      string      `json:"lastName"`
	DateOfBirth   interface{} `json:"dateOfBirth"`
	Sex           string      `json:"sex"`
	Status        string      `json:"status"`
	Address       string      `json:"address"`
	PostalCode    string      `json:"postalCode"`
	City          string      `json:"city"`
	Place         interface{} `json:"place"`
	Phone         interface{} `json:"phone"`
	MobilePhone   string      `json:"mobilePhone"`
	Email         string      `json:"email"`
	Comment       interface{} `json:"comment"`
	PatientGroups []struct {
		Name  string `json:"name"`
		Links struct {
			PatientGroup string `json:"patientGroup"`
		} `json:"links"`
	} `json:"patientGroups"`
	Relatives []interface{} `json:"relatives"`
	Links     struct {
		Self                   string `json:"self"`
		QuestionnaireSchedules string `json:"questionnaireSchedules"`
		Measurements           string `json:"measurements"`
		QuestionnaireResults   string `json:"questionnaireResults"`
		PatientThresholds      string `json:"patientThresholds"`
	} `json:"links"`
}

func (p PatientResult) String() string {
	return fmt.Sprintf("%s %s - %v", p.FirstName, p.LastName, p.DateOfBirth)
}

type Ignored struct {
	By     By     `json:"by,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type By struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Links     Links  `json:"links"`
}

/// MeasurementValue denotes an OTH measurement from the REST API
type MeasurementValue struct {
	Unit                 string      `json:"unit"`
	Value                interface{} `json:"value,omitempty"`
	IsAfterMeal          bool        `json:"isAfterMeal,omitempty"`
	IsBeforeMeal         bool        `json:"isBeforeMeal,omitempty"`
	IsControlMeasurement bool        `json:"isControlMeasurement,omitempty"`
	Systolic             float32     `json:"systolic,omitempty"`
	Diastolic            float32     `json:"diastolic,omitempty"`
	Ignored              Ignored     `json:"ignored,omitempty"`
}

type Links struct {
	Measurement string `json:"measurement,omitempty"`
	Patient     string `json:"patient,omitempty"`
	Clinician   string `json:"clinician,omitempty"`
}

type PrimaryDeviceIdentifier struct {
	MacAddress string `json:"macAddress"`
}
type Other struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}
type AdditionalDeviceIdentifiers struct {
	SystemID string `json:"systemId,omitempty"`
	Other    Other  `json:"other,omitempty"`
}
type DeviceMeasurement struct {
	ConnectionType              string                        `json:"connectionType"`
	Manufacturer                string                        `json:"manufacturer"`
	Model                       string                        `json:"model"`
	PrimaryDeviceIdentifier     PrimaryDeviceIdentifier       `json:"primaryDeviceIdentifier"`
	HardwareVersion             string                        `json:"hardwareVersion"`
	FirmwareVersion             string                        `json:"firmwareVersion"`
	SoftwareVersion             string                        `json:"softwareVersion"`
	AdditionalDeviceIdentifiers []AdditionalDeviceIdentifiers `json:"additionalDeviceIdentifiers"`
}

func (dm DeviceMeasurement) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("[%s/%s]", dm.Manufacturer, dm.Model))

	if len(dm.HardwareVersion) > 0 {
		sb.WriteString(fmt.Sprintf(" HV: %s", dm.HardwareVersion))
	}
	if len(dm.FirmwareVersion) > 0 {
		sb.WriteString(fmt.Sprintf(" FV: %s", dm.FirmwareVersion))
	}
	if len(dm.SoftwareVersion) > 0 {
		sb.WriteString(fmt.Sprintf(" SV: %s", dm.SoftwareVersion))
	}

	return sb.String()
}

/// Origin denotes how the measurement was takex
type Origin struct {
	ManualMeasurement struct {
		EnteredBy string `json:"enteredBy"`
	} `json:"manualMeasurement,omitempty"`
	DeviceMeasurement DeviceMeasurement `json:"deviceMeasurement,omitempty"`
}

func (o Origin) String() string {
	var sb strings.Builder

	if len(o.ManualMeasurement.EnteredBy) > 0 {
		sb.WriteString(fmt.Sprintf("Manual Measurement - clinician: %s", o.ManualMeasurement.EnteredBy))
	}

	if len(o.DeviceMeasurement.Manufacturer) > 0 {
		sb.WriteString("DM: ")
		sb.WriteString(o.DeviceMeasurement.String())
	}
	return sb.String()
}

type Measurement struct {
	Timestamp   time.Time        `json:"timestamp"`
	Type        string           `json:"type"`
	Measurement MeasurementValue `json:"measurement,omitempty"`
	Origin      Origin           `json:"origin,omitempty"`
	Links       Links            `json:"links"`
}

func (m Measurement) String() string {

	if m.Type == "blood_pressure" {
		return fmt.Sprintf("[%s] t: %s - %f/%f %s", m.Timestamp.Format(time.RFC822), m.Type, m.Measurement.Systolic, m.Measurement.Diastolic, m.Measurement.Unit)
	} else {
		return fmt.Sprintf("[%s] t: %s - %f %s patient: %s", m.Timestamp.Format(time.RFC822), m.Type, m.Measurement.Value, m.Measurement.Unit, m.Links.Patient)
	}
}

type MeasurementResponse struct {
	Results []Measurement `json:"results"`
	Total   int           `json:"total"`
	Max     int           `json:"max"`
	Offset  int           `json:"offset"`
	Links   struct {
		Self     string `json:"self"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
	} `json:"links"`
}
