package exporttypes

import (
	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
)

func NewPulseType() SimpleType {
	s := SimpleType{}

	s.npuCode = NPU_CODE_PULSE
	s.analysisText = "Hjerte—Systole; frekv. = ? × 1/min"
	s.resultUnitText = "x 1/min"
	s.decimal = 0
	s.isAlphaNumeric = false
	s.isToBeExported = true
	s.layoutResults = layoutZeroDigit
	s.devices = []MedicalDevice{NewNonin3230ExportType(), NewNonin9560ExportType(), NewAnDMedical767PBTCExportType(), NewAnDMedical767PBTCiExportType()}
	return s
}

func NewXdsPulseType() SimpleType {
	s := SimpleType{}

	s.npuCode = NPU_CODE_PULSE
	s.analysisText = "Hjerte—Systole; frekv. = ? * 1/min"
	s.resultUnitText = "1/min"
	s.decimal = 0
	s.isAlphaNumeric = false
	s.isToBeExported = true
	s.layoutResults = layoutZeroDigit
	s.devices = []MedicalDevice{NewNonin3230ExportType(), NewNonin9560ExportType(), NewAnDMedical767PBTCExportType(), NewAnDMedical767PBTCiExportType()}
	return s
}

func NewWeightType() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_WEIGHT
	s.isToBeExported = true
	s.resultUnitText = "kg"
	s.analysisText = "Pt—Legeme; masse = ? kg"
	s.decimal = 1
	s.layoutResults = layoutSingleDigit
	s.isAlphaNumeric = false
	s.devices = []MedicalDevice{NewAnDMedical321PBTCExportType(), NewAnDMedical351PBTCiExportType()}
	return s
}

func NewSaturation() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_SATURATION
	s.isToBeExported = true
	s.resultUnitText = ""
	s.analysisText = "Hb(Fe; O2-bind.; aB)—Oxygen(O2); mætn. = ?"
	s.decimal = 2
	s.layoutResults = layoutPercentDoubleDigit
	s.isAlphaNumeric = false
	s.devices = []MedicalDevice{NewNonin3230ExportType(), NewNonin9560ExportType()}
	return s
}

func NewRespiratoryRate() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_RESPIRATORY_RATE
	s.isToBeExported = true
	s.resultUnitText = "1/min"
	s.analysisText = "Pt—Respiration; frekvens = ? X 1/min"
	s.decimal = 0
	s.layoutResults = layoutZeroDigit
	s.isAlphaNumeric = false
	s.devices = []MedicalDevice{}
	return s
}

var nitriteValue = struct {
	Negative string
	Positive string
}{
	Negative: "Neg.",
	Positive: "Pos.",
}

var proteinValue = struct {
	Negative  string
	Plusminus string
	PlusOne   string
	PlusTwo   string
	PlusThree string
	PlusFour  string
}{
	Negative:  "Neg.",
	Plusminus: "+/-",
	PlusOne:   "+1",
	PlusTwo:   "+2",
	PlusThree: "+3",
	PlusFour:  "+4",
}

func NewUrineProtein() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_URINE_PROTEIN
	s.isToBeExported = true
	s.resultUnitText = ""
	s.analysisText = "U—Protein; arb.k.(proc.) = ?"
	s.decimal = 0
	s.values = []int{0, 1, 2, 3, 3}
	s.layoutResults = (func(m measurement.Measurement) string {
		value := ""
		mes := m.Measurement.Value.(string)
		switch mes {
		case proteinValue.Negative:
			value = "0"
		case proteinValue.Plusminus:
			value = "0"
		case proteinValue.PlusOne:
			value = "1"
		case proteinValue.PlusTwo:
			value = "2"
		case proteinValue.PlusThree:
			value = "3"
		case proteinValue.PlusFour:
			value = "3"

		}
		return value
	})
	s.isAlphaNumeric = false
	return s
}

func NewUrineLeukocytes() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_URINE_LEUKOCYTES
	s.isToBeExported = true
	s.resultUnitText = ""
	s.analysisText = "U—Leukocytter; arb.k.(proc.) = ?"
	s.decimal = 0
	s.values = []int{0, 1, 2, 3, 3}
	s.isAlphaNumeric = false
	s.layoutResults = (func(m measurement.Measurement) string {
		value := ""
		mes := m.Measurement.Value.(string)
		switch mes {
		case proteinValue.Negative:
			value = "0"
		case proteinValue.PlusOne:
			value = "1"
		case proteinValue.PlusTwo:
			value = "2"
		case proteinValue.PlusThree:
			value = "3"
		case proteinValue.PlusFour:
			value = "3"

		}
		return value
	})
	return s
}

func NewUrineNitrite() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_URINE_NITRITE
	s.isToBeExported = true
	s.resultUnitText = ""
	s.analysisText = "U—Nitrit; arb.k.(proc.) = ?"
	s.decimal = 0
	s.values = []int{0, 1}
	s.layoutResults = (func(m measurement.Measurement) string {
		value := ""
		mes := m.Measurement.Value.(string)
		switch mes {
		case nitriteValue.Positive:
			value = "1"
		case nitriteValue.Negative:
			value = "0"
		}
		return value
	})
	s.isAlphaNumeric = false
	return s
}
func NewUrineBlood() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_URINE_BLOOD
	s.isToBeExported = true
	s.resultUnitText = ""
	s.analysisText = "U—Nitrit; arb.k.(proc.) = ?"
	s.decimal = 0
	s.values = []int{0, 1}
	s.layoutResults = (func(m measurement.Measurement) string {
		value := ""
		mes := m.Measurement.Value.(string)
		switch mes {
		case nitriteValue.Positive:
			value = "1"
		case nitriteValue.Negative:
			value = "0"
		}
		return value
	})
	s.isAlphaNumeric = false
	return s
}

