package shared

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"
	"time"

	"github.com/KvalitetsIT/kih-telecare-exporter/app"
	"github.com/KvalitetsIT/kih-telecare-exporter/backend/kih/exporttypes"
	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
	"github.com/KvalitetsIT/kih-telecare-exporter/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type fields struct {
}

type args struct {
	m measurement.Measurement
}

func init() {
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)

	var err error
	config, err = app.InitConfig()
	if err != nil {
		log.Fatalf("Error creating config %v", err)
	}
	fmt.Println("Here")
}

var typetests = []struct {
	name     string
	fields   fields
	args     args
	npu      string
	exported bool
}{
	{"FEV1", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_FEV1}}, "MCS88015", true},
	{"Weight", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_WEIGHT}}, "NPU03804", true},
	{"Temperature", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_TEMPERATURE}}, "NPU08676", true},
	{"Saturation", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_SATURATION}}, "NPU03011", true},
	{"Blood Pressure", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_BLOOD_PRESSURE}}, "DNK05472,DNK05473", true},
	//{"Blood Pressure Diastolic", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_BLOOD_PRESSURE_DIASTOLIC}}, "DNK05473", true},
	{"Pulse", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_PULSE}}, "NPU21692", true},
	{"Blood sugar", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_BLOODSUGAR}}, "NPU22089", true},
	{"Urine Glusode", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_URINE_GLUCOSE}}, "NPU04207", true},
	{"Urine Nitrite", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_URINE_NITRITE}}, "NPU21578", true},
	{"Urine Leukocytes", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_URINE_LEUKOCYTES}}, "NPU03987", true},
	{"CRP", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_CRP}}, "NPU19748", true},
	{"Urine Erythrocytes", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_URINE_ERYTHROCYTES}}, "NPU03963", true},
	{"Fev Fev 6 Ratio", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_FEV1_FEV6_RATIO}}, "MCS88099", false},
}

func TestPopulateInstrumentFromMeasurement(t *testing.T) {

	tests := []struct {
		name  string
		file  string
		isNil bool
	}{
		{"Pulse nonin Onyx II", "testdata/pulse-mci00005.json", false},
		{"Pulse Nonin 3230 Bluetooth Smart Pulse Oximeter", "testdata/pulse-mci00013.json", false},
		{"Not mapped", "testdata/weight.json", true},
		{"MyGlucohealth", "testdata/blood_pressure.json", true},
	}

	for _, tt := range tests {
		data_file, err := ioutil.ReadFile(tt.file)
		if err != nil {
			t.Errorf("Error reading input file - %v", err)
		}

		var mm measurement.Measurement
		if err := json.Unmarshal(data_file, &mm); err != nil {
			t.Errorf("Error parsing json - %v", err)
		}

		log.Debugf("Origin %s", mm.Origin)
		i := mapMeasurmentToInstrument(mm)

		log.Debug("Got ", i)
		if nil == i {
			t.Errorf("Instrument is nil %v", i)
		}
		if len(i.SoftwareVersion) == 0 {
			t.Errorf("Software version is not set %v", i.SoftwareVersion)
		}
	}

}

func TestPopulateInstrument(t *testing.T) {

	tests := []struct {
		name       string
		file       string
		exportType exporttypes.MeasurementType
		medComID   string
		isNil      bool
	}{
		{"Pulse nonin Onyx II", "testdata/pulse-mci00005.json", exporttypes.NewPulseType(), "MCI00005", false},
		{"Pulse Nonin 3230 Bluetooth Smart Pulse Oximeter", "testdata/pulse-mci00013.json", exporttypes.NewPulseType(), "MCI00013", false},
		{"Not mapped", "testdata/weight.json", exporttypes.SimpleType{}, "", true},
	}

	for _, tt := range tests {
		data_file, err := ioutil.ReadFile(tt.file)
		if err != nil {
			t.Errorf("Error reading input file - %v", err)
		}

		var mm measurement.Measurement
		if err := json.Unmarshal(data_file, &mm); err != nil {
			t.Errorf("Error parsing json - %v", err)
		}

		i := measurmentToInstrument(tt.exportType, mm)

		log.Debug("Got ", i)
		if tt.isNil && i != nil {
			t.Errorf("Instrument should be nil ")
		}

		if !tt.isNil && tt.medComID != i.MedComID {
			t.Errorf("MedCom id should be %s - received %v", tt.medComID, i.MedComID)
		}
	}

}

