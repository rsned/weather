package datastructures

import (
	"math"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestFields(t *testing.T) {
	// TODO(rsned): Test for test fields.
}

var (
	// used to get a premade Less function for the SortSlices.
	ss sort.StringSlice
)

func TestPrefixLabels(t *testing.T) {

	tests := []struct {
		prefix string
		have   []string
		want   []string
	}{
		// Test prefix empty
		{
			prefix: "",
			have:   nil,
			want:   nil,
		},
		// Test prefix empty
		{
			prefix: "",
			have:   []string{"aa", "bb", "cc"},
			want:   []string{"aa", "bb", "cc"},
		},
		// Test prefix foolish chars
		{
			prefix: ".",
			have:   []string{"aa", "bb", "cc"},
			want:   []string{"aa", "bb", "cc"},
		},
		// Test normal behavior
		{
			prefix: "nested",
			have:   []string{"aa", "bb", "cc"},
			want:   []string{"nested.aa", "nested.bb", "nested.cc"},
		},
	}

	for _, test := range tests {
		got := prefixLabels(test.prefix, test.have)
		if diff := cmp.Diff(test.want, got, cmpopts.EquateEmpty(), cmpopts.SortSlices(ss.Less)); diff != "" {
			t.Errorf("prefixLables(%q, %+v) = %+v, want %+v\ndiff: %s",
				test.prefix, test.have, got, test.want, diff)
		}
	}
}

func TestFloatOrUnsetString(t *testing.T) {
	tests := []struct {
		have float64
		want string
	}{
		{
			have: 0,
			want: "0.00",
		},
		{
			have: -0,
			want: "0.00",
		},
		// Math extremes
		{
			have: math.NaN(),
			want: "NaN",
		},
		{
			have: math.Inf(-1),
			want: "-Inf",
		},
		{
			have: math.Inf(1),
			want: "+Inf",
		},
		{
			have: 987.6543,
			want: "987.65",
		},
		// Rounding when displaying less digits than value.
		{
			have: 456.7890,
			want: "456.79",
		},
		// Main case.
		{
			have: UnsetValue,
			want: UnsetValueString,
		},
		{
			have: -9999.00,
			want: UnsetValueString,
		},
		{
			have: -9999,
			want: UnsetValueString,
		},
		{
			have: -9999.0000000,
			want: UnsetValueString,
		},
	}

	for _, test := range tests {
		if got := floatOrUnsetString(test.have); got != test.want {
			t.Errorf("floatOrUnsetString(%v) = %q, want %q", test.have, got, test.want)
		}
	}
}
