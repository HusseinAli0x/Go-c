package entities

import (
	"Go_c/internal/domain/enums"
	domainErrors "Go_c/internal/domain/errors"

	"golang.org/x/crypto/bcrypt"
)

// User represents a system user (customer, driver, admin)
type User struct {
	Base

	Name         string         `json:"name" db:"name"`             // Full name
	Phone        *string        `json:"phone,omitempty" db:"phone"` // Optional phone
	Email        string         `json:"email" db:"email"`           // Required email
	PasswordHash string         `json:"-" db:"password_hash"`       // Hashed password
	Role         enums.UserRole `json:"role" db:"role"`             // Role
	ProfileImage *string        `json:"profile_image,omitempty" db:"profile_image"`
}

//
// ==========================
// Behaviors
// ==========================
//

// SetPassword hashes and stores password
func (u *User) SetPassword(password string) error {
	if len(password) < 6 {
		return domainErrors.ErrInvalidInput
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domainErrors.ErrInternal
	}

	u.PasswordHash = string(hash)
	u.UpdateTimestamp()
	return nil
}

// CheckPassword verifies password
func (u *User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return domainErrors.ErrInvalidPassword
	}
	return nil
}

// UpdateProfile updates user info
func (u *User) UpdateProfile(name string, phone *string) {
	u.Name = name
	u.Phone = phone
	u.UpdateTimestamp()
}

// SetProfileImage updates profile image
func (u *User) SetProfileImage(image string) {
	u.ProfileImage = &image
	u.UpdateTimestamp()
}

// Validate checks required fields
func (u *User) Validate() error {
	if u.Name == "" {
		return domainErrors.ErrInvalidInput
	}
	if u.Email == "" {
		return domainErrors.ErrInvalidInput
	}
	return nil
}

// Role helpers

func (u *User) IsDriver() bool {
	return u.Role == enums.RoleDriver
}

func (u *User) IsAdmin() bool {
	return u.Role == enums.RoleAdmin
}

func (u *User) IsCustomer() bool {
	return u.Role == enums.RoleCustomer
}
