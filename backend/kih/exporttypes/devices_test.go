package exporttypes

import (
	"testing"

	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
)

func TestNewAnDMedical302PBTExportType(t *testing.T) {

	differentDevice := measurement.DeviceMeasurement{Manufacturer: "Something", Model: "Some model"}

	manualOrigin := measurement.Origin{}
	manualOrigin.ManualMeasurement.EnteredBy = "clinician"

	tests := []struct {
		name     string
		medcomID string
		device   simpleDevice
		origin   measurement.Origin
		same     bool
	}{
		{name: "AndMedical302PT", medcomID: "MCI00010", device: NewAnDMedical302PBTExportType(), origin: measurement.Origin{DeviceMeasurement: measurement.DeviceMeasurement{Manufacturer: "A&D Medical", Model: "UT-302PlusBT Bluetooth"}}, same: true},
		{name: "AndMedical321PBT", medcomID: "MCI00002", device: NewAnDMedical321PBTCExportType(), origin: measurement.Origin{DeviceMeasurement: measurement.DeviceMeasurement{Manufacturer: "A&D Medical", Model: "UC-321PlusBT-C Bluetooth"}}, same: true},
		{name: "AnDMedical767PBTC", medcomID: "MCI00004", device: NewAnDMedical767PBTCExportType(), origin: measurement.Origin{DeviceMeasurement: measurement.DeviceMeasurement{Manufacturer: "A&D Medical", Model: "UA-767PlusBT-C Bluetooth"}}, same: true},
		{name: "AnDMedical767PBTCiExportType", medcomID: "MCI00012", device: NewAnDMedical767PBTCiExportType(), origin: measurement.Origin{DeviceMeasurement: measurement.DeviceMeasurement{Manufacturer: "A&D Medical", Model: "UA-767PlusBT-Ci Bluetooth"}}, same: true},
		{name: "Nonin3230ExportType", medcomID: "MCI00013", device: NewNonin3230ExportType(), origin: measurement.Origin{DeviceMeasurement: measurement.DeviceMeasurement{Manufacturer: "Nonin", Model: "3230 Bluetooth Smart Pulse Oximeter"}}, same: true},
		{name: "Nonin9560ExportType", medcomID: "MCI00005", device: NewNonin9560ExportType(), origin: measurement.Origin{DeviceMeasurement: measurement.DeviceMeasurement{Manufacturer: "Nonin", Model: "Onyx II 9560 Bluetooth Pulse Oximeter"}}, same: true},
		{name: "Vitalograph4000ExportType", medcomID: "MCI00014", device: NewVitalograph4000ExportType(), origin: measurement.Origin{DeviceMeasurement: measurement.DeviceMeasurement{Manufacturer: "Vitalograph", Model: "4000 Lung Monitor Bluetooth"}}, same: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSame := tt.device.CheckIfSameModel(tt.origin)
			if isSame != tt.same {
				t.Errorf("%+v should match as same device. got %vb ", tt.origin, isSame)
			}

			isSame = tt.device.CheckIfSameModel(measurement.Origin{DeviceMeasurement: differentDevice})
			if isSame == true {
				t.Errorf("%+v should not match as same device. got %v ", differentDevice, isSame)
			}

			isSame = tt.device.CheckIfSameModel(manualOrigin)
			if isSame == true {
				t.Errorf("%+v should not match as same device. got %v ", manualOrigin, isSame)
			}
			if tt.medcomID != tt.device.GetMedcomID() {
				t.Errorf("Expected %s got %s", tt.medcomID, tt.device.GetMedcomID())
			}
		})
	}
}
