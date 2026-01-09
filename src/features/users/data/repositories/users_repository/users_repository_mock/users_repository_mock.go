package users_respositories_mock

import (
	users_respositories "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/respositories"
	"github.com/google/uuid"
)

type UsersRepositoryMock struct{}

func NewUsersRepositoryMock() users_respositories.UsersRepository {
	return &UsersRepositoryMock{}
}

func (u *UsersRepositoryMock) Create(email, hashedPassword string) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (u *UsersRepositoryMock) FindByEmail(email string) (*users_models.User, error) {
	if !strings.EqualFold(email, "javierzavalaponce@gmail.com") {
		return nil, errors.New("user not found")
	}
	return &users_models.User{
		ID:             "123456789",
		Email:          email,
		HashedPassword: "$2a$10$7a8b9c0d1e2f3g4h5i6j7u8v9w0x1y2z3A4B5C6D7E8F9G0H1I2J3K",
	}, nil

}
