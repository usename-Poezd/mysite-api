package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/usename-Poezd/mysite-api/internal/domain"
	"github.com/usename-Poezd/mysite-api/internal/repository/postgresql"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}

type Repositories struct {
	Users UserRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users: postgresql.NewUsersRepo(db),
	}
}