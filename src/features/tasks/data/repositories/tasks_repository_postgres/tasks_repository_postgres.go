package tasks_repository_postgres

import (
	"database/sql"
	"fmt"

	tasks_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/models"
	task_repository "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/repositories"
	"github.com/google/uuid"
)

type TaskRepositoryPostgres struct {
	db *sql.DB
}

func NewTasksRepositoryPostgres(db *sql.DB) task_repository.TaskRepository {
	return &TaskRepositoryPostgres{
		db: db,
	}
}

func (u *TaskRepositoryPostgres) Create(task *tasks_models.Task, userUUID uuid.UUID) (uuid.UUID, error) {
	query := `INSERT INTO test_backend.tasks (name, description, user_id) VALUES ($1, $2, $3) RETURNING id`

	userUUID, e := uuid.Parse(task.User_ID)
	if e != nil {
		return fmt.Errorf("user_id inválido: %w", e)
	}

	var id uuid.UUID
	err := u.db.QueryRow(query, task.Title, task.Description, userUUID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