func TestExportReportFromMeasurement(t *testing.T) {
	weight_file, err := ioutil.ReadFile("../testdata/weight.json")
	if err != nil {
		t.Errorf("Error reading input file - %v", err)
	}

	var mm measurement.Measurement
	if err := json.Unmarshal(weight_file, &mm); err != nil {
		t.Errorf("Error parsing json - %v", err)
	}

	fmt.Println("m", mm)
	log.Debug("M", mm)

	mt, err := measurementToMeasurementType(mm)
	if err != nil {
		t.Errorf("Error - %v", err)
	}
	var lr []LaboratoryReportExtended

	lr, err = ReportFromMeasurement(exporttypes.GetKihdbExportTypes(), mm, mt)
	fmt.Printf("Got - %+v\n", lr)
	if err != nil {
		t.Errorf("Error parsing measurement %v", err)
	}
	if len(lr) > 0 {
		simpleValidation(t, exporttypes.GetKihdbExportTypes(), mm, mt, lr[0])
		validateProducerOflabResults(t, lr[0])

	} else {
		t.Error("No reports returned - should be 1")
	}

}

func measurementToMeasurementType(measurement measurement.Measurement) (repository.MeasurementExportState, error) {
	m := repository.MeasurementExportState{}
	log.Debug("Handling measuremnt - ", measurement.Links.Measurement)

	m.Measurement = measurement.Links.Measurement
	m.Status = repository.CREATED
	newUuid, err := uuid.NewUUID()
	if err != nil {
		errStr := fmt.Sprintf("Error generating UUID - %v", err)
		return m, errors.Wrap(err, errStr)
	}
	m.ID = newUuid

	return m, nil
}

// validate common values
func simpleValidation(t *testing.T, exportedTypes map[string]exporttypes.MeasurementType, mm measurement.Measurement, mt repository.MeasurementExportState, lr LaboratoryReportExtended) {
	exportType := exportedTypes[mm.Type]

	if lr.UuidIdentifier != mt.ID.String() {
		t.Errorf("Expected '%s' got '%s'", mt.ID, lr.UuidIdentifier)
	}
	if lr.ResultText != strconv.FormatFloat(float64(mm.Measurement.Value.(float64)), 'f', 1, 64) {
		t.Errorf("Expected %f - got %s", mm.Measurement.Value, lr.ResultText)
	}

	if lr.IupacIdentifier != exportType.GetNpuCode() {
		t.Errorf("Expected %s - got %s", exportType.GetNpuCode(), lr.IupacIdentifier)
	}

	if lr.ResultUnitText != "kg" {
		t.Errorf("Expected %s - got %s", exportType.GetResultUnitText(), lr.ResultUnitText)
	}

	// if lr.AnalysisText != exportType.AnalysisText {
	// 	t.Errorf("Expected %s - got %s", exportedTypes[mm.Type].AnalysisText, lr.AnalysisText)
	// }

	if lr.ResultTypeOfInterval != "unspecified" {
		t.Errorf("Expected unspecified - got %s", lr.ResultTypeOfInterval)
	}
	if lr.NationalSampleIdentifier != "9999999999" {
		t.Errorf("Expected 9999999999 - got %s", lr.NationalSampleIdentifier)
	}

	if lr.CreatedDateTime != mm.Timestamp.Format(time.RFC3339) {
		t.Errorf("Expected %s - got %s", mm.Timestamp.Format(time.RFC3339), lr.CreatedDateTime)
	}
}

