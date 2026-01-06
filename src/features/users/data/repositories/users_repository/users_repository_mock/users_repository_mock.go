package users_respositories_mock

import (
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
	//David(?): no stoy seguro aqui..
	//	mock returning the email as signData for simplicity
	return email, nil

}

func (u *UsersRepositoryMock) VerifyPassword(password string, hashedPassword string) (bool, error) {
	return true, nil
}
