package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"taskManager/internal/domain/lib"
	"taskManager/internal/infrastruct/errors"
)

const ParseData = "parsed_data"

var validate = validator.New()

func ParseAndValidateData[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var value T
		if err := c.BodyParser(&value); err != nil {
			return c.Status(http.StatusBadRequest).JSON(lib.Response{
				Code:    http.StatusBadRequest,
				Message: "Invalid request body" + err.Error(),
			})
		}

		if err := validate.Struct(value); err != nil {
			return c.Status(http.StatusBadRequest).JSON(lib.Response{
				Code:    http.StatusBadRequest,
				Message: "Failed to validate request body",
				Data:    errors.GetValidationErrors(err.(validator.ValidationErrors)),
			})
		}

		c.Locals(ParseData, value)
		return c.Next()
	}
}
