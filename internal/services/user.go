package services

import (
	"errors"
	"github.com/usename-Poezd/mysite-api/internal/domain"
	"github.com/usename-Poezd/mysite-api/internal/repository"
	"github.com/usename-Poezd/mysite-api/pkg/hasher"
)

type UserSignUpInput struct {
	Name     string
	Email    string
	Password string
}

type UserSignInInput struct {
	Email    string
	Password string
}


type UserService interface {
	SignUp(input UserSignUpInput) error
	SignIn(input UserSignInInput) (*domain.User, error)
}

type User struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *User {
	return &User{
		repo,
	}
}

func (s User) SignUp(input UserSignUpInput) error  {
	user := &domain.User{
		Name: input.Name,
		Email: input.Email,
		Password: hasher.HashAndSalt([]byte(input.Password)),
	}

	_, err := s.repo.Create(user)
	return err
}

func (s User) SignIn(input UserSignInInput) (*domain.User, error)  {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if !hasher.ComparePasswords(user.Password, []byte(input.Password)) {
		return nil, errors.New("passwords dont match")
	}

	return user, nil
}