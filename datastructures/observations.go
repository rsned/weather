package datastructures

import (
	"strings"
)

var (
	obsFields []string
)

func init() {
	obsFields = fields(&Observation{})
}

// Observation holds the set of all potentially useful fields at a given point in time.
// Dates and Times are expected to be in UTC time, adjusted as needed by the importer tools.
// Times are expected to only be at minute level granularity.  No seconds are stored.
// Any expectations on rounding and window times will be in the accompanying documentation.
// e.g., If an observation is for a 10 minute period, is the time recorded as 00 or 05, or 09?
type Observation struct {
	StationID string
	Date      string // Date UTC in YYYYMMDD format.
	Time      string // Time UTC in 24 HR HHMM format.

	TempC float64
	// TODO(rsned): Add additional value types.
}

// EmptyObservation returns a pre-set empty value with the missing sentinel
// values set on all relevant fields.
func EmptyObservation() *Observation {
	return &Observation{
		TempC: UnsetValue,
	}
}

func (o *Observation) String() string {
	return o.CSV(",")
}

// CSV returns this elements values as a CSV string.
func (o *Observation) CSV(delim string) string {
	return strings.Join(o.ValueColumns(), delim)
}

// HeaderColumns returns the labels for the columns in this entity.
func (o *Observation) HeaderColumns(prefix string) []string {
	// TODO(rsned): Cache the list by prefix to save on redundant work.
	return prefixLabels(prefix, obsFields)
}

// ValueColumns returns the values for this entity as a collection of strings
// in the same order as the HeaderColumns.
func (o *Observation) ValueColumns() []string {
	return []string{
		o.StationID,
		o.Date,
		o.Time,
		floatOrUnsetString(o.TempC),
	}
}
