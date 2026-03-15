package entities

import (
	"github.com/google/uuid"

	"Go_c/internal/domain/enums"
)

// Payment represents a payment made for a booking.
type Payment struct {
	Base                              // ID, CreatedAt, UpdatedAt
	BookingID     uuid.UUID           `json:"booking_id" db:"booking_id"`         // Related booking
	Amount        float64             `json:"amount" db:"amount"`                 // Payment amount
	PaymentMethod string              `json:"payment_method" db:"payment_method"` // Method (cash, card, wallet)
	Status        enums.PaymentStatus `json:"status" db:"status"`                 // Payment status
}

//
// ==========================
// Payment Behaviors
// ==========================
//

// MarkPaid marks the payment as completed.
func (p *Payment) MarkPaid() {
	p.Status = enums.PaymentPaid
	p.UpdateTimestamp()
}

// MarkFailed marks the payment as failed.
func (p *Payment) MarkFailed() {
	p.Status = enums.PaymentFailed
	p.UpdateTimestamp()
}

// IsPaid checks if the payment was successful.
func (p *Payment) IsPaid() bool {
	return p.Status == enums.PaymentPaid
}
