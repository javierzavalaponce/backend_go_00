package tasks_presentation_gin

import tasks_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/models"

type CreateTaskRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (s *CreateTaskRequest) ToCreateTaskData() (*tasks_models.Task, error) {

	return &tasks_models.Task{
		Title:       s.Name,
		Description: s.Description,
		Completed:   false,
		User_ID:     "",
	}, nil

}
