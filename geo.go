package is

import (
	"fmt"
	"strconv"
)

// Latitude validates if the given value is a valid geographic latitude.
// A latitude represents a geographic coordinate that specifies the north-south
// position of a point on the Earth's surface. It is an angle which ranges from
// -90 to +90 degrees.
//
// The function takes a float64 number representing the latitude and checks if
// it falls within the valid range.
//
// Example usage:
//
//	is.Latitude(-45.0) // Returns: true
//	is.Latitude(90.1)  // Returns: false
func Latitude[T string | float64](lat T) bool {
	v, err := strconv.ParseFloat(fmt.Sprint(lat), 64)
	if err != nil {
		return false
	}

	return v >= -90 && v <= 90
}

// Longitude validates if the given value is a valid geographic longitude.
// A longitude represents a geographic coordinate that specifies the east-west
// position of a point on the Earth's surface. It is an angle which ranges
// from -180 to +180 degrees.
//
// The function takes a float64 number representing the longitude and checks
// if it falls within the valid range.
//
// Example usage:
//
//	is.Longitude(-45.0) // Returns: true
//	is.Longitude(180.1) // Returns: false
func Longitude[T string | float64](lon T) bool {
	v, err := strconv.ParseFloat(fmt.Sprint(lon), 64)
	if err != nil {
		return false
	}

	return v >= -180 && v <= 180
}

// Coordinates validates if the given pair of values represent valid
// geographic coordinates. The pair consists of a latitude and a longitude.
//
// The function takes two float64 numbers representing the latitude and the
// longitude and checks if they fall within the valid ranges for their
// respective geographic coordinate systems.
//
// Example usage:
//
//	is.Coordinates(45.0, -123.1) // Returns: true
//	is.Coordinates(90.1, 180.1)  // Returns: false
func Coordinates[T string | float64](lat, lon T) bool {
	return Latitude(lat) && Longitude(lon)
}
