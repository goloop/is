package is

import (
	"regexp"
	"strconv"
)

// decimalNumberRegex matches a plain decimal number with an optional sign:
// integer, fractional, or both. It deliberately rejects every Go-specific
// float spelling that strconv.ParseFloat would otherwise accept — scientific
// notation ("1e1"), hexadecimal floats ("0x1p4"), and digit separators
// ("1_0") — so a coordinate parsed from user input cannot smuggle in an
// exotic literal.
var decimalNumberRegex = regexp.MustCompile(`^[+-]?(?:\d+(?:\.\d*)?|\.\d+)$`)

// parseCoordinate converts a coordinate value to float64. For float64 input
// it is returned as-is (no string round-trip, no allocation). For string
// input it must be a plain decimal number — see decimalNumberRegex.
func parseCoordinate[T string | float64](v T) (float64, bool) {
	switch x := any(v).(type) {
	case float64:
		return x, true
	case string:
		if !decimalNumberRegex.MatchString(x) {
			return 0, false
		}
		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return 0, false
		}
		return f, true
	}
	return 0, false
}

// Latitude validates if the given value is a valid geographic latitude.
// A latitude represents a geographic coordinate that specifies the north-south
// position of a point on the Earth's surface. It is an angle which ranges from
// -90 to +90 degrees.
//
// The value may be a float64 or a string. A string must be a plain decimal
// number (an optional sign followed by digits and an optional fractional
// part); scientific notation, hexadecimal floats, and digit separators are
// rejected.
//
// Example usage:
//
//	is.Latitude(-45.0)    // Returns: true
//	is.Latitude("45.5")   // Returns: true
//	is.Latitude(90.1)     // Returns: false (out of range)
//	is.Latitude("1e1")    // Returns: false (scientific notation)
//	is.Latitude("0x1p4")  // Returns: false (hexadecimal float)
func Latitude[T string | float64](lat T) bool {
	v, ok := parseCoordinate(lat)
	return ok && v >= -90 && v <= 90
}

// Longitude validates if the given value is a valid geographic longitude.
// A longitude represents a geographic coordinate that specifies the east-west
// position of a point on the Earth's surface. It is an angle which ranges
// from -180 to +180 degrees.
//
// The value may be a float64 or a string. A string must be a plain decimal
// number (an optional sign followed by digits and an optional fractional
// part); scientific notation, hexadecimal floats, and digit separators are
// rejected.
//
// Example usage:
//
//	is.Longitude(-45.0)   // Returns: true
//	is.Longitude("123.4") // Returns: true
//	is.Longitude(180.1)   // Returns: false (out of range)
//	is.Longitude("1_0")   // Returns: false (digit separator)
func Longitude[T string | float64](lon T) bool {
	v, ok := parseCoordinate(lon)
	return ok && v >= -180 && v <= 180
}

// Coordinates validates if the given pair of values represent valid
// geographic coordinates. The pair consists of a latitude and a longitude.
//
// The function takes two values (float64 or string) representing the latitude
// and the longitude and checks if they fall within the valid ranges for their
// respective geographic coordinate systems.
//
// Example usage:
//
//	is.Coordinates(45.0, -123.1) // Returns: true
//	is.Coordinates(90.1, 180.1)  // Returns: false
func Coordinates[T string | float64](lat, lon T) bool {
	return Latitude(lat) && Longitude(lon)
}
