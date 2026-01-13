package users_respositories

import (
	users_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/models"
	"github.com/google/uuid"
)

type UsersRepository interface {
	Create(email, hashedPassword string) (uuid.UUID, error)
	FindByEmail(email string) (*users_models.User, error)
	Read(id string) (*users_models.User, error)
}
