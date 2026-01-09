package users_repository_postgres

import (
	"database/sql"

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
