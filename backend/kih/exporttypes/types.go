package exporttypes

import (
	"fmt"

	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
)

type MeasurementType interface {
	IsToBeExported() bool
	GetDevices() []MedicalDevice
	GetResultText(m measurement.Measurement) string
	GetAnalysisText() string
	GetResultUnitText() string
	GetNpuCode() string
}

func (s SimpleType) String() string {
	return fmt.Sprintf("[%s] E: %v Units: %s - AT: %s - digits: %d", s.npuCode, s.isToBeExported, s.resultUnitText, s.analysisText, s.decimal)
}

type SimpleType struct {
	npuCode        string
	isToBeExported bool
	resultUnitText string
	analysisText   string
	devices        []MedicalDevice
	decimal        int
	isAlphaNumeric bool
	layoutResults  func(m measurement.Measurement) string
	values         interface{}
}

// Implementation
func (s SimpleType) GetNpuCode() string {
	return s.npuCode
}

func (s SimpleType) IsToBeExported() bool {
	return s.isToBeExported
}

func (s SimpleType) GetResultUnitText() string {
	return s.resultUnitText
}
func (s SimpleType) GetAnalysisText() string {
	return s.analysisText
}
func (s SimpleType) IsAlphaNumeric() bool {
	return s.isAlphaNumeric
}
func (s SimpleType) GetDevices() []MedicalDevice {
	return s.devices
}
func (s SimpleType) GetResultText(m measurement.Measurement) string {
	if nil != s.layoutResults {
		return s.layoutResults(m)
	}
	return ""
}

type MedicalDevice interface {
	GetMedcomID() string
	GetManufacturer() string
	GetModel() string
	CheckIfSameModel(origin measurement.Origin) bool
}

// simpleDevice
type simpleDevice struct {
	medComId        string
	manufacturer    string
	model           string
	productType     string
	compareFunction func(origin measurement.Origin) bool
}

func (d simpleDevice) GetMedcomID() string {
	return d.medComId
}

func (d simpleDevice) GetManufacturer() string {
	return d.manufacturer
}

func (d simpleDevice) GetModel() string {
	return d.model
}

func (d simpleDevice) CheckIfSameModel(origin measurement.Origin) bool {
	return d.compareFunction(origin)
}

func (d simpleDevice) String() string {
	return fmt.Sprintf("[%s] model: %s - medcom id: %s", d.manufacturer, d.model, d.medComId)
}

// Setup the types to export
func init() {
	exportTypes = make(map[string]MeasurementType)
	exportTypes[TYPE_NAME_PULSE] = NewPulseType()
	exportTypes[TYPE_NAME_WEIGHT] = NewWeightType()
	exportTypes[TYPE_NAME_URINE_LEUKOCYTES] = NewUrineLeukocytes()
	exportTypes[TYPE_NAME_SATURATION] = NewSaturation()
	exportTypes[TYPE_NAME_URINE_NITRITE] = NewUrineNitrite()
	//exportTypes[TYPE_NAME_URINE_BLOOD] = NewUrineBlood()
	exportTypes[TYPE_NAME_URINE_BLOOD] = NewUrineErythrocytes()
	exportTypes[TYPE_NAME_URINE_PROTEIN] = NewUrineProtein()
	exportTypes[TYPE_NAME_URINE_GLUCOSE] = NewUrineGlucose()
	exportTypes[TYPE_NAME_URINE_ERYTHROCYTES] = NewUrineErythrocytes()
	exportTypes[TYPE_NAME_BLOODSUGAR] = NewBloodSugar()
	exportTypes[TYPE_NAME_BLOOD_PRESSURE] = NewBloodPressureType()
	exportTypes[TYPE_NAME_CRP] = NewCrp()
	exportTypes[TYPE_NAME_FEV1] = NewFev1()
	exportTypes[TYPE_NAME_FEV1_FEV6_RATIO] = NewFev1Fev6()
	exportTypes[TYPE_NAME_FEV6] = NewFev6()
	exportTypes[TYPE_NAME_TEMPERATURE] = NewTemperature()
}

var exportTypes map[string]MeasurementType

func GetExportTypes() map[string]MeasurementType {
	return exportTypes
}

