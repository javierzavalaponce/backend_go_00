package users_security_service_mock

import users_services "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/services"

type SecurityServiceMock struct{}

var _ users_services.SecurityService = (*SecurityServiceMock)(nil)

func NewSecurityServiceMock() *SecurityServiceMock {
	return &SecurityServiceMock{}
}

func (s *SecurityServiceMock) GenerateToken(userId string) (string, error) {
	return "asdfsadfsdfsdfdsafsdaf", nil
}

func (s *SecurityServiceMock) GeneratePasswordHash(password string) (string, error) {
	return "asdfsadfsdfsdfdsafsdaf", nil
}
