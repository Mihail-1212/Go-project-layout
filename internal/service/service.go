package service

import (
	repo "github.com/mihail-1212/go-project-layout/internal/repository"
	"github.com/mihail-1212/go-project-layout/pkg/domain"
)

type UserService interface {
	GetAllUsers() ([]domain.User, error)
}

type Services struct {
	UserService
}

func NewServices(repo *repo.Repository) *Services {
	return &Services{
		UserService: NewUser(repo),
	}
}
