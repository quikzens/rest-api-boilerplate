package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// hashPassword returns the bcrypt hash of the password
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}

// checkPassword checks if the provided password is correct or not
func checkPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
