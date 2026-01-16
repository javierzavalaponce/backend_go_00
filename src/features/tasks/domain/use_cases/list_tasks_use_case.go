package tasks_use_cases

import (
	core_services "github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/domain/services"
	tasks_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/models"
	tasks_repository "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/repositories"
)

type ListTasksUseCase struct {
	tasksRepository tasks_repository.TaskRepository
	securityService core_services.SecurityService
}

func NewListTasksUseCase(
	tasksRepository tasks_repository.TaskRepository,
	securityService core_services.SecurityService,
) *ListTasksUseCase {
	return &ListTasksUseCase{
		tasksRepository: tasksRepository,
		securityService: securityService,
	}
}

func (t *ListTasksUseCase) Execute(token string) (*[]tasks_models.Task, error) {
	userID, err := t.securityService.ValidateToken(token)
	if err != nil {
		return nil, err
	}
	listTasks, err := t.tasksRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	return &listTasks, nil
}
