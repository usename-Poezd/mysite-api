package v1

import (
"github.com/go-playground/validator/v10"
"github.com/gofiber/fiber/v2"
	"github.com/usename-Poezd/mysite-api/internal/domain"
	"github.com/usename-Poezd/mysite-api/internal/services"
	"github.com/usename-Poezd/mysite-api/pkg/auth"
)

type SignInInput struct {
	Password string `json:"password" validate:"required,min=6,required"`
	Email string `json:"email" validate:"required,email,min=6,max=32,required"`
}

type SignInResponse struct {
	Token 	string		`json:"token"`
	Data	*domain.User `json:"data"`
}

// SignIn
// @Summary User SignIn
// @Tags users-auth
// @Description sign in user
// @ModuleID SignUp
// @Accept  json
// @Produce  json
// @Param input body SignInInput true "sign in info"
// @Success 200 {object} SignInResponse
// @Failure 400,404 {object} SignInInput
// @Failure 422 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} SignInInput
// @Router /auth/v1/sign-in [post]
func (h Handler) SignIn(c *fiber.Ctx) error  {
	input := new(SignInInput)
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

	user, err := h.services.User.SignIn(services.UserSignInInput{
		Email: input.Email,
		Password: input.Password,
	})
	if err != nil {
		return err
	}

	token, err := auth.MakeToken(user)

	if err != nil {
		return err
	}

	return c.JSON(SignInResponse{
		token,
		user,
	})
}

