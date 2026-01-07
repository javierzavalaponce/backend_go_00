package users_respositories_mock

import (
	"errors"
	"strings"

	users_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/models"
	users_respositories "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/respositories"
	"github.com/google/uuid"
)

type UsersRepositoryMock struct{}

var _ users_respositories.UsersRepository = (*UsersRepositoryMock)(nil)

func NewUsersRepositoryMock() *UsersRepositoryMock {
	return &UsersRepositoryMock{}
}

func (u *UsersRepositoryMock) Create(signData *users_models.SignData, hashedPassword string) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (u *UsersRepositoryMock) FindByEmail(email string) (string, error) {
	if strings.EqualFold(email, "javierzavalaponce@gmail.com") {
		return email, nil
	}
	return "", errors.New("user not found")
}

// FindPasswordByEmail regresa el hash de la contraseña para un email dado
func (u *UsersRepositoryMock) FindPasswordByEmail(signData *users_models.SignData) (string, error) {
	if strings.EqualFold(signData.Email, "javierzavalaponce@gmail.com") {
		if signData.Password == "123456" {
			return "$2a$10$7a8b9c0d1e2f3g4h5i6j7u8v9w0x1y2z3A4B5C6D7E8F9G0H1I2J3K", nil // bcrypt hash for "correct_password"
		}
		return "", errors.New("incorrect password")
	}
	return "", errors.New("user not found")
}