func validateProducerOflabResults(t *testing.T, lr LaboratoryReportExtended) {
	if "POT" != lr.ProducerOfLabResult.IdentifierCode {
		t.Errorf("Expected Producer of lab result - POT - got %s", lr.ProducerOfLabResult.IdentifierCode)
	}
	if "Patient målt" != lr.ProducerOfLabResult.Identifier {
		t.Errorf("Expected Producer of lab result - Patient målt  - got %s", lr.ProducerOfLabResult.Identifier)
	}
}

func TestHandleOrigin(t *testing.T) {
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)

	tests := []struct {
		testfile     string
		transferedby string
		medcomid     string
		exporttype   exporttypes.MeasurementType
	}{
		{"weight", "automatic", "", exporttypes.NewWeightType()},
		{"weight_typed", "typed", "", exporttypes.NewWeightType()},
		{"weight_no_origin", "typed", "", exporttypes.NewWeightType()},
		{"weight_manual", "typedbyhcprof", "", exporttypes.NewWeightType()},
		{"blood_pressure", "automatic", "", exporttypes.NewBloodPressureType()},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Input - %s - outcome - %s", tt.testfile, tt.transferedby), func(t *testing.T) {
			f, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s.json", tt.testfile))
			if err != nil {
				t.Errorf("Error reading input file - %v", err)
			}

			var mm measurement.Measurement
			if err := json.Unmarshal(f, &mm); err != nil {
				t.Errorf("Error parsing json - %v", err)
			}

			lr := LaboratoryReportExtended{}

			if err := handleOrigin(tt.exporttype, mm, &lr); err != nil {
				t.Errorf("Error parsing origin - %v", err)
			}

			if tt.transferedby != lr.MeasurementTransferredBy {
				fmt.Printf("Got %+v", lr)
				t.Errorf("Expected typed got %s", lr.MeasurementTransferredBy)
			}

			if nil != lr.Instrument && tt.medcomid != lr.Instrument.MedComID {
				fmt.Printf("Got %+v", lr)
				t.Errorf("Instrument - id Expected '%s' got '%s'", tt.medcomid, lr.Instrument.MedComID)

			}
		})
	}
}

