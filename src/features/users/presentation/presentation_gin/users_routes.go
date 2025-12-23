package users_presentation_gin

import "github.com/gin-gonic/gin"

func SetupRoutes(r gin.IRouter, userHandlers *UsersHandlers) {
	usersRoute := r.Group("/users")
	usersRoute.POST("/sign-up", userHandlers.SignUp)
	usersRoute.POST("/sign-in", userHandlers.SignIn)
	usersRoute.GET("/profile", userHandlers.Profile)
}
