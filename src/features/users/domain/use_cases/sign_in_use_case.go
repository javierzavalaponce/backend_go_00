package users_use_cases

import (
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
	// Buscar el usuario en el repositorio
	userId, err := s.usersRepository.FindByEmail(signData.Email)
	if err != nil {
		return "", err
	}

	// Verificar la contraseña
	/*
		isPasswordValid, err := s.securityService.VerifyPassword(signData.Password, userId)
		if err != nil {
			return "", err
		}
		if !isPasswordValid {
			return "", err.New("invalid credentials")
		}
	*/
	// Generar el token de autenticación
	token, err := s.securityService.GenerateToken(userId)
	if err != nil {
		return "", err
	}

	return token, nil
}
