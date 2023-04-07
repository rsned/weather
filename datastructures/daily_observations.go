package datastructures

import "strings"

var (
	dailyObsFields []string
)

func init() {
	dailyObsFields = fields(&DailyObservation{})
}

// DailyObservation holds daily summary weather observation information.
// This generally covers the min, mean, and max of the values tracked in the
// Observation type.
type DailyObservation struct {
	StationID string
	Date      string
	TempCMin  float64
	TempCMean float64
	TempCMax  float64
}

// EmptyDailyObservation returns a pre-set empty value with the missing sentinel
// values set on all relevant fields.
func EmptyDailyObservation() *DailyObservation {
	return &DailyObservation{}
}

func (a *DailyObservation) String() string {
	return a.CSV(",")
}

// CSV returns this elements values as a CSV string.
func (a *DailyObservation) CSV(delim string) string {
	return strings.Join(a.ValueColumns(), delim)
}

// HeaderColumns returns the labels for the columns in this entity.
func (a *DailyObservation) HeaderColumns(prefix string) []string {
	// TODO(rsned): Cache the list by prefix to save on redundant work.
	return prefixLabels(prefix, dailyObsFields)
}

// ValueColumns returns the values for this entity as a collection of strings
// in the same order as the HeaderColumns.
func (a *DailyObservation) ValueColumns() []string {
	return []string{}
}
