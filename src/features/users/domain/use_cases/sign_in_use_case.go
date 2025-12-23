package users_use_cases

import (
	users_respositories "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/respositories"
	users_services "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/services"
)

type SignInUseCase struct {
	usersRepository users_respositories.UsersRepository
	securityService users_services.SecurityService
}

func NewSignInUseCase(
	usersRepository users_respositories.UsersRepository,
	securityService users_services.SecurityService,
) *SignInUseCase {

	return &SignInUseCase{
		usersRepository: usersRepository,
		securityService: securityService,
	}
}

func (s *SignInUseCase) Exec() (string, error) {
	// Aquí iría la lógica para verificar las credenciales del usuario
	token, err := s.securityService.GenerateToken("user-id-placeholder")
	if err != nil {
		return "", err
	}

	return token, nil
}
