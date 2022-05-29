package database

import (
	"database/sql"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	if len(url) == 0 {
		var err error
		url, err = loadServerUrl()

		if err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func loadServerUrl() (string, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"postgres://%s:%s@postgres/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
	), nil
}

func (repo *PostgresRepository) Close() {
	repo.db.Close()
}