func NewUrineErythrocytes() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_URINE_ERYTHROCYTES
	s.isToBeExported = true
	s.resultUnitText = ""
	s.analysisText = "U—Erythrocytter; arb.k.(proc.) = ?"
	s.decimal = 0
	s.values = []int{0, 1, 2, 3}
	s.isAlphaNumeric = false
	s.layoutResults = (func(m measurement.Measurement) string {
		value := ""
		mes := m.Measurement.Value.(string)
		switch mes {
		case proteinValue.Negative:
			value = "0"
		case proteinValue.PlusOne:
			value = "1"
		case proteinValue.PlusTwo:
			value = "2"
		case proteinValue.PlusThree:
			value = "3"
		case proteinValue.PlusFour:
			value = "3"

		}
		return value
	})

	return s
}

// func NewUrine() simpleType {
// 	s := simpleType{}
// 	s.npuCode = "NPU3958"
// 	s.isToBeExported = false
// 	s.resultUnitText = "mg/L"
// 	s.analysisText = "P—C-reaktivt protein; massek. = ? mg/L"
// 	s.decimal = 0
// 	s.isAlphaNumeric = false

// 	return s
// }

func NewUrine() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_URINE
	s.isToBeExported = false
	s.resultUnitText = ""
	s.analysisText = "U—Protein; arb.k.(proc.) = ?"
	s.decimal = 0
	s.values = []int{0, 0, 1, 2, 3, 3}
	s.isAlphaNumeric = false

	return s
}

func NewUrineGlucose() SimpleType {
	s := SimpleType{}

	s.npuCode = NPU_CODE_URINE_GLUCOSE
	s.isToBeExported = true
	s.resultUnitText = ""
	s.analysisText = "U—Glucose; arb.k.(proc.) = ?"
	s.decimal = 0
	s.values = []int{0, 1, 2, 3, 4}
	s.isAlphaNumeric = false
	s.layoutResults = (func(m measurement.Measurement) string {
		value := ""
		mes := m.Measurement.Value.(string)
		switch mes {
		case proteinValue.Negative:
			value = "0"
		case proteinValue.PlusOne:
			value = "1"
		case proteinValue.PlusTwo:
			value = "2"
		case proteinValue.PlusThree:
			value = "3"
		case proteinValue.PlusFour:
			value = "3"

		}
		return value
	})

	return s
}

func NewTemperature() SimpleType {
	s := SimpleType{}

	s.npuCode = NPU_CODE_TEMPERATURE
	s.isToBeExported = true
	s.resultUnitText = "°C"
	s.analysisText = "Pt—Legeme; temp. = ? °C"
	s.decimal = 1
	s.layoutResults = layoutSingleDigit
	s.isAlphaNumeric = false
	s.devices = []MedicalDevice{NewAnDMedical302PBTExportType()}
	return s
}

func NewFev1Fev6() SimpleType {
	s := SimpleType{}

	s.npuCode = NPU_CODE_FEV1_FEV6_RATIO
	s.isToBeExported = false
	s.resultUnitText = ""
	s.analysisText = "Lunge—FEV1/FEV6 ratio = ?"
	s.layoutResults = layoutPercentDoubleDigit
	s.decimal = 2
	s.isAlphaNumeric = false

	return s
}

func NewFev6() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_FEV6
	s.isToBeExported = true
	s.resultUnitText = "L"
	s.analysisText = "Lunge—Lungefunktionsundersøgelse COPD FEV6; vol. = ? L"
	s.layoutResults = layoutDoubleDigit
	s.decimal = 2
	s.isAlphaNumeric = false
	return s
}

func NewFev1() SimpleType {
	s := SimpleType{}

	s.npuCode = NPU_CODE_FEV1
	s.isToBeExported = true
	s.resultUnitText = "L"
	s.analysisText = "Lunge—Lungefunktionsundersøgelse FEV1; vol. = ? L"
	// Devices=        []MedcomDevice{MedcomDevice{MedComId= "MCI00014" Manufacturer= "Vitalograph"}}
	s.layoutResults = layoutDoubleDigit
	s.decimal = 2
	s.isAlphaNumeric = false

	return s
}

func NewCrp() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_CRP
	s.isToBeExported = true
	s.resultUnitText = "mg/L"
	s.analysisText = "P—C-reaktivt protein; massek. = ? mg/L"
	s.decimal = 0
	s.layoutResults = layoutZeroDigit
	s.isAlphaNumeric = false
	return s
}

func NewBloodSugar() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_BLOODSUGAR
	s.isToBeExported = true
	s.resultUnitText = "mmol/L"
	s.analysisText = "P(kB)—Glucose; stofk. = ? mmol/L"
	s.layoutResults = layoutSingleDigit
	s.decimal = 1
	s.isAlphaNumeric = false

	return s
}
