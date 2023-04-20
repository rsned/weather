package datastructures

import (
	"strings"
)

// These are used to aid in the reflection based understanding of the data structures
// so that manually curated lists of fields need to be kept up to date.
var (
	stationFields []string
)

func init() {
	stationFields = fields(&Station{})
}

// Station represents one physical location that records weather observations, and
// the set of related identitifiers and information about its capabilities.
type Station struct {
	// ID is a unique ID generated for this station. It is explicitly
	// distinct from the various other NGO and national based identifiers.
	ID string `beam:"id"`

	Name string `beam:"name"`

	// Identifiers contains the various system specific identifiers for this
	// station such as WMO, NCEID, etc.
	Identifiers *Identifiers `beam:"identifiers"`

	// Geography is a collection of geographical identifiers for the stations
	// location. This includes both administrative levels where known (such as
	// the containing Region/Country, State/Province/Prefecture/etc., Locality,
	// Postal Code), as well as some common data points (Latitude/Longitude,
	// spatial geometry cells, etc.)
	Geography *Geography `beam:"geography"`

	// Attributions contains data useable to identify which systems data
	// were incorporated to the data about this Station.
	Attributions *Attributions `beam:"attributions"`

	StartDate   string `beam:"start_date"`
	EndDate     string `beam:"end_date"`
	LastUpdated string `beam:"last_updated"`
}

func EmptyStation() *Station {
	return &Station{
		Identifiers:  &Identifiers{},
		Geography:    &Geography{},
		Attributions: &Attributions{},
	}
}

func (s *Station) String() string {
	return s.CSV(",")
}

func (s *Station) CSV(delim string) string {
	// TODO(rsned): Move this to use full CSV encoding.
	return strings.Join(s.ValueColumns(), delim)
}

func (s *Station) HeaderColumns(prefix string) []string {
	// TODO(rsned): Cache the list by prefix to save on redundant work.
	cols := prefixLabels(prefix, stationFields)[0:2]
	cols = append(cols, s.Identifiers.HeaderColumns("ids")...)
	cols = append(cols, s.Geography.HeaderColumns("geo")...)
	cols = append(cols, s.Attributions.HeaderColumns("attr")...)
	cols = append(cols, prefixLabels(prefix, stationFields)[2:]...)

	return cols
}

func (s *Station) ValueColumns() []string {
	cols := []string{
		s.ID,
		s.Name,
	}
	cols = append(cols, s.Identifiers.ValueColumns()...)
	cols = append(cols, s.Geography.ValueColumns()...)
	cols = append(cols, s.Attributions.ValueColumns()...)
	cols = append(cols, []string{
		s.StartDate,
		s.EndDate,
		s.LastUpdated,
	}...)
	return cols
}
