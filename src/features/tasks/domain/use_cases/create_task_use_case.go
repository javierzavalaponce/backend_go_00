package tasks_use_cases

import (
	core_services "github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/domain/services"
	tasks_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/models"
	tasks_repository "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/repositories"
	"github.com/google/uuid"
)

type CreateTasksUseCase struct {
	tasksRepository tasks_repository.TaskRepository
	securityService core_services.SecurityService
}

func NewCreateTaskUseCase(
	tasksRepository tasks_repository.TaskRepository,
	securityService core_services.SecurityService,
) *CreateTasksUseCase {
	return &CreateTasksUseCase{
		tasksRepository: tasksRepository,
		securityService: securityService,
	}
}

func (t *CreateTasksUseCase) Execute(request *tasks_models.Task, token string) (uuid.UUID, error) {
	var id uuid.UUID
	userID, err := t.securityService.ValidateToken(token)
	if err != nil {
		return id, err
	}
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return id, err
	}
	id, err = t.tasksRepository.Create(request, userUUID)
	if err != nil {
		return id, err
	}
	return id, nil
}
