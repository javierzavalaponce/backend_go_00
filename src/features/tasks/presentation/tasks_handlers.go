package tasks_presentation_gin

import (
	"fmt"
	"net/http"

	tasks_use_cases "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/tasks/domain/use_cases"
	"github.com/gin-gonic/gin"
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
	/*use mock repo...*/
	fmt.Println("Obtiene la lista -ALL-")
	c.JSON(http.StatusOK, "list")
}

// READ (one)
func (h *TasksHandlers) getTaskByID(c *gin.Context) {
	/*use mock repo...*/
	fmt.Println("Obtiene un item por ID")
	c.JSON(http.StatusOK, "task")
}

// javier
func (h *TasksHandlers) createTask(c *gin.Context) {
	var createTaskRequest CreateTaskRequest
	if err := c.ShouldBindJSON(&createTaskRequest); err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	req, err := createTaskRequest.ToCreateTaskData()

	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	authHeader := c.GetHeader("Authorization")
	req.User_ID = authHeader[len("Bearer "):]
	if authHeader == "" {
		c.Status(http.StatusUnauthorized)
		return
	}
	token := authHeader[len("Bearer "):]
	err = h.createTaskUseCase.Execute(req, token)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, "created")
}
