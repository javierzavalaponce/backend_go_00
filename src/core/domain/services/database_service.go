package core_services

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DatabaseService interface {
	Close() error
	GetDB() *sql.DB
}
