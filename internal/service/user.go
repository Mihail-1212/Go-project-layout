package service

import (
	repo "github.com/mihail-1212/go-project-layout/internal/repository"
	"github.com/mihail-1212/go-project-layout/pkg/domain"
)

type User struct {
	repo *repo.Repository
}

func NewUser(repo *repo.Repository) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) GetAllUsers() ([]domain.User, error) {
	return u.repo.User.GetAllUsers()
}
