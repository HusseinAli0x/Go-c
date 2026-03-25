package entities

import (
	"Go_c/internal/domain/enums"
	domainErrors "Go_c/internal/domain/errors"
	"time"

	"github.com/google/uuid"
)

// Payment represents a payment for a booking
type Payment struct {
	Base

	BookingID     uuid.UUID           `json:"booking_id" db:"booking_id"`         // Related booking
	Amount        float64             `json:"amount" db:"amount"`                 // Payment amount
	PaymentMethod string              `json:"payment_method" db:"payment_method"` // e.g. card, cash
	Status        enums.PaymentStatus `json:"status" db:"status"`                 // Payment status
	DeletedAt     *time.Time          `json:"deleted_at,omitempty" db:"deleted_at"`
}

// ==========================
// Behaviors
// ==========================
//

// MarkPaid marks payment as paid
func (p *Payment) MarkPaid() error {
	if p.Status == enums.PaymentPaid {
		return domainErrors.ErrPaymentAlreadyExists
	}
	p.Status = enums.PaymentPaid
	p.UpdateTimestamp()
	return nil
}

// MarkFailed marks payment as failed
func (p *Payment) MarkFailed() {
	p.Status = enums.PaymentFailed
	p.UpdateTimestamp()
}

// IsPending checks if payment is still pending
func (p *Payment) IsPending() bool {
	return p.Status == enums.PaymentPending
}

// SoftDelete marks payment as deleted
func (p *Payment) SoftDelete() {
	now := time.Now()
	p.DeletedAt = &now
	p.UpdateTimestamp()
}

// Validate checks required fields
func (p *Payment) Validate() error {
	if p.BookingID == uuid.Nil {
		return domainErrors.ErrInvalidInput
	}
	if p.Amount <= 0 {
		return domainErrors.ErrInvalidInput
	}
	if p.PaymentMethod == "" {
		return domainErrors.ErrInvalidInput
	}
	return nil
}
