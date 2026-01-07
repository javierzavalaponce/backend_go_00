package users_presentation_gin

import (
	"fmt"
	"net/http"

	users_use_cases "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/use_cases"
	users_presentation_dtos "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/presentation/presentation_gin/dtos"
	"github.com/gin-gonic/gin"
)

type UsersHandlers struct {
	signUpUseCase *users_use_cases.SignUpUseCase
	signInUseCase *users_use_cases.SignInUseCase
}

func NewUsersHandlers(
	signUpUseCase *users_use_cases.SignUpUseCase,
	signInUseCase *users_use_cases.SignInUseCase,
) *UsersHandlers {
	return &UsersHandlers{
		signUpUseCase: signUpUseCase,
		signInUseCase: signInUseCase,
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
	// David (?):
	// al bindear el json del request al struct SignInRequest
	// se requier de validaciones adicionales?
	// o es suficiente con el binding "required" en el struct?
	var signInRequest users_presentation_dtos.SignInRequest
	if err := c.ShouldBindJSON(&signInRequest); err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	signData, err := signInRequest.ToSignInData()
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	token, err := usersHandlers.signInUseCase.Execute(signData)
	if err != nil {
		fmt.Println(err)
		//como seleccionar el status code adecuado? David
		//c.Status(http.StatusInternalServerError)
		c.Status(http.StatusUnauthorized)

		return
	}

	response := users_presentation_dtos.NewSignResponse(token)
	c.JSON(http.StatusOK, response)
}

func (usersHandlers *UsersHandlers) Profile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User profile"})
}
