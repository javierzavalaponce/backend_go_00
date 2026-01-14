package task_repository

import tasks_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/models"

type TaskRepository interface {
	Create(task *tasks_models.Task) error
}
