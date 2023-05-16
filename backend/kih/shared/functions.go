package shared

import (
	"fmt"
	"time"

	"bitbucket.org/opentelehealth/exporter/backend/kih/exporttypes"
	"bitbucket.org/opentelehealth/exporter/measurement"
	"bitbucket.org/opentelehealth/exporter/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

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

// Map OTH Measurment to Laboratory Report structure
func ReportFromMeasurement(exportedTypes map[string]exporttypes.MeasurementType, m measurement.Measurement, mr repository.MeasurementExportState) ([]LaboratoryReportExtended, error) {

	var reports []LaboratoryReportExtended
	var err error
	exportType, ok := exportedTypes[m.Type]

	if !ok {
		return reports, fmt.Errorf("Export type for measurement type %s not found", m.Type)
	}

	switch exportType.(type) {
	case exporttypes.SimpleType:
		reports, err = handleSimpleType(reports, m, mr, exportType)
		if err != nil {
			return reports, fmt.Errorf("Error converting to format - %v", err)
		}

	case exporttypes.BloodPressureType:
		reports, err = handleBloodPressure(reports, m, mr, exportType)
		if err != nil {
			return reports, fmt.Errorf("Error converting to format - %v", err)
		}
	default:
		return reports, fmt.Errorf("Unknown measurement type for %v", exportType)
	}

	log.Debug("Returning - # of reports ", len(reports))
	return reports, nil
}

// handle measurements device to instrument based on MedCom standard decisions
func measurmentToInstrument(exportType exporttypes.MeasurementType, m measurement.Measurement) *InstrumentType {
	log.Debug("mapping origin ", m.Origin.DeviceMeasurement)
	var instrument *InstrumentType

	var device exporttypes.MedicalDevice
	for _, d := range exportType.GetDevices() {
		log.Debug("Found - ", d)
		if d.CheckIfSameModel(m.Origin) {
			log.Debug("Origin", m.Origin, " MATCHES", d)
			device = d
		} else {
			log.Debug("Origin", m.Origin, " does not match", d)
		}
	}
	if nil != device {
		log.Debug("Device found")
		instrument = &InstrumentType{}
		instrument.SoftwareVersion = device.GetManufacturer()
		instrument.MedComID = device.GetMedcomID()
		instrument.Model = device.GetModel()
	} else {
		log.Debug("Value of device is ", device)
	}

	return instrument
}

// map origin to OIO instrument structure
func mapMeasurmentToInstrument(m measurement.Measurement) *InstrumentType {
	log.Debug("mapping origin ", m.Origin.DeviceMeasurement)
	instrument := &InstrumentType{}

	instrument.Manufacturer = m.Origin.DeviceMeasurement.Manufacturer
	instrument.Model = m.Origin.DeviceMeasurement.Model
	instrument.ProductType = m.Origin.DeviceMeasurement.ConnectionType

	if len(m.Origin.DeviceMeasurement.FirmwareVersion) > 0 {
		instrument.SoftwareVersion = m.Origin.DeviceMeasurement.FirmwareVersion
	}
	if len(m.Origin.DeviceMeasurement.SoftwareVersion) > 0 {
		if len(instrument.SoftwareVersion) > 0 {
			instrument.SoftwareVersion = fmt.Sprintf("%s/%s", instrument.SoftwareVersion, m.Origin.DeviceMeasurement.SoftwareVersion)

		} else {
			instrument.SoftwareVersion = m.Origin.DeviceMeasurement.SoftwareVersion
		}
	}

	if len(m.Origin.DeviceMeasurement.HardwareVersion) > 0 {
		if len(instrument.SoftwareVersion) > 0 {
			instrument.SoftwareVersion = fmt.Sprintf("%s/%s", instrument.SoftwareVersion, m.Origin.DeviceMeasurement.HardwareVersion)

		} else {
			instrument.SoftwareVersion = m.Origin.DeviceMeasurement.HardwareVersion
		}
	}

	if len(instrument.SoftwareVersion) == 0 {
		instrument.SoftwareVersion = instrument.Model
	}

	return instrument
}

// Handle Origin fill out instrument data and values for transfer based on measurements Origin
func handleOrigin(exportType exporttypes.MeasurementType, m measurement.Measurement, lr *LaboratoryReportExtended) error {
	log.Debugf("Got origin: %+v", m.Origin)
	if len(m.Origin.ManualMeasurement.EnteredBy) > 0 {
		if "clinician" == m.Origin.ManualMeasurement.EnteredBy {
			lr.MeasurementTransferredBy = MEASUREMENT_TRANSFERED_BY_HCPROF
		} else {
			lr.MeasurementTransferredBy = MEASUREMENT_TRANSFERED_BY_TYPED
		}
	}
	log.Debug(fmt.Sprintf("Device: %s", m.Origin.DeviceMeasurement.Model))
	if len(m.Origin.DeviceMeasurement.Model) == 0 {
		log.Debug("Empty instrument")
	} else {
		lr.MeasurementTransferredBy = MEASUREMENT_TRANSFERED_BY_AUTOMATIC

		if config.Export.NoDeviceWhiteList {
			log.Debug("Using device information from measurement")
			lr.Instrument = mapMeasurmentToInstrument(m)
		} else {
			log.Debug("Using device information setup")
			lr.Instrument = measurmentToInstrument(exportType, m)
		}
	}

	if len(lr.MeasurementTransferredBy) == 0 {
		log.Debug("Setting default value")
		lr.MeasurementTransferredBy = MEASUREMENT_TRANSFERED_BY_TYPED
	}

	return nil
}

// Set handle generic mappings
func performGenericMapping(lr *LaboratoryReportExtended) {
	lr.ProducerOfLabResult = &ProducerOfLabResult{}
	lr.ProducerOfLabResult.IdentifierCode = POT_CODE
	lr.ProducerOfLabResult.Identifier = MEASURED_BY_PATIENT
	lr.MeasurementTransferredBy = MEASUREMENT_TRANSFER_TYPE
	lr.MeasurementLocation = MEASUREMENT_LOCATION_TYPE
	lr.MeasurementScheduled = MEAUREMENT_SCHEDULED_TYPE
	lr.NationalSampleIdentifier = NATIONAL_SAMPLE_IDENTIFIER
	lr.ResultTypeOfInterval = UNSPECIFIED
}

// func performMeasurementMapping(m measurement.Measurement, lr *LaboratoryReportExtended) error {
// 	exportType, ok := exportedTypes[m.Type]

// 	if !ok {
// 		return fmt.Errorf("Export type for measurement type %s not found", m.Type)
// 	}

// 	switch exportType.(type) {
// 	case exporttypes.SimpleType:
// 		err := performSimpleMapping(m, lr, exportType)
// 	case exporttypes.BloodPressureType:
// 		t := exportType.(exporttypes.BloodPressureType)
// 		return t.GetNpuCodes()
// 	default:
// 		return fmt.Errorf("Unknown measurement type for %v", exportType)
// 	}

// 	lr.AnalysisText = m.Type
// 	lr.ResultUnitText = m.Measurement.Unit

// 	lr.ResultText = fmt.Sprintf("%f", m.Measurement.Value)

// 	/*
// 	   if (isAlphanumeric) {
// 	       type.resultEncodingIdentifier = EncodingIdentifierType.ALPHANUMERIC
// 	   } else {
// 	       type.resultEncodingIdentifier = EncodingIdentifierType.NUMERIC
// 	   }
// 	*/

// 	return nil
// }

func performBaseMapping(exportType exporttypes.MeasurementType, lr *LaboratoryReportExtended, m measurement.Measurement, mr repository.MeasurementExportState) {
	lr.UuidIdentifier = mr.ID.String()
	lr.CreatedDateTime = m.Timestamp.Format(time.RFC3339)
	lr.IupacIdentifier = exportType.GetNpuCode()
}

func handleSimpleType(reports []LaboratoryReportExtended, m measurement.Measurement, mr repository.MeasurementExportState, exportType exporttypes.MeasurementType) ([]LaboratoryReportExtended, error) {
	r := LaboratoryReportExtended{}
	performBaseMapping(exportType, &r, m, mr)
	// Do the generic mapping
	performGenericMapping(&r)

	if err := handleOrigin(exportType, m, &r); err != nil {
		return reports, errors.Wrap(err, fmt.Sprintf("Error parsing origin %v", err))
	}

	r.AnalysisText = exportType.GetAnalysisText()
	r.ResultUnitText = exportType.GetResultUnitText()
	r.ResultText = exportType.GetResultText(m)

	reports = append(reports, r)

	return reports, nil
}

func handleBloodPressure(reports []LaboratoryReportExtended, m measurement.Measurement, mr repository.MeasurementExportState, exportType exporttypes.MeasurementType) ([]LaboratoryReportExtended, error) {
	// var err error
	t := exportType.(exporttypes.BloodPressureType)
	systolic := LaboratoryReportExtended{}
	diastolic := LaboratoryReportExtended{}
	performBaseMapping(exportType, &systolic, m, mr)
	performGenericMapping(&systolic)
	if err := handleOrigin(exportType, m, &systolic); err != nil {
		return reports, errors.Wrap(err, fmt.Sprintf("Error parsing origin %v", err))
	}

	systolic.IupacIdentifier = t.GetSystolic().GetNpuCode()
	systolic.AnalysisText = t.GetSystolic().GetAnalysisText()
	systolic.ResultUnitText = t.GetSystolic().GetResultUnitText()
	systolic.ResultText = t.GetSystolic().GetResultText(m)

	performBaseMapping(exportType, &diastolic, m, mr)
	performGenericMapping(&diastolic)
	if err := handleOrigin(exportType, m, &diastolic); err != nil {
		return reports, errors.Wrap(err, fmt.Sprintf("Error parsing origin %v", err))
	}

	diastolic.UuidIdentifier = uuid.New().String()
	diastolic.IupacIdentifier = t.GetDiastolic().GetNpuCode()
	diastolic.AnalysisText = t.GetDiastolic().GetAnalysisText()
	diastolic.ResultUnitText = t.GetDiastolic().GetResultUnitText()
	diastolic.ResultText = t.GetDiastolic().GetResultText(m)

	reports = append(reports, systolic)
	reports = append(reports, diastolic)

	return reports, nil
}
