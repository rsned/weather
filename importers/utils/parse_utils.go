package utils

import (
	"strconv"
	"strings"
)

// ParseInt attempts to get the numerical value from the string, returning the
// given default if it is unable to do so.
func ParseInt(s string, def int64) int64 {
	if i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64); err == nil {
		return i
	}
	return def
}

// ParseIntBounded attempts to parse an integer value from the given string. If the
// value is an integer, and lies within the range [min, max], it is returned.
// Otherwise the default is returned.
func ParseIntBounded(s string, min, max, def int64) int64 {
	i := ParseInt(s, def)
	if i == def {
		return def
	}

	if i < min || i > max {
		return def
	}

	return i
}

// ParseIntScaled attempts to parse an integer value from the given string, and
// if successful, scale it by the given amount. Otherwise the default is returned.
func ParseIntScaled(s string, scale float64, def int64) int64 {
	i := ParseInt(s, def)
	if i == def {
		return def
	}

	return int64(float64(i) * scale)
}

// ParseFloat attempts to get the numerical value from the string, returning the
// given default if it is unable to do so.
//
// String representations of NaN, (+/-)Inf are also supported.
func ParseFloat(s string, def float64) float64 {
	if f, err := strconv.ParseFloat(strings.TrimSpace(s), 64); err == nil {
		return f
	}
	return def
}

// ParseFloatBounded attempts to parse an float value from the given string.
// If the value is a float, and lies within the range [min, max], it is returned.
// Otherwise the default is returned.
func ParseFloatBounded(s string, min, max, def float64) float64 {
	f := ParseFloat(s, def)
	if f == def {
		return def
	}

	if f < min || f > max {
		return def
	}
	return f
}

// ParseFloatScaled attempts to parse an float value from the given string, and
// if successful, scale it by the given amount. Otherwise the default is returned.
func ParseFloatScaled(s string, scale, def float64) float64 {
	f := ParseFloat(s, def)

	if f == def {
		return def
	}

	return f * scale
}
