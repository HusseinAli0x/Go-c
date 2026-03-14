package entities

import (
	"Go_c/internal/domain/enums"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// User entity with Base for ID and timestamps
type User struct {
	Base                        // ID, CreatedAt, UpdatedAt
	Name         string         `json:"name" db:"name"`
	Phone        string         `json:"phone" db:"phone"`
	Email        *string        `json:"email,omitempty" db:"email"`
	PasswordHash string         `json:"-" db:"password_hash"`
	Role         enums.UserRole `json:"role" db:"role"`
	ProfileImage *string        `json:"profile_image,omitempty" db:"profile_image"`
	Version      int            `json:"version" db:"version"`
}

// ==========================
// User Behaviors
// ==========================

// SetPassword hashes the plain password and sets PasswordHash
func (u *User) SetPassword(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

// CheckPassword verifies the password against hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

// PromoteToAdmin changes the role to admin
func (u *User) PromoteToAdmin() {
	u.Role = enums.RoleAdmin
}