// Get export types for KIH DB
func GetKihdbExportTypes() map[string]MeasurementType {
	exportTypes = make(map[string]MeasurementType)
	exportTypes[TYPE_NAME_PULSE] = NewPulseType()
	exportTypes[TYPE_NAME_WEIGHT] = NewWeightType()
	exportTypes[TYPE_NAME_URINE_LEUKOCYTES] = NewUrineLeukocytes()
	exportTypes[TYPE_NAME_SATURATION] = NewSaturation()
	exportTypes[TYPE_NAME_URINE_NITRITE] = NewUrineNitrite()
	//exportTypes[TYPE_NAME_URINE_BLOOD] = NewUrineBlood()
	exportTypes[TYPE_NAME_URINE_BLOOD] = NewUrineErythrocytes()
	exportTypes[TYPE_NAME_URINE_PROTEIN] = NewUrineProtein()
	exportTypes[TYPE_NAME_URINE_GLUCOSE] = NewUrineGlucose()
	exportTypes[TYPE_NAME_URINE_ERYTHROCYTES] = NewUrineErythrocytes()
	exportTypes[TYPE_NAME_BLOODSUGAR] = NewBloodSugar()
	exportTypes[TYPE_NAME_BLOOD_PRESSURE] = NewBloodPressureType()
	exportTypes[TYPE_NAME_CRP] = NewCrp()
	exportTypes[TYPE_NAME_FEV1] = NewFev1()
	exportTypes[TYPE_NAME_FEV1_FEV6_RATIO] = NewFev1Fev6()
	exportTypes[TYPE_NAME_FEV6] = NewFev6()
	exportTypes[TYPE_NAME_TEMPERATURE] = NewTemperature()

	return exportTypes
}

// Get export types for OIO XDS export
func GetOioXdsExportTypes() map[string]MeasurementType {
	exportTypes = make(map[string]MeasurementType)
	exportTypes[TYPE_NAME_PULSE] = NewXdsPulseType()
	exportTypes[TYPE_NAME_WEIGHT] = NewWeightType()
	exportTypes[TYPE_NAME_URINE_LEUKOCYTES] = NewUrineLeukocytes()
	exportTypes[TYPE_NAME_SATURATION] = NewSaturation()
	exportTypes[TYPE_NAME_URINE_NITRITE] = NewUrineNitrite()
	exportTypes[TYPE_NAME_URINE_BLOOD] = NewUrineErythrocytes()
	exportTypes[TYPE_NAME_URINE_PROTEIN] = NewUrineProtein()
	exportTypes[TYPE_NAME_URINE_GLUCOSE] = NewUrineGlucose()
	exportTypes[TYPE_NAME_URINE_ERYTHROCYTES] = NewUrineErythrocytes()
	exportTypes[TYPE_NAME_BLOODSUGAR] = NewBloodSugar()
	exportTypes[TYPE_NAME_BLOOD_PRESSURE] = NewBloodPressureType()
	exportTypes[TYPE_NAME_CRP] = NewCrp()
	exportTypes[TYPE_NAME_FEV1] = NewFev1()
	exportTypes[TYPE_NAME_FEV1_FEV6_RATIO] = NewFev1Fev6()
	exportTypes[TYPE_NAME_FEV6] = NewFev6()
	exportTypes[TYPE_NAME_TEMPERATURE] = NewTemperature()
	exportTypes[TYPE_NAME_RESPIRATORY_RATE] = NewRespiratoryRate()

	return exportTypes
}

// Compiled handled names for measurement types
const (
	// vital signs
	NPU_CODE_SATURATION               = "NPU03011"
	NPU_CODE_BLOOD_PRESSURE_SYSTOLIC  = "DNK05472"
	NPU_CODE_BLOOD_PRESSURE_DIASTOLIC = "DNK05473"
	NPU_CODE_TEMPERATURE              = "NPU08676"
	NPU_CODE_PULSE                    = "NPU21692"
	// results
	NPU_CODE_WEIGHT             = "NPU03804"
	NPU_CODE_URINE              = "NPU04206"
	NPU_CODE_URINE_BLOOD        = "NPU04206"
	NPU_CODE_URINE_GLUCOSE      = "NPU04207"
	NPU_CODE_URINE_NITRITE      = "NPU21578"
	NPU_CODE_URINE_PROTEIN      = "NPU04206"
	NPU_CODE_URINE_LEUKOCYTES   = "NPU03987"
	NPU_CODE_URINE_ERYTHROCYTES = "NPU03963"
	NPU_CODE_BLOODSUGAR         = "NPU22089"
	NPU_CODE_CRP                = "NPU19748"
	NPU_CODE_FEV1               = "MCS88015"
	NPU_CODE_FEV6               = "MCS88100"
	NPU_CODE_FEV1_FEV6_RATIO    = "MCS88099"
	NPU_CODE_RESPIRATORY_RATE   = "MCS88122"
	// NPU_CODE_CTG                      = ""
	// NPU_CODE_ECG                      = ""
	// NPU_CODE_FEF25_75                 = ""
	// NPU_CODE_FEV1_PERCENTAGE          = ""
	// NPU_CODE_FEV6_PERCENTAGE          = ""
	// NPU_CODE_HEIGHT                   = ""
	// NPU_CODE_HEMOGLOBIN               = ""
	// NPU_CODE_LEAK_50                  = ""
	// NPU_CODE_LEAK_95                  = ""
	// NPU_CODE_PAIN_SCALE               = ""
	// NPU_CODE_PEAK_FLOW                = ""
	// NPU_CODE_RESPIRATORY_RATE         = "respiratory_rate"
	// NPU_CODE_RESPIRATORY_RATE_50      = "respiratory_rate_50%"
	// NPU_CODE_RESPIRATORY_RATE_95      = "respiratory_rate_95%"
	// NPU_CODE_SIT_TO_STAND             = "sit_to_stand"
	// NPU_CODE_SATURATION_50            = "saturation_50%"
	// NPU_CODE_SATURATION_95            = "saturation_95%"
	// NPU_CODE_SPIROMETER               = "spirometry"
	// NPU_CODE_TIDAL_VOLUME_50          = "tidal_volume_50%"
	// NPU_CODE_TIDAL_VOLUME_95          = "tidal_volume_95%"
)

