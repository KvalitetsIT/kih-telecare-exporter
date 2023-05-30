package exporttypes

import (
	"testing"

	"github.com/KvalitetsIT/kih-telecare-exporter/measurement"
)

func TestFloatConversion(t *testing.T) {
	tests := []struct {
		name   string
		input  interface{}
		output float64
	}{
		{"String", "26.2", 26.2},
		{"Int", 26, 26.0},
		{"Float", 26.2, 26.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := handleConversionToFloat(tt.input)
			if tt.output != res {
				t.Errorf("Expected '%f' got '%f'", tt.output, res)
			}
		})
	}
}

func TestSingleDigitLayout(t *testing.T) {
	tests := []struct {
		name   string
		input  interface{}
		output string
	}{
		{"No digits String", "26", "26.0"},
		{"one digits String", "26.0", "26.0"},
		{"Two digits String", "26.00", "26.0"},
		{"Int", 26, "26.0"},
		{"Float one digits ", 26.0, "26.0"},
		{"Float two digits ", 26.02, "26.0"},
		{"Float two digits rounding ", 26.06, "26.1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := measurement.Measurement{}
			m.Measurement.Value = tt.input

			res := layoutSingleDigit(m)
			if tt.output != res {
				t.Errorf("Expected '%v' got '%v'", tt.output, res)
			}
		})
	}

}

func TestLayoutDoubleDigiti(t *testing.T) {
	tests := []struct {
		name   string
		input  interface{}
		output string
	}{
		{"No digits String", "26", "26.00"},
		{"one digits String", "26.0", "26.00"},
		{"Two digits String", "26.010", "26.01"},
		{"Two digits String", "26.015", "26.02"},
		{"Int", 26, "26.00"},
		{"Float one digits ", 26.0, "26.00"},
		{"Float two digits ", 26.02, "26.02"},
		{"Float two digits rounding ", 26.060, "26.06"},
		{"Float two digits rounding ", 26.006, "26.01"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := measurement.Measurement{}
			m.Measurement.Value = tt.input

			res := layoutDoubleDigit(m)
			if tt.output != res {
				t.Errorf("Expected '%v' got '%v'", tt.output, res)
			}
		})
	}
}

func TestLayoutPercentDoubleDigits(t *testing.T) {
	tests := []struct {
		name   string
		input  interface{}
		output string
	}{
		{"No digits String", "26", "0.26"},
		{"one digits String", "26.0", "0.26"},
		{"Two digits String", "26.60", "0.27"},
		{"Two digits String", "26.015", "0.26"},
		{"Int", 26, "0.26"},
		{"Float one digits ", 26.0, "0.26"},
		{"Float two digits ", 26.02, "0.26"},
		{"Float two digits rounding ", 26.60, "0.27"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := measurement.Measurement{}
			m.Measurement.Value = tt.input

			res := layoutPercentDoubleDigit(m)
			if tt.output != res {
				t.Errorf("Expected '%v' got '%v'", tt.output, res)
			}
		})
	}
}