func TestReportFromMeasurementSimpleTypes(t *testing.T) {

	now := time.Now()

	exportedTypes := exporttypes.GetKihdbExportTypes()

	typetests := []struct {
		name         string
		measurement  measurement.Measurement
		npu          string
		exported     bool
		value        string
		unit         string
		analysisText string
	}{
		{"FEV1",
			measurement.Measurement{Type: exporttypes.TYPE_NAME_FEV1,
				Timestamp:   now,
				Measurement: measurement.MeasurementValue{Unit: "%", Value: 4.98}},
			"MCS88015", true, "4.98", "L", "Lunge—Lungefunktionsundersøgelse FEV1; vol. = ? L"},
		{"Weight",
			measurement.Measurement{Type: exporttypes.TYPE_NAME_WEIGHT,
				Timestamp:   now,
				Measurement: measurement.MeasurementValue{Unit: "kg", Value: 98}},
			"NPU03804", true, "98.0", "kg", "Pt—Legeme; masse = ? kg"},
		{"Temperature",
			measurement.Measurement{Type: exporttypes.TYPE_NAME_TEMPERATURE,
				Timestamp:   now,
				Measurement: measurement.MeasurementValue{Unit: "M", Value: 37.0}},
			"NPU08676", true, "37.0", "°C", "Pt—Legeme; temp. = ? °C"},
		{"Saturation",
			measurement.Measurement{Type: exporttypes.TYPE_NAME_SATURATION,
				Timestamp:   now,
				Measurement: measurement.MeasurementValue{Unit: "%", Value: 98}},
			"NPU03011", true, "0.98", "", "Hb(Fe; O2-bind.; aB)—Oxygen(O2); mætn. = ?"},
		{"Pulse",
			measurement.Measurement{Type: exporttypes.TYPE_NAME_PULSE,
				Timestamp:   now,
				Measurement: measurement.MeasurementValue{Unit: "BPM", Value: 58}},
			"NPU21692", true, "58", "x 1/min", "Hjerte—Systole; frekv. = ? × 1/min"},
		{"BloodSugar",
			measurement.Measurement{Type: exporttypes.TYPE_NAME_BLOODSUGAR,
				Timestamp:   now,
				Measurement: measurement.MeasurementValue{Unit: "%", Value: 8.4}},
			"NPU22089", true, "8.4", "mmol/L", "P(kB)—Glucose; stofk. = ? mmol/L"},
		// {"Urine Glucose",
		// 	measurement.Measurement{Type: exporttypes.TYPE_NAME_URINE_GLUCOSE,
		// 		Timestamp:   now,
		// 		Measurement: measurement.MeasurementValue{Unit: "%", Value: "+2"}},
		// 	"NPU04207", true, "28", "mg/L", "P—C-reaktivt protein; massek. = ? mg/l"},
		// {"Urine Nitrite",
		// 	measurement.Measurement{Type: exporttypes.TYPE_NAME_URINE_NITRITE,
		// 		Timestamp:   now,
		// 		Measurement: measurement.MeasurementValue{Unit: "%", Value: "+2"}},
		// 	"NPU21578", true, "28", "mg/L", "P—C-reaktivt protein; massek. = ? mg/L"},
		// {"Urine Leukocytes",
		// 	measurement.Measurement{Type: exporttypes.TYPE_NAME_URINE_LEUKOCYTES,
		// 		Timestamp:   now,
		// 		Measurement: measurement.MeasurementValue{Unit: "%", Value: "+2"}},
		// 	"NPU03978", true, "28", "mg/L", "P—C-reaktivt protein; massek. = ? mg/L"},
		{"CRP",
			measurement.Measurement{Type: exporttypes.TYPE_NAME_CRP,
				Timestamp:   now,
				Measurement: measurement.MeasurementValue{Unit: "%", Value: 28}},
			"NPU19748", true, "28", "mg/L", "P—C-reaktivt protein; massek. = ? mg/L"},
		// {"Urine Erythrocytes",
		// 	measurement.Measurement{Type: exporttypes.TYPE_NAME_URINE_ERYTHROCYTES,
		// 		Timestamp:   now,
		// 		Measurement: measurement.MeasurementValue{Unit: "%", Value: 28}},
		// 	"NPU03963", true, "28", "mg/L", "P—C-reaktivt protein; massek. = ? mg/L"},
		// {"Blood Pressure", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_BLOOD_PRESSURE}}, "DNK05472,DNK05473", true},
		// //{"Blood Pressure Diastolic", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_BLOOD_PRESSURE_DIASTOLIC}}, "DNK05473", true},
		// {"Blood sugar", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_BLOODSUGAR}}, "NPU22089", true},
		// {"Urine Glusode", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_URINE_GLUCOSE}}, "NPU04207", true},
		// {"Urine Nitrite", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_URINE_NITRITE}}, "NPU21578", true},
		// {"Urine Leukocytes", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_URINE_LEUKOCYTES}}, "NPU03987", true},
		// {"Urine Erythrocytes", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_URINE_ERYTHROCYTES}}, "NPU03963", true},
		// {"Fev Fev 6 Ratio", fields{}, args{m: measurement.Measurement{Type: exporttypes.TYPE_NAME_FEV1_FEV6_RATIO}}, "MCS88099", false},
	}
	for _, tt := range typetests {
		t.Run(tt.name, func(t *testing.T) {
			mr := repository.MeasurementExportState{}
			mr.ID = uuid.New()
			reports, err := ReportFromMeasurement(exportedTypes, tt.measurement, mr)
			if err != nil {
				t.Error("Error converting measurement ", err)
			}

			fmt.Printf("Received : \n%+v\n", reports)
			if len(reports) != 1 {
				t.Error("Error converting measurement - expected more results ", reports)
			}
			report := reports[0]
			if report.ResultText != tt.value {
				t.Errorf("Expected '%s' got '%s'", tt.value, report.ResultText)
			}

			if report.ResultUnitText != tt.unit {
				t.Errorf("Expected '%s' got '%s'", tt.unit, report.ResultUnitText)
			}

			if report.IupacIdentifier != tt.npu {
				t.Errorf("Expected '%s' got '%s'", tt.npu, report.IupacIdentifier)
			}
			if report.AnalysisText != tt.analysisText {
				t.Errorf("Expected '%s' got '%s'", tt.analysisText, report.AnalysisText)
			}
			if mr.ID.String() != report.UuidIdentifier {
				t.Errorf("Expected '%s' got '%s'", mr.ID.String(), report.UuidIdentifier)
			}
			if now.Format(time.RFC3339) != report.CreatedDateTime {
				t.Errorf("Expected '%s' got '%s'", now.String(), report.CreatedDateTime)

			}
		})
	}
}

