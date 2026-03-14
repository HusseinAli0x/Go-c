package entities

import (
	"math"
	"time"

	"Go_c/internal/domain/enums"

	"github.com/google/uuid"
)

// Booking represents a ride booking made by a user.
// Inherits ID, CreatedAt, UpdatedAt from Base.
type Booking struct {
	Base                                // ID, CreatedAt, UpdatedAt
	UserID          uuid.UUID           `json:"user_id" db:"user_id"`                     // The user who made the booking
	DriverID        uuid.UUID           `json:"driver_id" db:"driver_id"`                 // The assigned driver
	CarID           uuid.UUID           `json:"car_id" db:"car_id"`                       // The assigned car
	PickupLocation  Point               `json:"pickup_location" db:"pickup_location"`     // Pickup coordinates
	DropoffLocation Point               `json:"dropoff_location" db:"dropoff_location"`   // Dropoff coordinates
	DistanceKM      float64             `json:"distance_km" db:"distance_km"`             // Estimated distance
	Price           float64             `json:"price" db:"price"`                         // Booking price
	Status          enums.BookingStatus `json:"status" db:"status"`                       // Current booking status
	PaymentStatus   string              `json:"payment_status" db:"payment_status"`       // Payment status (pending, paid, failed)
	RequestedAt     time.Time           `json:"requested_at" db:"requested_at"`           // When booking was requested
	AcceptedAt      *time.Time          `json:"accepted_at,omitempty" db:"accepted_at"`   // When driver accepted
	StartedAt       *time.Time          `json:"started_at,omitempty" db:"started_at"`     // When ride started
	CompletedAt     *time.Time          `json:"completed_at,omitempty" db:"completed_at"` // When ride completed
	Version         int                 `json:"version" db:"version"`                     // Optimistic locking version
}

//
// ==========================
// Booking Behaviors
// ==========================
//

// Accept sets the booking status to "accepted" and records the accepted timestamp.
func (b *Booking) Accept() {
	now := time.Now()
	b.Status = enums.BookingAccepted
	b.AcceptedAt = &now
	b.UpdateTimestamp()
}

// Arriving sets the booking status to "arriving".
func (b *Booking) Arriving() {
	b.Status = enums.BookingArriving
	b.UpdateTimestamp()
}

// Start sets the booking status to "started" and records the start time.
func (b *Booking) Start() {
	now := time.Now()
	b.Status = enums.BookingStarted
	b.StartedAt = &now
	b.UpdateTimestamp()
}

// Complete sets the booking status to "completed" and records the completion time.
func (b *Booking) Complete() {
	now := time.Now()
	b.Status = enums.BookingCompleted
	b.CompletedAt = &now
	b.UpdateTimestamp()
}

// Cancel sets the booking status to "cancelled".
func (b *Booking) Cancel() {
	b.Status = enums.BookingCancelled
	b.UpdateTimestamp()
}

// IsActive checks whether the booking is still ongoing (not completed or cancelled).
func (b *Booking) IsActive() bool {
	return b.Status != enums.BookingCompleted && b.Status != enums.BookingCancelled
}

// UpdateDistanceAndPrice recalculates distance and price.
func (b *Booking) UpdateDistanceAndPrice(distanceKM, price float64) {
	b.DistanceKM = distanceKM
	b.Price = price
	b.UpdateTimestamp()
}

// DistanceToPickup calculates the distance in meters from a point to the pickup location.
func (b *Booking) DistanceToPickup(point Point) float64 {
	return haversineDistance(b.PickupLocation, point)
}

// DistanceToDropoff calculates the distance in meters from a point to the dropoff location.
func (b *Booking) DistanceToDropoff(point Point) float64 {
	return haversineDistance(b.DropoffLocation, point)
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
