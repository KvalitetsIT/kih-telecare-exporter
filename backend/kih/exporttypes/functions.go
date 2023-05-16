package exporttypes

import (
	"reflect"
	"strconv"

	"bitbucket.org/opentelehealth/exporter/measurement"
	"github.com/sirupsen/logrus"
)

func layoutZeroDigit(m measurement.Measurement) string {
	return strconv.FormatFloat(handleConversionToFloat(m.Measurement.Value), 'f', 0, 64)
}

func layoutSingleDigit(m measurement.Measurement) string {
	return strconv.FormatFloat(handleConversionToFloat(m.Measurement.Value), 'f', 1, 64)
}

func handleConversionToFloat(val interface{}) float64 {
	var value float64
	typeOfValue := reflect.ValueOf(val)
	switch typeOfValue.Kind() {
	case reflect.Int:
		value = float64(val.(int))
	case reflect.Float32:
		value = float64(val.(float32))
	case reflect.Float64:
		value = float64(val.(float64))
	case reflect.String:
		val, err := strconv.ParseFloat(val.(string), 64)
		if err != nil {
			logrus.Errorln("Error parsing number, ", val)
			return 0.0
		}
		value = val
	}

	return value
}

func layoutDoubleDigit(m measurement.Measurement) string {
	return strconv.FormatFloat(handleConversionToFloat(m.Measurement.Value), 'f', 2, 64)
}

func layoutPercentDoubleDigit(m measurement.Measurement) string {
	return strconv.FormatFloat(handleConversionToFloat(m.Measurement.Value)/100, 'f', 2, 64)
}

func GetDevicesForType(t string) []MedicalDevice {
	// exportedType, ok := exportedTypes[t]
	// if !ok {
	// 	return []MedcomDevice{}
	// }

	// return exportedType[0].Devices
	return []MedicalDevice{}
}