func TestReportFromMeasurementSimpleTypesFromFile(t *testing.T) {
	now := time.Now()
	typetests := []struct {
		name         string
		datafile     string
		shouldFail   bool
		npu          string
		exported     bool
		value        string
		unit         string
		analysisText string
	}{
		{"Temperature", "testdata/temperature.json", false, "NPU08676", true, "36.4", "°C", "Pt—Legeme; temp. = ? °C"},
		{"Saturation", "testdata/saturation.json", false, "NPU03011", true, "0.96", "", "Hb(Fe; O2-bind.; aB)—Oxygen(O2); mætn. = ?"},
		{"Pulse", "testdata/pulse.json", false, "NPU21692", true, "58", "x 1/min", "Hjerte—Systole; frekv. = ? × 1/min"},
		{"FEV1 / FEV6", "testdata/fev1-fev6-ratio.json", false, "MCS88099", true, "0.80", "", "Lunge—FEV1/FEV6 ratio = ?"},
		{"Fef25-75", "testdata/fef25.json", true, "MCS88100", true, "5.00", "L", "Lunge—Lungefunktionsundersøgelse COPD FEV6; vol. = ? L"},
		{"Fev6", "testdata/fev6.json", false, "MCS88100", true, "5.00", "L", "Lunge—Lungefunktionsundersøgelse COPD FEV6; vol. = ? L"},
		{"Fev1", "testdata/fev1.json", false, "MCS88015", true, "4.00", "L", "Lunge—Lungefunktionsundersøgelse FEV1; vol. = ? L"},
		// Weight
		{"Weight", "testdata/weight.json", false, "NPU03804", true, "84.9", "kg", "Pt—Legeme; masse = ? kg"},
		// Urine Nitrite
		{"Urine Nitrite Positive", "testdata/nitrite_in_urine_positive.json", false, "NPU21578", true, "1", "", "U—Nitrit; arb.k.(proc.) = ?"},
		{"Urine Nitrite Negative", "testdata/nitrite_in_urine_negative.json", false, "NPU21578", true, "0", "", "U—Nitrit; arb.k.(proc.) = ?"},
		// Urine Protein
		{"Urine Protein Negative", "testdata/protein_in_urine_negative.json", false, "NPU04206", true, "0", "", "U—Protein; arb.k.(proc.) = ?"},
		{"Urine Protein Plus Minus", "testdata/protein_in_urine_plus_minus.json", false, "NPU04206", true, "0", "", "U—Protein; arb.k.(proc.) = ?"},
		{"Urine Protein Plus One", "testdata/protein_in_urine_plus_one.json", false, "NPU04206", true, "1", "", "U—Protein; arb.k.(proc.) = ?"},
		{"Urine Protein Plus Two", "testdata/protein_in_urine_plus_two.json", false, "NPU04206", true, "2", "", "U—Protein; arb.k.(proc.) = ?"},
		{"Urine Protein Plus Three", "testdata/protein_in_urine_plus_three.json", false, "NPU04206", true, "3", "", "U—Protein; arb.k.(proc.) = ?"},
		{"Urine Protein Plus Four", "testdata/protein_in_urine_plus_four.json", false, "NPU04206", true, "3", "", "U—Protein; arb.k.(proc.) = ?"},
		// Leukocytes
		{"Urine Leukocytes Negative", "testdata/leukocytes_in_urine_negative.json", false, "NPU03987", true, "0", "", "U—Leukocytter; arb.k.(proc.) = ?"},
		{"Urine Leukocytes Plus One", "testdata/leukocytes_in_urine_plus_one.json", false, "NPU03987", true, "1", "", "U—Leukocytter; arb.k.(proc.) = ?"},
		{"Urine Leukocytes Plus Two", "testdata/leukocytes_in_urine_plus_two.json", false, "NPU03987", true, "2", "", "U—Leukocytter; arb.k.(proc.) = ?"},
		{"Urine Leukocytes Plus Three", "testdata/leukocytes_in_urine_plus_three.json", false, "NPU03987", true, "3", "", "U—Leukocytter; arb.k.(proc.) = ?"},
		{"Urine Leukocytes Plus Four", "testdata/leukocytes_in_urine_plus_four.json", false, "NPU03987", true, "3", "", "U—Leukocytter; arb.k.(proc.) = ?"},
		// blood_urine
		{"Urine Blood Negative", "testdata/blood_in_urine_negative.json", false, "NPU03963", true, "0", "", "U—Erythrocytter; arb.k.(proc.) = ?"},
		{"Urine Blood Plus One", "testdata/blood_in_urine_plus_one.json", false, "NPU03963", true, "1", "", "U—Erythrocytter; arb.k.(proc.) = ?"},
		{"Urine blood Plus Two", "testdata/blood_in_urine_plus_two.json", false, "NPU03963", true, "2", "", "U—Erythrocytter; arb.k.(proc.) = ?"},
		{"Urine blood Plus Three", "testdata/blood_in_urine_plus_three.json", false, "NPU03963", true, "3", "", "U—Erythrocytter; arb.k.(proc.) = ?"},
		{"Urine blood Plus Four", "testdata/blood_in_urine_plus_four.json", false, "NPU03963", true, "3", "", "U—Erythrocytter; arb.k.(proc.) = ?"},
		// Urine Glucose
		{"Urine Glucose Negative", "testdata/glucose_in_urine_negative.json", false, "NPU04207", true, "0", "", "U—Glucose; arb.k.(proc.) = ?"},
		{"Urine Glucose Plus One", "testdata/glucose_in_urine_plus_one.json", false, "NPU04207", true, "1", "", "U—Glucose; arb.k.(proc.) = ?"},
		{"Urine Glucose Plus Two", "testdata/glucose_in_urine_plus_two.json", false, "NPU04207", true, "2", "", "U—Glucose; arb.k.(proc.) = ?"},
		{"Urine Glucose Plus Three", "testdata/glucose_in_urine_plus_three.json", false, "NPU04207", true, "3", "", "U—Glucose; arb.k.(proc.) = ?"},
		{"Urine Glucose Plus Four", "testdata/glucose_in_urine_plus_four.json", false, "NPU04207", true, "3", "", "U—Glucose; arb.k.(proc.) = ?"},
	}
	for _, tt := range typetests {
		t.Run(tt.name, func(t *testing.T) {
			mr := repository.MeasurementExportState{}
			mr.ID = uuid.New()
			inputdata, err := ioutil.ReadFile(tt.datafile)
			if err != nil {
				t.Errorf("Error reading file %v", err)
				t.SkipNow()
			}
			var m measurement.Measurement
			if err := json.Unmarshal(inputdata, &m); err != nil {
				t.Errorf("Error converting measurement - %v", err)
			}
			m.Timestamp = now
			fmt.Println("M ", m)
			reports, err := ReportFromMeasurement(exporttypes.GetKihdbExportTypes(), m, mr)
			if err != nil {
				if tt.shouldFail {
					t.SkipNow()
				}
				t.Error("Error converting measurement ", err)
			} else {
				if tt.shouldFail {
					t.Error("Test should fail, but didnt")
				}
			}

			fmt.Printf("Received : \n%+v\n", reports)
			if len(reports) != 1 {
				t.Error("Error converting measurement - expected more results ", reports)
			}
			report := reports[0]
			if report.ResultText != tt.value {
				t.Errorf("Expected '%s' got '%s'", tt.value, report.ResultText)
			}

			if report.ResultUnitText != tt.unit {
				t.Errorf("Expected '%s' got '%s'", tt.unit, report.ResultUnitText)
			}

			if report.IupacIdentifier != tt.npu {
				t.Errorf("Expected '%s' got '%s'", tt.npu, report.IupacIdentifier)
			}
			if report.AnalysisText != tt.analysisText {
				t.Errorf("Expected '%s' got '%s'", tt.analysisText, report.AnalysisText)
			}
			if mr.ID.String() != report.UuidIdentifier {
				t.Errorf("Expected '%s' got '%s'", mr.ID.String(), report.UuidIdentifier)
			}
			if now.Format(time.RFC3339) != report.CreatedDateTime {
				t.Errorf("Expected '%s' got '%s'", now.String(), report.CreatedDateTime)

			}
		})
	}
}

