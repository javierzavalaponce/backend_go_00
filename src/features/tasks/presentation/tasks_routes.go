package tasks_presentation_gin

import "github.com/gin-gonic/gin"

func SetupRoutes(r gin.IRouter, tasksHandlers *TasksHandlers) {
	tasksRoute := r.Group("/tasks")
	tasksRoute.GET("/:id", tasksHandlers.getTaskByID)
	tasksRoute.GET("/", tasksHandlers.getTasks)
}
