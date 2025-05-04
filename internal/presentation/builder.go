package presentation

import (
	"github.com/gofiber/fiber/v2"
	"taskManager/internal/application"
	"taskManager/internal/presentation/user"
)

type Presentation struct {
	User *user.UserPresentation
}

func BuildPresentation(
	router fiber.Router,
	application *application.ApplicationModule,
) *Presentation {
	return &Presentation{
		User: user.NewUserPresentation(router, application),
	}
}
