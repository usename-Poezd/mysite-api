package http

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/usename-Poezd/mysite-api/docs"
	"github.com/usename-Poezd/mysite-api/internal/delivery/http/auth"
	v1 "github.com/usename-Poezd/mysite-api/internal/delivery/http/v1"
	"github.com/usename-Poezd/mysite-api/internal/services"
)

type Handler struct {
	app 		*fiber.App
	services 	*services.Services
}

func NewHandler(app *fiber.App, services *services.Services) *Handler {
	return &Handler{
		app,
		services,
	}
}

func (h Handler) Init() {

	h.app.Get("/swagger/*", swagger.Handler) // default

	api := h.app.Group("api")


	v1.NewHandler(h.services).Init(api)
	auth.NewHandler(h.services).Init(api)
}
