package exporttypes

import (
	"fmt"
	"strconv"

	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
)

type BloodPressureType struct {
	systolic       SimpleType
	diastolic      SimpleType
	isToBeExported bool
}

func (b BloodPressureType) GetDevices() []MedicalDevice {
	res := []MedicalDevice{}

	res = append(res, b.systolic.GetDevices()...)
	res = append(res, b.diastolic.GetDevices()...)
	return res
}

func (b BloodPressureType) IsToBeExported() bool {
	return true
}
func (b BloodPressureType) GetAnalysisText() string {
	return ""
}
func (b BloodPressureType) GetResultUnitText() string {
	return ""
}
func (b BloodPressureType) GetResultText(m measurement.Measurement) string {
	return ""
}
func (b BloodPressureType) GetNpuCode() string {
	return fmt.Sprintf("%s,%s", b.systolic.GetNpuCode(), b.diastolic.GetNpuCode())
}

func (b BloodPressureType) GetSystolic() SimpleType {
	return b.systolic
}

func (b BloodPressureType) GetDiastolic() SimpleType {
	return b.diastolic
}

func NewBloodPressureType() BloodPressureType {
	bp := BloodPressureType{systolic: newSystolicBloodPressure(), diastolic: newDiastolicBloodPressure()}

	bp.isToBeExported = true
	return bp
}
func newSystolicBloodPressure() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_BLOOD_PRESSURE_SYSTOLIC
	s.isToBeExported = true
	s.resultUnitText = "mmHg"
	s.analysisText = "Arm—Blodtryk(systolisk); tryk = ? mmHg"
	s.decimal = 0
	s.layoutResults = func(m measurement.Measurement) string {
		return strconv.FormatFloat(float64(m.Measurement.Systolic), 'f', 0, 64)
	}

	s.isAlphaNumeric = false
	s.devices = []MedicalDevice{NewAnDMedical767PBTCExportType(), NewAnDMedical767PBTCiExportType()}
	return s
}

func newDiastolicBloodPressure() SimpleType {
	s := SimpleType{}
	s.npuCode = NPU_CODE_BLOOD_PRESSURE_DIASTOLIC
	s.isToBeExported = true
	s.resultUnitText = "mmHg"
	s.analysisText = "Arm—Blodtryk(diastolisk); tryk = ? mmHg"
	s.devices = []MedicalDevice{NewAnDMedical767PBTCExportType(), NewAnDMedical767PBTCiExportType()}
	s.layoutResults = func(m measurement.Measurement) string {
		return strconv.FormatFloat(float64(m.Measurement.Diastolic), 'f', 0, 64)
	}
	s.decimal = 0
	s.isAlphaNumeric = false

	return s
}
