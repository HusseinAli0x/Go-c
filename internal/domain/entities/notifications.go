package entities

import "github.com/google/uuid"

// Notification represents a system notification sent to a user.
type Notification struct {
	Base             // ID, CreatedAt, UpdatedAt
	UserID uuid.UUID `json:"user_id" db:"user_id"` // Target user
	Title  string    `json:"title" db:"title"`     // Notification title
	Body   string    `json:"body" db:"body"`       // Notification message
	IsRead bool      `json:"is_read" db:"is_read"` // Read status
}

//
// ==========================
// Notification Behaviors
// ==========================
//

// MarkAsRead marks the notification as read.
func (n *Notification) MarkAsRead() {
	n.IsRead = true
	n.UpdateTimestamp()
}

// MarkAsUnread marks the notification as unread.
func (n *Notification) MarkAsUnread() {
	n.IsRead = false
	n.UpdateTimestamp()
}

// IsUnread checks whether the notification has not been read yet.
func (n *Notification) IsUnread() bool {
	return !n.IsRead
}
