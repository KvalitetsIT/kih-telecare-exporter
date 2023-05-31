// Package shared holds the structures shared between versions of the KIH SOAP interface
package shared

import (
	"encoding/xml"
	"fmt"

	"github.com/KvalitetsIT/kih-telecare-exporter/app"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var config *app.Config

func Init(conf *app.Config) {
	log = logrus.New()
	config = conf
	log.SetLevel(config.Logger.Level)
}

type ProducerOfLabResult struct {
	Identifier     string `xml:"urn:oio:medcom:chronicdataset:1.0.0 Identifier"`
	IdentifierCode string `xml:"urn:oio:medcom:chronicdataset:1.0.0 IdentifierCode"`
}
type InstrumentType struct {
	MedComID        string `xml:"urn:oio:medcom:chronicdataset:1.0.1 MedComID,omitempty" json:"MedComId,omitempty"`
	Manufacturer    string `xml:"urn:oio:medcom:chronicdataset:1.0.1 Manufacturer,omitempty" json:"Manufacturer,omitempty"`
	ProductType     string `xml:"urn:oio:medcom:chronicdataset:1.0.1 ProductType,omitempty" json:"ProductType,omitempty"`
	Model           string `xml:"urn:oio:medcom:chronicdataset:1.0.1 Model,omitempty" json:"Model,omitempty"`
	SoftwareVersion string `xml:"urn:oio:medcom:chronicdataset:1.0.1 SoftwareVersion,omitempty" json:"SoftwareVersion,omitempty"`
}

func (i InstrumentType) String() string {
	return fmt.Sprintf("[%s] type: %s, model %s, version %s", i.Manufacturer, i.ProductType, i.Model, i.SoftwareVersion)
}

type LaboratoryReportExtended struct {
	UuidIdentifier                string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 UuidIdentifier" json:"UuidIdentifier,omitempty"`
	CreatedDateTime               string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 CreatedDateTime" json:"CreatedDateTime,omitempty"`
	AnalysisText                  string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 AnalysisText" json:"AnalysisText,omitempty"`
	ResultText                    string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 ResultText" json:"ResultText,omitempty"`
	ResultEncodingIdentifier      string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 ResultEncodingIdentifier,omitempty" json:"ResultEncodingIdentifier,omitempty"`
	ResultOperatorIdentifier      string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 ResultOperatorIdentifier,omitempty" json:"ResultOperatorIdentifier,omitempty"`
	ResultUnitText                string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 ResultUnitText" json:"ResultUnitText,omitempty"`
	ResultAbnormalIdentifier      string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 ResultAbnormalIdentifier,omitempty" json:"ResultAbnormalIdentifier,omitempty"`
	ResultMinimumText             string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 ResultMinimumText,omitempty" json:"ResultMinimumText,omitempty"`
	ResultMaximumText             string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 ResultMaximumText,omitempty" json:"ResultMaximumText,omitempty"`
	ResultTypeOfInterval          string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 ResultTypeOfInterval,omitempty" json:"ResultTypeOfInterval,omitempty"`
	NationalSampleIdentifier      string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 NationalSampleIdentifier" json:"NationalSampleIdentifier,omitempty"`
	IupacIdentifier               string               `xml:"urn:oio:medcom:chronicdataset:1.0.0 IupacIdentifier" json:"IupacIdentifier,omitempty"`
	ProducerOfLabResult           *ProducerOfLabResult `xml:"urn:oio:medcom:chronicdataset:1.0.0 ProducerOfLabResult" json:"ProducerOfLabResult,omitempty"`
	Instrument                    *InstrumentType      `xml:"urn:oio:medcom:chronicdataset:1.0.0 Instrument,omitempty" json:"Instrument,omitempty"`
	MeasurementTransferredBy      string               `xml:"urn:oio:medcom:chronicdataset:1.0.1 MeasurementTransferredBy" json:"MeasurementTransferredBy,omitempty"`
	MeasurementLocation           string               `xml:"urn:oio:medcom:chronicdataset:1.0.1 MeasurementLocation" json:"MeasurementLocation,omitempty"`
	MeasuringDataClassification   string               `xml:"urn:oio:medcom:chronicdataset:1.0.1 MeasuringDataClassification,omitempty" json:"MeasuringDataClassification,omitempty"`
	MeasurementDuration           string               `xml:"urn:oio:medcom:chronicdataset:1.0.1 MeasurementDuration,omitempty" json:"MeasurementDuration,omitempty"`
	MeasurementScheduled          string               `xml:"urn:oio:medcom:chronicdataset:1.0.1 MeasurementScheduled" json:"MeasurementScheduled,omitempty"`
	HealthCareProfessionalComment string               `xml:"urn:oio:medcom:chronicdataset:1.0.1 HealthCareProfessionalComment,omitempty" json:"HealthCareProfessionalComment,omitempty"`
	MeasuringCircumstances        string               `xml:"urn:oio:medcom:chronicdataset:1.0.1 MeasuringCircumstances,omitempty" json:"MeasuringCircumstances,omitempty"`
}

type SelfMonitoredSample struct {
	LaboratoryReportExtendedCollection struct {
		LaboratoryReportExtended []LaboratoryReportExtended `xml:"urn:oio:medcom:chronicdataset:1.0.1 LaboratoryReportExtended"`
	} `xml:"urn:oio:medcom:chronicdataset:1.0.1 LaboratoryReportExtendedCollection"`
	CreatedByText string `xml:"urn:oio:medcom:chronicdataset:1.0.0 CreatedByText"`
}

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	//	soapenv string   `xml:"xmlns:soapenv,attr"`
	//	URN     string   `xml:"xmlns:urn,attr"`
	//	Urn1    string   `xml:"xmlns:urn1,attr"`
	//	Urn2    string   `xml:"xmlns:urn2,attr"`
	///dgws.HeaderAttributes

	Body interface{} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}
