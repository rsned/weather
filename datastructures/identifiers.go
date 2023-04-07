package datastructures

import (
	"strings"
)

// These are used to aid in the reflection based understanding of the data structures
// so that manually curated lists of fields need to be kept up to date.
var (
	identifiersFields []string
)

func init() {
	identifiersFields = fields(&Identifiers{})
}

type Identifier struct {
	IDKey string `beam:"id_key" json:"id_key"`
	Value string `beam:"value" json:"value"`
}

// Identifiers contains the set of known current identifiers for a given Station.
// Historical identifiers or a time series list of changes are not a part of this
// type.
type Identifiers struct {
	WmoID     string `beam:"wmo_id" json:"wmo_id"`
	GhcnID    string `beam:"ghcn_id" json:"ghcn_id"`
	GhcnIDAlt string `beam:"ghcn_id_alt" json:"ghcn_id_alt"`

	IATA string `beam:"iata" json:"iata"`
	ICAO string `beam:"icao" json:"icao"`

	// RegionalAviationCodes is a map of ISO 3166-1 region code to the aviation
	// code from that regions air authority.
	// e.g., US => "SFO"
	RegionalAviationCodes map[string]string `beam:"regional_aviation_codes" json:regional_aviation_codes"`

	// RegionalSpecificIDs is a map of ISO 3166-1 region codes to the collection
	// of identifiers assigned by that regions authority.
	// e.g., US => [EPA:11432 ICOADS:3312 HCDN:AL293]
	RegionalIDs map[string]string `beam:"regional_ids" json:"regional_ids"`

	// TODO(rsned): Add more identifiers.
}

func (i *Identifiers) String() string {
	return i.CSV(",")
}

// CSV returns this elements values as a CSV string.
func (i *Identifiers) CSV(delim string) string {
	// TODO(rsned): Move this to use full CSV encoding.
	return strings.Join(i.ValueColumns(), delim)
}

// HeaderColumns returns the labels for the columns in this entity.
func (i *Identifiers) HeaderColumns(prefix string) []string {
	// TODO(rsned): Cache the list by prefix to save on redundant work.
	return prefixLabels(prefix, identifiersFields)
}

// ValueColumns returns the values for this entity as a collection of strings
// in the same order as the HeaderColumns.
func (i *Identifiers) ValueColumns() []string {
	return []string{
		i.WmoID,
		i.GhcnID,
		i.GhcnIDAlt,
		i.ICAO,
		i.IATA,
		"map of regional airport codes",
		"map of regional ids",
	}
}
