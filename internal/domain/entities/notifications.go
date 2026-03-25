package entities

import (
	"Go_c/internal/domain/errors"
	"time"

	"github.com/google/uuid"
)

// Notification represents a message sent to a user
type Notification struct {
	Base

	UserID uuid.UUID `json:"user_id" db:"user_id"` // Recipient
	Title  string    `json:"title" db:"title"`     // Notification title
	Body   string    `json:"body" db:"body"`       // Notification body
	IsRead bool      `json:"is_read" db:"is_read"` // Whether notification has been read
}

// ==========================
// Behaviors
// ==========================
//

// MarkRead marks notification as read
func (n *Notification) MarkRead() {
	n.IsRead = true
	n.UpdateTimestamp()
}

// MarkUnread marks notification as unread
func (n *Notification) MarkUnread() {
	n.IsRead = false
	n.UpdateTimestamp()
}

// UpdateContent updates title and body
func (n *Notification) UpdateContent(title, body string) error {
	if title == "" || body == "" {
		return errors.ErrInvalidInput
	}

	n.Title = title
	n.Body = body
	n.UpdateTimestamp()
	return nil
}

// IsRecent checks if notification was created/updated within given duration (seconds)
func (n *Notification) IsRecent(durationInSeconds int64) bool {
	elapsed := time.Since(n.UpdatedAt).Seconds()
	return int64(elapsed) <= durationInSeconds
}
