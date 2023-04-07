/*
Package datastructures contains the basic data types used. The main goal is to store
data in a straighforward unified manner to make it easier to use the data once it is
ingested.  (e.g., Converting all units to SI standards, time shifting to use UTC for
all datapoints so the user doesn't have to figure out forward and backward timezone
offsets themselves.)
*/
package datastructures

import (
	"fmt"
	"reflect"
	"strings"
)

type Outputter interface {
	CSV() string
	HeaderColumns() []string
	ValueColumns() []string
}

func fields(s any) []string {
	var f []string
	te := reflect.TypeOf(s).Elem()
	for _, v := range reflect.VisibleFields(te) {
		f = append(f, v.Name)
		// Name, PkgPath, Tag, Offset, Index, Anonymous
		// Tag = space separated k:v pairs
		// e.g. Tag:beam:"id" json:"id"
	}

	return f
}

func prefixLabels(prefix string, labels []string) []string {
	if strings.TrimSpace(prefix) != "" {
		prefix = prefix + "."
	}

	fields := make([]string, len(labels))
	for i, v := range labels {
		fields[i] = fmt.Sprintf("%s%s", prefix, v)
	}
	return fields
}
