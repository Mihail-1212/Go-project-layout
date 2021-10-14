package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mihail-1212/go-project-layout/internal/repository/postgres"
	"github.com/mihail-1212/go-project-layout/pkg/domain"
)

type User interface {
	GetAllUsers() ([]domain.User, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: postgres.NewUserPostgres(db),
	}
}
