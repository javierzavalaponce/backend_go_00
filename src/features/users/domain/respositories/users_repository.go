package users_respositories

import (
	users_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/models"
	"github.com/google/uuid"
)

type UsersRepository interface {
	Create(user *users_models.SignData, hashedPassword string) (uuid.UUID, error)
}
