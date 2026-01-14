package main

import (
	"github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/data/services/config_service"
	"github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/data/services/database_service/database_service_postgres"
	"github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/data/services/logging_service"

	security_service_impl "github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/data/services/security_service"

	tasks_repository_postgres "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/data/repositories/tasks_repository_postgres"
	tasks_use_cases "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/use_cases"
	tasks_presentation_gin "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/presentation"
	users_respositories_postgres "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/data/repositories/users_repository/users_repository_postgres"
	users_use_cases "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/use_cases"
	users_presentation_gin "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/presentation/presentation_gin"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfgService := config_service.NewConfigServiceEnv()
	loggingService := logging_service.NewLoggingServiceGolog()
	cfg := cfgService.Read()

	databaseService, err := database_service_postgres.NewDatabaseServicePostgres(
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbName,
		cfg.DbSslMode,
		cfg.MigrationsPath,
	)

	if databaseService == nil {
		dbErr := err
		loggingService.Error("Failed to initialize database", map[string]interface{}{
			"error": dbErr,
		})
		return
	}

	if err != nil {
		panic(err)
	}

	defer databaseService.Close()

	securityService := security_service_impl.NewSecurityServiceImpl(cfg.SessionSecret)
	usersRepository := users_respositories_postgres.NewUsersRepositoryPostgres(databaseService.GetDB())
	tasksRepository := tasks_repository_postgres.NewTasksRepositoryPostgres(databaseService.GetDB())

	signUpUseCase := users_use_cases.NewSignUpUseCase(
		usersRepository,
		securityService,
	)

	signInUseCase := users_use_cases.NewSignInUseCase(
		usersRepository,
		securityService,
	)

	getProfileUseCase := users_use_cases.NewGetProfileUseCase(
		usersRepository,
		securityService,
	)

	userHandlers := users_presentation_gin.NewUsersHandlers(
		signUpUseCase,
		signInUseCase,
		getProfileUseCase,
	)

	// Tasks Handlers
	createTaskUseCase := tasks_use_cases.NewCreateTaskUseCase(
		tasksRepository,
		securityService,
	)

	tasksHandlers := tasks_presentation_gin.NewTasksHandlers(
		createTaskUseCase,
	)

	users_presentation_gin.SetupRoutes(r, userHandlers)
	tasks_presentation_gin.SetupRoutes(r, tasksHandlers)

	r.Run(":3000")
}