const (
	TYPE_NAME_WEIGHT                             = "weight"
	TYPE_NAME_URINE_BLOOD                        = "blood_in_urine"
	TYPE_NAME_URINE                              = "protein_in_urine"
	TYPE_NAME_URINE_GLUCOSE                      = "glucose_in_urine"
	TYPE_NAME_URINE_NITRITE                      = "nitrite_in_urine"
	TYPE_NAME_URINE_PROTEIN                      = "protein_in_urine"
	TYPE_NAME_URINE_COMBI                        = "urine_measurement"
	TYPE_NAME_URINE_LEUKOCYTES                   = "leukocytes_in_urine"
	TYPE_NAME_URINE_ERYTHROCYTES                 = "erythrocytes_in_urine"
	TYPE_NAME_ACTIVITY                           = "activity"
	TYPE_NAME_BLOOD_PRESSURE                     = "blood_pressure"
	TYPE_NAME_BLOOD_PRESSURE_SYSTOLIC            = "blood_pressure_systolic"
	TYPE_NAME_BLOOD_PRESSURE_DIASTOLIC           = "blood_pressure_diastolic"
	TYPE_NAME_BLOODSUGAR                         = "bloodsugar"
	TYPE_NAME_CONTINUOUS_BLOOD_SUGAR_MEASUREMENT = "continuous_blood_sugar_measurement"
	TYPE_NAME_CRP                                = "crp"
	TYPE_NAME_CTG                                = "ctg"
	TYPE_NAME_ECG                                = "ecg"
	TYPE_NAME_DAILY_STEPS                        = "daily_steps"
	TYPE_NAME_DAILY_STEPS_WEEKLY_AVERAGE         = "daily_steps_weekly_average"
	TYPE_NAME_DURATION                           = "duration"
	TYPE_NAME_DURATION_HOURS                     = "duration_hours"
	TYPE_NAME_FEV1                               = "fev1"
	TYPE_NAME_FEV6                               = "fev6"
	TYPE_NAME_FEV1_FEV6_RATIO                    = "fev1/fev6"
	TYPE_NAME_FEF25_75                           = "fef25-75%"
	TYPE_NAME_FEV1_PERCENTAGE                    = "fev1%"
	TYPE_NAME_FEV6_PERCENTAGE                    = "fev6%"
	TYPE_NAME_HEIGHT                             = "height"
	TYPE_NAME_HEMOGLOBIN                         = "hemoglobin"
	TYPE_NAME_LEAK_50                            = "leak_50%"
	TYPE_NAME_LEAK_95                            = "leak_95%"
	TYPE_NAME_PAIN_SCALE                         = "pain_scale"
	TYPE_NAME_PEAK_FLOW                          = "peak_flow"
	TYPE_NAME_PULSE                              = "pulse"
	TYPE_NAME_RESPIRATORY_RATE                   = "respiratory_rate"
	TYPE_NAME_RESPIRATORY_RATE_50                = "respiratory_rate_50%"
	TYPE_NAME_RESPIRATORY_RATE_95                = "respiratory_rate_95%"
	TYPE_NAME_SIT_TO_STAND                       = "sit_to_stand"
	TYPE_NAME_SATURATION                         = "saturation"
	TYPE_NAME_SATURATION_50                      = "saturation_50%"
	TYPE_NAME_SATURATION_95                      = "saturation_95%"
	TYPE_NAME_SPIROMETER                         = "spirometry"
	TYPE_NAME_TEMPERATURE                        = "temperature"
	TYPE_NAME_TIDAL_VOLUME_50                    = "tidal_volume_50%"
	TYPE_NAME_TIDAL_VOLUME_95                    = "tidal_volume_95%"
)

const (
	PRODUCT_TYPE_LUNGMONITOR          = "Lung Monitor"
	PRODUCT_TYPE_PULSEOXIMETER        = "Pulse Oximeter"
	PRODUCT_TYPE_BLOODPRESSUREMONITOR = "Blood Pressure Monitor"
	PRODUCT_TYPE_URINEANALYZER        = "Urine Analyzer"
	PRODUCT_TYPE_WEIGHT               = "Weight"
	PRODUCT_TYPE_THERMOMETER          = "Thermometer"
)
