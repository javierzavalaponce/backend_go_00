package users_security_service_mock

import (
	"time"

	core_services "github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/domain/services"
)

type SecurityServiceMock struct{}

var _ core_services.SecurityService = (*SecurityServiceMock)(nil)

func NewSecurityServiceMock() *SecurityServiceMock {
	return &SecurityServiceMock{}
}

func (s *SecurityServiceMock) GenerateToken(userID string, duration time.Duration) (string, error) {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", nil
}
func (s *SecurityServiceMock) ValidateToken(token string) (string, error) {
	return "userID12345", nil
}
func (s *SecurityServiceMock) HashPassword(password string) (string, error) {
	return "asdfsadfsdfsdfdsafsdaf", nil
}
func (s *SecurityServiceMock) ValidatePassword(hashedPassword, password string) (bool, error) {
	return true, nil
}
