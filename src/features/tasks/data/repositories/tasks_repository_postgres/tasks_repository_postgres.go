package tasks_repository_postgres

import (
	"database/sql"

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
	var taskId uuid.UUID
	err := u.db.QueryRow(query, task.Title, task.Description, userUUID).Scan(&taskId)
	if err != nil {
		return taskId, err
	}

	return taskId, nil
}

func (u *TaskRepositoryPostgres) FindByUserID(userID string) ([]tasks_models.Task, error) {
	query := `
		SELECT
			name,
			description,
			completed,
			user_id
		FROM test_backend.tasks
		WHERE user_id = $1
		  AND deleted_at IS NULL
		ORDER BY created_at DESC
	`

	rows, err := u.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []tasks_models.Task

	for rows.Next() {
		var task tasks_models.Task

		err := rows.Scan(
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.User_ID,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
