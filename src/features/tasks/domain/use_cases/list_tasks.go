package tasks_use_cases

type ListTasksUseCase struct {
	// Add necessary repositories and services here
}

func (t *ListTasksUseCase) ListTasks() ([]string, error) {
	// Implement the logic to list tasks here
	return []string{"task1", "task2"}, nil
}
