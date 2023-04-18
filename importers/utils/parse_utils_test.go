package utils

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// TODO(rsned): Move these somewhere more general.
const (
	UnsetValue = -9999
	eps        = 1e-15
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		have string
		def  int64
		want int64
	}{
		// Empty string inputs.
		{
			have: "",
		},
		{
			have: "",
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Bad input strings.
		{
			have: "soccer",
			def:  UnsetValue,
			want: UnsetValue,
		},
		{
			have: "1e6",
			def:  UnsetValue,
			want: UnsetValue,
		},
		{
			have: "12+3i",
			def:  UnsetValue,
			want: UnsetValue,
		},
		{
			have: "  - 37.4",
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Normal cases.
		{
			have: "111",
			def:  UnsetValue,
			want: 111,
		},
		{
			have: "0",
			def:  UnsetValue,
			want: 0,
		},
		{
			have: "-23456789",
			def:  UnsetValue,
			want: -23456789,
		},
		{
			have: "986",
			def:  UnsetValue,
			want: 986,
		},
		// Test improbable values.
		{
			have: "9223372036854775807",
			def:  UnsetValue,
			want: math.MaxInt,
		},
		{
			have: "NaN",
			def:  UnsetValue,
			want: UnsetValue,
		},
		{
			have: "Inf",
			def:  UnsetValue,
			want: UnsetValue,
		},
		// TODO(rsned): More test cases as needed.
	}
	for _, test := range tests {
		got := ParseInt(test.have, test.def)
		if got != test.want {
			t.Errorf("ParseInt(%q, %d) = %d, want %d", test.have, test.def, got, test.want)
		}
	}
}

func TestParseIntBounded(t *testing.T) {
	tests := []struct {
		have string
		min  int64
		max  int64
		def  int64
		want int64
	}{
		// Empty string inputs.
		{
			have: "",
		},
		{
			have: "",
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Bad input strings.
		{
			have: "soccer",
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Normal cases inside bounds.
		{
			have: "-5678",
			min:  -10000,
			max:  10000,
			def:  UnsetValue,
			want: -5678,
		},
		// Normal cases outside bounds.
		{
			have: "986",
			min:  0,
			max:  100,
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Poor bounds choices. (max < min)
		{
			have: "75   ",
			min:  100,
			max:  0,
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Test improbable values.
		{
			have: "-92233720368547758080", // math.MinInt*10
			def:  UnsetValue,
			want: UnsetValue,
		},
		// TODO(rsned): More test cases as needed.
	}
	for _, test := range tests {
		got := ParseIntBounded(test.have, test.min, test.max, test.def)
		if got != test.want {
			t.Errorf("ParseIntBounded(%q, %d, %d, %d) = %d, want %d",
				test.have, test.min, test.max, test.def, got, test.want)
		}
	}
}

func TestParseIntScaled(t *testing.T) {
	tests := []struct {
		have  string
		scale float64
		def   int64
		want  int64
	}{
		// Empty string inputs.
		{
			have: "",
		},
		{
			have: "",
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Bad input strings.
		{
			have: "soccer",
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Normal cases.
		{
			have: "0",
			def:  UnsetValue,
			want: 0,
		},
		{
			have:  "4115",
			scale: 20.0,
			def:   UnsetValue,
			want:  82300,
		},
		{
			have:  "10300",
			scale: 0.1,
			def:   UnsetValue,
			want:  1030,
		},
		// Test improbable values.
		{
			have:  "10300",
			scale: math.NaN(),
			def:   UnsetValue,
			want:  math.MinInt,
		},
		{
			have:  "10300",
			scale: math.Inf(-1),
			def:   UnsetValue,
			want:  math.MinInt,
		},
		// TODO(rsned): More test cases as needed.
	}
	for _, test := range tests {
		got := ParseIntScaled(test.have, test.scale, test.def)
		if got != test.want {
			t.Errorf("ParseIntScaled(%q, %f, %d) = %d, want %d", test.have, test.scale, test.def, got, test.want)
		}
	}
}

func TestParseFloat(t *testing.T) {
	tests := []struct {
		have string
		def  float64
		want float64
	}{
		// Empty string inputs.
		{
			have: "",
		},
		{
			have: "",
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Bad input strings.
		{
			have: "soccer",
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Normal cases.
		{
			have: "12.34",
			def:  UnsetValue,
			want: 12.34,
		},
		{
			have: "1e-30",
			def:  UnsetValue,
			want: 1e-30,
		},
		{
			have: "   6.02e23",
			def:  UnsetValue,
			want: 6.02e23,
		},
		{
			have: "-17.2     ",
			def:  UnsetValue,
			want: -17.2,
		},
		{
			have: "1234567.89",
			def:  UnsetValue,
			want: 1234567.89,
		},
		{
			have: "-0.33333444445555",
			def:  UnsetValue,
			want: -0.33333444445555,
		},
		// Test improbable values.
		{
			have: "Inf",
			def:  UnsetValue,
			want: math.Inf(+1),
		},
		// TODO(rsned): More test cases as needed.
	}
	for _, test := range tests {
		got := ParseFloat(test.have, test.def)
		if diff := cmp.Diff(test.want, got, cmpopts.EquateNaNs(),
			cmpopts.EquateApprox(eps, eps)); diff != "" {
			t.Errorf("ParseFloat(%q, %f) = %f, want %f, diff: %s", test.have, test.def, got, test.want, diff)
		}
	}
}

func TestParseFloatBounded(t *testing.T) {
	tests := []struct {
		have string
		min  float64
		max  float64
		def  float64
		want float64
	}{
		// Empty string inputs.
		{
			have: "",
		},
		{
			have: "",
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Bad input strings.
		{
			have: "soccer",
			def:  UnsetValue,
			want: UnsetValue,
		},
		{
			have: "57.3+18.9*    21.0000",
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Normal cases inside bounds.
		{
			have: "37.4 ",
			min:  0,
			max:  40,
			def:  UnsetValue,
			want: 37.4,
		},
		{
			have: "5280",
			min:  math.MinInt,
			max:  math.MaxInt,
			def:  UnsetValue,
			want: 5280,
		},
		// Normal cases outside bounds.
		{
			have: "-500",
			min:  -273.15,
			max:  math.MaxInt,
			def:  UnsetValue,
			want: UnsetValue,
		},
		{
			have: "31587.23",
			min:  0.0,
			max:  1200.0,
			def:  UnsetValue,
			want: UnsetValue,
		},
		// Poor bounds choices. (max < min)
		{
			have: "0",
			min:  math.MaxInt,
			max:  math.MinInt,
			def:  UnsetValue,
			want: UnsetValue,
		},
		// TODO(rsned): More test cases as needed.
	}
	for _, test := range tests {
		got := ParseFloatBounded(test.have, test.min, test.max, test.def)
		if diff := cmp.Diff(test.want, got, cmpopts.EquateNaNs(),
			cmpopts.EquateApprox(eps, eps)); diff != "" {
			t.Errorf("ParseFloatBounded(%q, %f, %f, %f) = %f, want %f, diff: %s", test.have, test.min, test.max, test.def, got, test.want, diff)
		}
	}
}

func TestParseFloatScaled(t *testing.T) {
	tests := []struct {
		have  string
		scale float64
		def   float64
		want  float64
	}{
		// Empty string inputs.
		{
			have: "",
		},
		{
			have: "",
			def:  UnsetValue,
			want: UnsetValue,
		},
		{
			have:  "",
			scale: -100.0,
			def:   UnsetValue,
			want:  UnsetValue,
		},
		// Bad input strings.
		{
			have:  "curling",
			scale: 0.1,
			def:   UnsetValue,
			want:  UnsetValue,
		},
		{
			have:  " --121",
			scale: 1.0,
			def:   UnsetValue,
			want:  UnsetValue,
		},
		{
			have:  "sqrt(2)",
			scale: 1.0,
			def:   UnsetValue,
			want:  UnsetValue,
		},
		// Normal cases.
		{
			have: "0",
			def:  UnsetValue,
			want: 0,
		},
		{
			have:  "411.5",
			scale: 20.0,
			def:   UnsetValue,
			want:  8230,
		},
		{
			have:  "11081",
			scale: 0.1,
			def:   UnsetValue,
			want:  1108.1,
		},

		// Test improbable values.
		{
			have:  "11081",
			scale: math.NaN(),
			def:   UnsetValue,
			want:  math.NaN(),
		},
		{
			have:  "+Inf",
			scale: math.NaN(),
			def:   UnsetValue,
			want:  math.NaN(),
		},
		{
			have:  "11081",
			scale: math.Inf(-1),
			def:   UnsetValue,
			want:  math.Inf(-1),
		},
		// TODO(rsned): More test cases as needed.
	}
	for _, test := range tests {
		got := ParseFloatScaled(test.have, test.scale, test.def)
		if diff := cmp.Diff(test.want, got, cmpopts.EquateNaNs(),
			cmpopts.EquateApprox(eps, eps)); diff != "" {
			t.Errorf("ParseFloatScaled(%q, %f, %f) = %f, want %f, diff: %s", test.have, test.scale, test.def, got, test.want, diff)
		}
	}
}
