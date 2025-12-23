package users_presentation_dtos

import (
	users_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/models"
)

type SignRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *SignRequest) ToSignData() (*users_models.SignData, error) {
	return &users_models.SignData{
		Email:    s.Email,
		Password: s.Password,
	}, nil
}

func (s *SignRequest) String() string {
	return s.Email + " " + s.Password[:3] + "..."
}
