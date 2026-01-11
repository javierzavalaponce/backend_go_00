package users_use_cases

import (
	core_services "github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/domain/services"
	users_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/models"
	users_respositories "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/respositories"
)

type SignUpUseCase struct {
	usersRepository users_respositories.UsersRepository
	securityService core_services.SecurityService
}

func NewSignUpUseCase(
	usersRepository users_respositories.UsersRepository,
	securityService core_services.SecurityService,
) *SignUpUseCase {
	return &SignUpUseCase{
		usersRepository: usersRepository,
		securityService: securityService,
	}
}

func (s *SignUpUseCase) Execute(signData *users_models.SignData) (string, error) {
	hashedPassword, err := s.securityService.HashPassword(signData.Password)
	if err != nil {
		return "", err
	}

	userId, err := s.usersRepository.Create(signData.Email, hashedPassword)
	if err != nil {
		return "", err
	}

	token, err := s.securityService.GenerateToken(userId.String(), 24*60*60)
	if err != nil {
		return "", err
	}

	return token, nil
}
