package exporttypes

// currently, this is a just a hard code list. The main drawback with that is, that it requires a release to handle that mapping.
// Perhaps it would be prudent to make it more dynamically. One way was to do to create a dirrectory with a file per measuerement type
//
// Eash file would contain a list of debice. for each device we could specificy the following:
// - medcom id
// - manufacturer
// - model
// - hardwareversion
// - softwareversion
// - fimrwareversion

// Then new devices can be added on the fly. The files can be parsed some what easier and get build into a decent structure to search in

import (
	"strings"

	"bitbucket.org/opentelehealth/exporter/measurement"
)

func NewAnDMedical302PBTExportType() simpleDevice {
	d := simpleDevice{}

	d.medComId = "MCI00010"
	d.manufacturer = "A&D Medical"
	d.productType = PRODUCT_TYPE_THERMOMETER
	d.model = "UT-302PlusBT Bluetooth"

	d.compareFunction = func(origin measurement.Origin) bool {
		if strings.EqualFold(d.manufacturer, origin.DeviceMeasurement.Manufacturer) {
			if strings.Contains(strings.ToLower(origin.DeviceMeasurement.Model), "302") {
				return true
			}
		}
		return false
	}
	return d
}

func NewAnDMedical321PBTCExportType() simpleDevice {
	d := simpleDevice{}

	d.medComId = "MCI00002"
	d.manufacturer = "A&D Medical"
	d.productType = PRODUCT_TYPE_WEIGHT
	d.model = "UC-321PlusBT-C Bluetooth"

	d.compareFunction = func(origin measurement.Origin) bool {
		if strings.EqualFold(d.manufacturer, origin.DeviceMeasurement.Manufacturer) {
			if strings.Contains(strings.ToLower(origin.DeviceMeasurement.Model), "321") {
				return true
			}
		}
		return false
	}
	return d
}

func NewAnDMedical351PBTCiExportType() simpleDevice {
	d := simpleDevice{}
	d.medComId = "MCI00011"
	d.productType = PRODUCT_TYPE_WEIGHT
	d.manufacturer = "A&D Medical"
	d.model = "UC-351PlusBT-Ci Bluetooth"

	d.compareFunction = func(origin measurement.Origin) bool {
		if strings.EqualFold(d.manufacturer, origin.DeviceMeasurement.Manufacturer) {
			if strings.Contains(strings.ToLower(origin.DeviceMeasurement.Model), "351") {
				return true
			}
		}
		return false
	}

	return d
}

func NewAnDMedical767PBTCExportType() simpleDevice {

	d := simpleDevice{}

	d.medComId = "MCI00004"
	d.productType = PRODUCT_TYPE_BLOODPRESSUREMONITOR
	d.manufacturer = "A&D Medical"
	d.model = "UA-767PlusBT-C Bluetooth"

	d.compareFunction = func(origin measurement.Origin) bool {
		model := strings.ToLower(origin.DeviceMeasurement.Model)
		if strings.EqualFold(d.manufacturer, origin.DeviceMeasurement.Manufacturer) {
			if strings.Contains(model, "767") {
				if strings.Contains(model, "c") {
					return true
				}
			}
		}
		return false
	}
	return d
}

func NewAnDMedical767PBTCiExportType() simpleDevice {
	d := simpleDevice{}

	d.medComId = "MCI00012"
	d.productType = PRODUCT_TYPE_BLOODPRESSUREMONITOR
	d.manufacturer = "A&D Medical"
	d.model = "UA-767PlusBT-Ci Bluetooth"

	d.compareFunction = func(origin measurement.Origin) bool {
		model := strings.ToLower(origin.DeviceMeasurement.Model)
		if strings.EqualFold(d.manufacturer, origin.DeviceMeasurement.Manufacturer) {
			if strings.Contains(model, "767") {
				if strings.Contains(model, "ci") {
					return true
				}
			}
		}
		return false
	}
	return d
}

func NewNonin3230ExportType() simpleDevice {

	d := simpleDevice{}

	d.medComId = "MCI00013"
	d.productType = PRODUCT_TYPE_PULSEOXIMETER
	d.manufacturer = "Nonin"
	d.model = "3230 Bluetooth Smart Pulse Oximeter"
	d.compareFunction = func(origin measurement.Origin) bool {
		if strings.EqualFold(d.manufacturer, origin.DeviceMeasurement.Manufacturer) {
			if strings.Contains(strings.ToLower(origin.DeviceMeasurement.Model), "3230") {
				return true
			}
		}
		return false
	}
	return d
}

func NewNonin9560ExportType() simpleDevice {
	d := simpleDevice{}

	d.medComId = "MCI00005"
	d.productType = PRODUCT_TYPE_PULSEOXIMETER
	d.manufacturer = "Nonin"
	d.model = "Onyx II 9560 Bluetooth Pulse Oximeter"
	d.compareFunction = func(origin measurement.Origin) bool {
		if strings.EqualFold(d.manufacturer, origin.DeviceMeasurement.Manufacturer) {
			if strings.Contains(strings.ToLower(origin.DeviceMeasurement.Model), "9560") {
				return true
			}
		}
		return false
	}

	return d
}

func NewVitalograph4000ExportType() simpleDevice {
	d := simpleDevice{}

	d.medComId = "MCI00014"
	d.productType = PRODUCT_TYPE_LUNGMONITOR
	d.manufacturer = "Vitalograph"
	d.model = "4000 Lung Monitor Bluetooth"

	d.compareFunction = func(origin measurement.Origin) bool {
		return strings.EqualFold(d.manufacturer, origin.DeviceMeasurement.Manufacturer)
	}

	return d
}
