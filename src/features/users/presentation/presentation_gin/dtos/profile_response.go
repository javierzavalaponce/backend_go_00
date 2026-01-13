package users_presentation_dtos

import users_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/models"

type ProfileOutputDto struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func ProfileOutputDtoFromUser(user users_models.User) ProfileOutputDto {
	return ProfileOutputDto{
		ID:    user.ID,
		Email: user.Email,
	}
}
