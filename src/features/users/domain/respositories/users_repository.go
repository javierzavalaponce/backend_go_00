package users_respositories

import (
	"github.com/google/uuid"
)

type UsersRepository interface {
	Create(email, hashedPassword string) (uuid.UUID, error)
}
