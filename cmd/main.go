package main

import (
	"github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/data/services/database_service/database_service_postgres"
	users_respositories_postgres "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/data/repositories/users_repository/users_repository_postgres"
	users_security_service_mock "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/data/services/security_service/security_service_mock"
	users_use_cases "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/use_cases"
	users_presentation_gin "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/presentation/presentation_gin"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	databaseService, err := database_service_postgres.NewDatabaseServicePostgres(
		"localhost",
		"5432",
		"USER",
		"PASSWORD",
		"TEST_BACKEND_DB",
		"disable",
		"src/core/data/services/database_service/database_service_postgres/migrations",
	)
	if err != nil {
		panic(err)
	}

	defer databaseService.Close()

	securityService := users_security_service_mock.NewSecurityServiceMock()
	usersRepository := users_respositories_postgres.NewUsersRepositoryPostgres(databaseService.GetDB())

	signUpUseCase := users_use_cases.NewSignUpUseCase(
		usersRepository,
		securityService,
	)

	signInUseCase := users_use_cases.NewSignInUseCase(
		usersRepository,
		securityService,
	)

	userHandlers := users_presentation_gin.NewUsersHandlers(
		signUpUseCase,
		signInUseCase,
	)

	tasksHandlers := tasks_presentation_gin.NewTasksHandlers()

	users_presentation_gin.SetupRoutes(r, userHandlers)
	tasks_presentation_gin.SetupRoutes(r, tasksHandlers)

	r.Run(":3000")
}
