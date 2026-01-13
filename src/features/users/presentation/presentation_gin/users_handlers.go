package users_presentation_gin

import (
	"fmt"
	"net/http"

	users_use_cases "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/use_cases"
	users_presentation_dtos "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/presentation/presentation_gin/dtos"
	"github.com/gin-gonic/gin"
)

type UsersHandlers struct {
	signUpUseCase     *users_use_cases.SignUpUseCase
	signInUseCase     *users_use_cases.SignInUseCase
	getProfileUseCase *users_use_cases.GetProfileUseCase
}

func NewUsersHandlers(
	signUpUseCase *users_use_cases.SignUpUseCase,
	signInUseCase *users_use_cases.SignInUseCase,
	getProfileUseCase *users_use_cases.GetProfileUseCase,
) *UsersHandlers {
	return &UsersHandlers{
		signUpUseCase:     signUpUseCase,
		signInUseCase:     signInUseCase,
		getProfileUseCase: getProfileUseCase,
	}
}

func (usersHandlers *UsersHandlers) SignUp(c *gin.Context) {
	var signRequest users_presentation_dtos.SignRequest
	if err := c.ShouldBindJSON(&signRequest); err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	signData, err := signRequest.ToSignData()
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	token, err := usersHandlers.signUpUseCase.Execute(signData)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	response := users_presentation_dtos.NewSignResponse(token)
	c.JSON(http.StatusCreated, response)
}

func (usersHandlers *UsersHandlers) SignIn(c *gin.Context) {
	var signInRequest users_presentation_dtos.SignRequest
	if err := c.ShouldBindJSON(&signInRequest); err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	signData, err := signInRequest.ToSignData()
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	token, err := usersHandlers.signInUseCase.Execute(signData)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusUnauthorized)
		return
	}

	response := users_presentation_dtos.NewSignResponse(token)
	c.JSON(http.StatusOK, response)
}

func (usersHandlers *UsersHandlers) Profile(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		/*
			h.LoggingService.Error("GetProfile error: ", map[string]interface{}{
				"error": "Authorization header is required",
			})
		*/
		c.Status(http.StatusUnauthorized)
		return
	}

	token := authHeader[len("Bearer "):]
	user, err := usersHandlers.getProfileUseCase.Execute(token, nil)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, users_presentation_dtos.ProfileOutputDtoFromUser(*user))
	//c.JSON(http.StatusOK, gin.H{"message": "User profile"})
}
