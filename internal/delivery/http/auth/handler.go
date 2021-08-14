package auth

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/usename-Poezd/mysite-api/internal/delivery/http/auth/v1"
	"github.com/usename-Poezd/mysite-api/internal/services"
)

type Handler struct {
	services 	*services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services,
	}
}

func (h Handler) Init(router fiber.Router) {
	api := router.Group("auth")

	v1.NewHandler(h.services).Init(api)
}
