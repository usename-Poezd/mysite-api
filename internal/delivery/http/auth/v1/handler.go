package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/usename-Poezd/mysite-api/internal/services"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services,
	}
}

func (h *Handler) Init(router fiber.Router)  {
	v1 := router.Group("v1")

	v1.Post("/sign-up", h.SignUp)
	v1.Post("/sign-in", h.SignIn)
}
