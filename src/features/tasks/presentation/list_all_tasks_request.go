package tasks_presentation_gin

type ListAllTaskRequest struct {
	Token string `json:"token_jwt" binding:"required"`
}

func NewListAllTaskData(token string) *ListAllTaskRequest {
	return &ListAllTaskRequest{Token: token}
}
