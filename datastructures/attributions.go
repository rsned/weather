package datastructures

import (
	"strings"
)

// These are used to aid in the reflection based understanding of the data structures
// so that manually curated lists of fields need to be kept up to date.
var (
	attributionsFields []string
)

func init() {
	attributionsFields = fields(&Attributions{})
}

// Attributions is a collection of attribution messages and tags for data used in a
// station or observation.
type Attributions struct {
}

func (a *Attributions) String() string {
	return a.CSV(",")
}

// CSV returns this elements values as a CSV string.
func (a *Attributions) CSV(delim string) string {
	return strings.Join(a.ValueColumns(), delim)
}

// HeaderColumns returns the labels for the columns in this entity.
func (a *Attributions) HeaderColumns(prefix string) []string {
	// TODO(rsned): Cache the list by prefix to save on redundant work.
	return prefixLabels(prefix, attributionsFields)
}

// ValueColumns returns the values for this entity as a collection of strings
// in the same order as the HeaderColumns.
func (a *Attributions) ValueColumns() []string {
	return []string{}
}
