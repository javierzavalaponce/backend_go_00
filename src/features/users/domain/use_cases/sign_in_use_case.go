package users_use_cases

import (
	"errors"

	users_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/models"
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

// Sign in, inicio de sesion
func (s *SignInUseCase) Execute(signData *users_models.SignData) (string, error) {
	// 1. Buscar el usuario en el repositorio
	user, err := s.usersRepository.FindByEmail(signData.Email)
	if err != nil {
		return "", err
	}

	// Verificar la contraseña
	isPasswordValid, err := s.securityService.VerifyPassword(user.HashedPassword, signData.Password)
	if err != nil {
		return "", err
	}
	if !isPasswordValid {
		return "", errors.New("invalid credentials")
	}

	// Generar el token de autenticación
	token, err := s.securityService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
