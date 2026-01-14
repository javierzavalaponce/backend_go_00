package tasks_models

type Task struct {
	Title       string `json:"title"`
	Description string `json:"descr"`
	Completed   bool   `json:"completed"`
	User_ID     string `json:"user-id"`
}
