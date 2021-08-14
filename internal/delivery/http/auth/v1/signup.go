package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/usename-Poezd/mysite-api/internal/services"
)

type SignUpInput struct {
	Name string `json:"name"  validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Email string `json:"email" validate:"required,email,min=6,max=32"`
}

// SignUp
// @Summary User SignUp
// @Tags users-auth
// @Description create user account
// @ModuleID SignUp
// @Accept  json
// @Produce  json
// @Param input body SignUpInput true "sign up info"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} SignUpInput
// @Failure 422 {object} ErrorResponse
// @Failure 500 {object} SignUpInput
// @Failure default {object} SignUpInput
// @Router /auth/v1/sign-up [post]
func (h Handler) SignUp(c *fiber.Ctx) error  {
	input := new(SignUpInput)
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	var errors []*ErrorResponse
	if err := validator.New().Struct(input); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	if errors != nil {
		return c.Status(422).JSON(errors)
	}

	err := h.services.User.SignUp(services.UserSignUpInput{
		Name: input.Name,
		Email: input.Email,
		Password: input.Password,
	})
	if err != nil {
		return err
	}

	return c.JSON(input)
}
