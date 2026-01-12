package users_repository_postgres

import (
	"database/sql"

	users_models "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/models"
	users_respositories "github.com/DEINSI-DEVELOP/test_backend_go.git/src/features/users/domain/respositories"
	"github.com/google/uuid"
)

type UsersRepositoryPostgres struct {
	db *sql.DB
}

func NewUsersRepositoryPostgres(db *sql.DB) users_respositories.UsersRepository {
	return &UsersRepositoryPostgres{
		db: db,
	}
}

func (u *UsersRepositoryPostgres) Create(email, hashedPassword string) (uuid.UUID, error) {
	query := `INSERT INTO test_backend.users (email, password) VALUES ($1, $2) RETURNING id`
	var id uuid.UUID
	err := u.db.QueryRow(query, email, hashedPassword).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (u *UsersRepositoryPostgres) FindByEmail(email string) (*users_models.User, error) {

	query := `SELECT id, email, password FROM test_backend.users WHERE email = $1`
	var user users_models.User
	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.HashedPassword)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
