package core_services

import "time"

type SecurityService interface {
	GenerateToken(userID string, duration time.Duration) (string, error)
	ValidateToken(token string) (string, error)
	HashPassword(password string) (string, error)
	ValidatePassword(hashedPassword, password string) (bool, error)
}
