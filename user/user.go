package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User -
type User struct {
	Username       string
	HashedPassword string
	Permissions    []string
	// Role           string
}

// NewUser -
func NewUser(username string, password string, permissions []string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &User{
		Username:       username,
		HashedPassword: string(hashedPassword),
		Permissions:    permissions,
		// Role:           role,
	}

	return user, nil
}

// IsCorrectPassword -
func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

// Clone -
func (user *User) Clone() *User {
	return &User{
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Permissions:    user.Permissions,
		// Role:           user.Role,
	}
}
