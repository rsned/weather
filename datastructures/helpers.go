package datastructures

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	UnsetValue       = -9999
	UnsetValueString = "-9999"
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
	if strings.TrimSpace(prefix) == "." {
		prefix = ""
	}
	if strings.TrimSpace(prefix) != "" {
		prefix = prefix + "."
	}

	fields := make([]string, len(labels))
	for i, v := range labels {
		fields[i] = fmt.Sprintf("%s%s", prefix, v)
	}
	return fields
}

func floatOrUnsetString(v float64) string {
	if v == UnsetValue {
		return UnsetValueString
	}
	return fmt.Sprintf("%0.2f", v)
}
