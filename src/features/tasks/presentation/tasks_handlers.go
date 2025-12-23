package tasks_presentation_gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TasksHandlers struct {
}

func NewTasksHandlers() *TasksHandlers {
	return &TasksHandlers{}
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
