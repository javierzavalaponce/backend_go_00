package main

import (
	users_respositories_mock "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/data/repositories/users_repository/users_repository_mock"
	users_security_service_mock "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/data/services/security_service/security_service_mock"
	users_use_cases "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/use_cases"
	users_presentation_gin "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/presentation/presentation_gin"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	securityService := users_security_service_mock.NewSecurityServiceMock()
	usersRepository := users_respositories_mock.NewUsersRepositoryMock()

	signUpUseCase := users_use_cases.NewSignUpUseCase(
		usersRepository,
		securityService,
	)

	userHandlers := users_presentation_gin.NewUsersHandlers(signUpUseCase)

	users_presentation_gin.SetupRoutes(r, userHandlers)

	r.Run(":3000")
}
