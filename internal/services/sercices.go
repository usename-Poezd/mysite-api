package services

import "github.com/usename-Poezd/mysite-api/internal/repository"

type Services struct {
	User	UserService
}

func NewServices(repos *repository.Repositories) *Services  {
	return &Services{
		User: NewUserService(repos.Users),
	}
}