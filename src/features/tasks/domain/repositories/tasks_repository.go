package task_repository

import (
	tasks_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/models"
	"github.com/google/uuid"
)

type TaskRepository interface {
	Create(task *tasks_models.Task, userUUID uuid.UUID) (uuid.UUID, error)
	FindByUserID(userID string) ([]tasks_models.Task, error)
}
