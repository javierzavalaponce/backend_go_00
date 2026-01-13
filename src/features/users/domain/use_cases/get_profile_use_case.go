package users_use_cases

import (
	core_services "github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/domain/services"
	users_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/models"
	users_respositories "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/respositories"
)

type GetProfileUseCase struct {
	usersRepository users_respositories.UsersRepository
	securityService core_services.SecurityService
}

func NewGetProfileUseCase(
	usersRepository users_respositories.UsersRepository,
	securityService core_services.SecurityService,
) *GetProfileUseCase {
	return &GetProfileUseCase{
		usersRepository: usersRepository,
		securityService: securityService,
	}
}

func (s *GetProfileUseCase) Execute(token string, signData *users_models.SignData) (*users_models.User, error) {
	// Validar el token de autenticación
	userID, err := s.securityService.ValidateToken(token)
	if err != nil {
		return &users_models.User{}, err
	}

	// 1. Buscar el usuario en el repositorio
	user, err := s.usersRepository.Read(userID)
	if err != nil {
		return &users_models.User{}, err
	}
	return user, err
}
