package users_use_cases

import (
	users_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/models"
	users_respositories "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/respositories"
	users_services "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/services"
)

type SignUpUseCase struct {
	usersRepository users_respositories.UsersRepository
	securityService users_services.SecurityService
}

func NewSignUpUseCase(
	usersRepository users_respositories.UsersRepository,
	securityService users_services.SecurityService,
) *SignUpUseCase {
	return &SignUpUseCase{
		usersRepository: usersRepository,
		securityService: securityService,
	}
}

func (s *SignUpUseCase) Execute(signData *users_models.SignData) (string, error) {
	hashedPassword, err := s.securityService.GeneratePasswordHash(signData.Password)
	if err != nil {
		return "", err
	}

	userId, err := s.usersRepository.Create(signData, hashedPassword)
	if err != nil {
		return "", err
	}

	token, err := s.securityService.GenerateToken(userId.String())
	if err != nil {
		return "", err
	}

	return token, nil
}
