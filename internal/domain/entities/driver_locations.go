package entities

import (
	domainErrors "Go_c/internal/domain/errors"
	"math"
	"time"

	"github.com/google/uuid"
)

// DriverLocation represents the real-time GPS location of a driver
type DriverLocation struct {
	Base               // ID, CreatedAt, UpdatedAt
	DriverID uuid.UUID `json:"driver_id" db:"driver_id"` // Reference to the driver
	Location Point     `json:"location" db:"location"`   // Current geographic location
}

//
// ==========================
// Behaviors
// ==========================
//

// UpdateLocation sets a new latitude and longitude
func (d *DriverLocation) UpdateLocation(lat, lng float64) error {
	if lat < -90 || lat > 90 || lng < -180 || lng > 180 {
		return domainErrors.ErrInvalidInput
	}

	d.Location.Latitude = lat
	d.Location.Longitude = lng
	d.UpdateTimestamp()
	return nil
}

// IsRecent checks if the location was updated within the given duration (in seconds)
func (d *DriverLocation) IsRecent(durationInSeconds int64) bool {
	elapsed := time.Since(d.UpdatedAt).Seconds()
	return int64(elapsed) <= durationInSeconds
}

// DistanceTo calculates distance in meters from this location to a point
func (d *DriverLocation) DistanceTo(p Point) float64 {
	return haversineDistance(d.Location, p)
}

// ==========================
// Helper: Haversine Formula
// ==========================
func haversineDistance(a, b Point) float64 {
	const R = 6371000 // Earth radius in meters
	lat1 := a.Latitude * math.Pi / 180
	lat2 := b.Latitude * math.Pi / 180
	deltaLat := (b.Latitude - a.Latitude) * math.Pi / 180
	deltaLng := (b.Longitude - a.Longitude) * math.Pi / 180

	sinLat := math.Sin(deltaLat / 2)
	sinLng := math.Sin(deltaLng / 2)

	c := 2 * math.Atan2(math.Sqrt(sinLat*sinLat+math.Cos(lat1)*math.Cos(lat2)*sinLng*sinLng),
		math.Sqrt(1-(sinLat*sinLat+math.Cos(lat1)*math.Cos(lat2)*sinLng*sinLng)))
	return R * c
}
