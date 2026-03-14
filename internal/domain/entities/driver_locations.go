package entities

import (
	"math"
	"time"

	"github.com/google/uuid"
)

// DriverLocation represents the current location of a driver.
// Inherits ID, CreatedAt, UpdatedAt from Base.
type DriverLocation struct {
	Base               // ID, CreatedAt, UpdatedAt
	DriverID uuid.UUID `json:"driver_id" db:"driver_id"` // Reference to the driver
	Location Point     `json:"location" db:"location"`   // Current geographic location
}

//
// ==========================
// DriverLocation Behaviors
// ==========================
//

// UpdateLocation updates the driver's current location and timestamp.
func (d *DriverLocation) UpdateLocation(lat, lng float64) {
	d.Location.Latitude = lat
	d.Location.Longitude = lng
	d.UpdateTimestamp()
}

// IsRecent checks if the location was updated within the given duration (in seconds).
func (d *DriverLocation) IsRecent(durationInSeconds int64) bool {
	elapsed := time.Since(d.UpdatedAt).Seconds()
	return int64(elapsed) <= durationInSeconds
}

// DistanceToPoint calculates approximate distance in meters to another point.
// Uses Haversine formula (good enough for GPS distance in small projects).
func DistanceToPoint(a, b Point) float64 {
	const R = 6371000
	dLat := (b.Latitude - a.Latitude) * math.Pi / 180
	dLng := (b.Longitude - a.Longitude) * math.Pi / 180
	lat1 := a.Latitude * math.Pi / 180
	lat2 := b.Latitude * math.Pi / 180

	sinDLat := math.Sin(dLat / 2)
	sinDLng := math.Sin(dLng / 2)

	aVal := sinDLat*sinDLat + math.Cos(lat1)*math.Cos(lat2)*sinDLng*sinDLng
	c := 2 * math.Atan2(math.Sqrt(aVal), math.Sqrt(1-aVal))

	return R * c
}

// helper functions
func sinSquared(x float64) float64   { return (1 - cos(2*x)) / 2 }
func cos(x float64) float64          { return math.Cos(x) }
func atan2Sqrt(y, x float64) float64 { return math.Atan2(math.Sqrt(y), math.Sqrt(x)) }
