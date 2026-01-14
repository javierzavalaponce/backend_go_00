package users_presentation_dtos

//del
import (
	"fmt"

	users_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/models"
)

type SignRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ToSignData convierte el DTO SignRequest a un modelo SignData
// siempre se valida en este punto
func (s *SignRequest) ToSignData() (*users_models.SignData, error) {
	if s.Email == "" || s.Password == "" {
		return nil, fmt.Errorf("email and password are required")
	}
	return &users_models.SignData{
		Email:    s.Email,
		Password: s.Password,
	}, nil
}

func (s *SignRequest) String() string {
	return s.Email + " " + s.Password[:3] + "..."
}
