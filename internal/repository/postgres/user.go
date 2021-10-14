package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mihail-1212/go-project-layout/pkg/domain"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (u *UserPostgres) GetAllUsers() ([]domain.User, error) {
	var users []domain.User

	query := fmt.Sprintf("SELECT id, username, \"password\" FROM %s", userTable)
	rows, err := u.db.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
