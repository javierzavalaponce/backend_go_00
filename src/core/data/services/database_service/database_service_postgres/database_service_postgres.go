package database_service_postgres

import (
	"database/sql"
	"fmt"
	"log"

	core_services "github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/domain/services"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DatabaseServicePostgres struct {
	DB *sql.DB
}

func NewDatabaseServicePostgres(host, port, user, password, dbname, sslMode, migrationsPath string) (core_services.DatabaseService, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: "schema_migrations",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		dbname,
		driver,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize migration instance: %w", err)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No new migrations to apply.")
		} else if err.Error() == "dirty database" {
			return nil, fmt.Errorf("database is dirty. Run `migrate force <version>` to fix: %w", err)
		} else {
			return nil, fmt.Errorf("migration failed: %w", err)
		}
	} else {
		log.Println("Database migrations applied successfully.")
	}

	return &DatabaseServicePostgres{DB: db}, nil
}

func (s *DatabaseServicePostgres) GetDB() *sql.DB {
	return s.DB
}

func (s *DatabaseServicePostgres) Close() error {
	return s.DB.Close()
}
