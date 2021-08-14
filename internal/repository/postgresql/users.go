package postgresql

import (
	"github.com/jmoiron/sqlx"
	"github.com/usename-Poezd/mysite-api/internal/domain"
	"time"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{
		db,
	}
}

func (r UsersRepository) Create(user *domain.User) (*domain.User, error)  {
	user.UpdatedAt = time.Now()
	_, err := r.db.NamedExec(`INSERT INTO users (name, password, email, updated_at) VALUES (:name, :password, :email, :updated_at)`, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UsersRepository) FindByEmail(email string) (*domain.User, error)  {
	user := domain.User{}
	if err := r.db.Get(&user, "SELECT * FROM users WHERE email=$1 LIMIT 1", email); err != nil {
		return nil, err
	}

	return &user, nil
}
