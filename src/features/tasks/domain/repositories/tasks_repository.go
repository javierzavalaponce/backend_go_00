package task_repository

import tasks_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/models"

type TaskRepository interface {
	Save(task *tasks_models.Task) error
	FindAll() ([]tasks_models.Task, error)
	FindByID(id int) (*tasks_models.Task, error)
	Update(task *tasks_models.Task) error
	Delete(id int) error
}
