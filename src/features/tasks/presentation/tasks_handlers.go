package tasks_presentation_gin

import (
	"fmt"
	"net/http"

	tasks_use_cases "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/use_cases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TasksHandlers struct {
	createTaskUseCase *tasks_use_cases.CreateTasksUseCase
}

func NewTasksHandlers(
	createTaskUseCase *tasks_use_cases.CreateTasksUseCase,
) *TasksHandlers {
	return &TasksHandlers{
		createTaskUseCase: createTaskUseCase,
	}
}

func (h *TasksHandlers) getTasks(c *gin.Context) {
	fmt.Println("Obtiene la lista -ALL-")
	c.JSON(http.StatusOK, "list")
}

func (h *TasksHandlers) getTaskByID(c *gin.Context) {
	fmt.Println("Obtiene un item por ID")
	c.JSON(http.StatusOK, "task")
}

func (h *TasksHandlers) createTask(c *gin.Context) {
	var createTaskRequest CreateTaskRequest
	if err := c.ShouldBindJSON(&createTaskRequest); err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}
	httpRequest, err := createTaskRequest.TaskData()
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}
	authHeader := c.GetHeader("Authorization")
	httpRequest.User_ID = authHeader[len("Bearer "):]
	if authHeader == "" {
		c.Status(http.StatusUnauthorized)
		return
	}
	token := authHeader[len("Bearer "):]
	var id uuid.UUID
	id, err = h.createTaskUseCase.Execute(httpRequest, token)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, id.String())
}