func TestReportFromMeasurementComplexTypesFromFile(t *testing.T) {
	now := time.Now()
	typetests := []struct {
		name         string
		datafile     string
		shouldFail   bool
		npu          []string
		exported     bool
		value        []string
		unit         []string
		analysisText []string
	}{
		{"Blood Pressure", "testdata/blood_pressure.json", false, []string{"DNK05472", "DNK05473"}, true, []string{"130", "80"}, []string{"mmHg", "mmHg"}, []string{"Arm—Blodtryk(systolisk); tryk = ? mmHg", "Arm—Blodtryk(diastolisk); tryk = ? mmHg"}},
	}
	for _, tt := range typetests {
		t.Run(tt.name, func(t *testing.T) {
			mr := repository.MeasurementExportState{}
			mr.ID = uuid.New()
			inputdata, err := ioutil.ReadFile(tt.datafile)
			if err != nil {
				t.Errorf("Error reading file %v", err)
				t.SkipNow()
			}
			var m measurement.Measurement
			if err := json.Unmarshal(inputdata, &m); err != nil {
				t.Errorf("Error converting measurement - %v", err)
			}
			m.Timestamp = now
			fmt.Println("M ", m)
			reports, err := ReportFromMeasurement(exporttypes.GetKihdbExportTypes(), m, mr)
			if err != nil {
				if tt.shouldFail {
					t.SkipNow()
				}
				t.Error("Error converting measurement ", err)
			} else {
				if tt.shouldFail {
					t.Error("Test should fail, but didnt")
				}
			}

			fmt.Printf("Received : \n%+v\n", reports)
			if len(reports) != len(tt.npu) {
				t.Errorf("Error converting measurement - got %d expected %d", len(tt.npu), len(reports))
			}
			for i, v := range reports {
				report := v

				if report.ResultUnitText != tt.unit[i] {
					t.Errorf("Expected '%s' got '%s'", tt.unit[i], report.ResultUnitText)
				}

				if report.IupacIdentifier != tt.npu[i] {
					t.Errorf("Expected '%s' got '%s'", tt.npu[i], report.IupacIdentifier)
				}
				if report.AnalysisText != tt.analysisText[i] {
					t.Errorf("Expected '%s' got '%s'", tt.analysisText[i], report.AnalysisText)
				}
				if now.Format(time.RFC3339) != report.CreatedDateTime {
					t.Errorf("Expected '%s' got '%s'", now.String(), report.CreatedDateTime)

				}
			}

		})
	}
}
