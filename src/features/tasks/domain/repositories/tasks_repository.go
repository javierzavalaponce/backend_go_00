package task_repository

import tasks_models

type TaskRepository interface {
	Save(task *tasks_models.Task) error
	FindAll() ([]tasks_models.Task, error)
	FindByID(id int) (*tasks_models.Task, error)
	Update(task *tasks_models.Task) error
	Delete(id int) error
}
