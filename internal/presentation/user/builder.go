package user

import (
	"github.com/gofiber/fiber/v2"
	"taskManager/internal/application"
	"taskManager/internal/application/user/dto/request"
	"taskManager/internal/infrastruct/middleware"
)

type UserPresentation struct {
	router      fiber.Router
	application *application.ApplicationModule
}

func NewUserPresentation(router fiber.Router,
	application *application.ApplicationModule) *UserPresentation {
	return &UserPresentation{
		router:      router,
		application: application,
	}
}

func (u UserPresentation) BuildUserPresentation() {
	route := u.router.Group(KEY)
	route.Name("User endpoints")

	route.Post(CREATE, middleware.ParseAndValidateData[request.RequestCreateUser](), func(ctx *fiber.Ctx) error {
		return create(ctx, u.application.User.Create)
	})

	route.Get(GET_ALL, func(ctx *fiber.Ctx) error {
		return getAllUsers(ctx, u.application.User.GetAll)
	})

	route.Get(GET, func(ctx *fiber.Ctx) error {
		return getUser(ctx, u.application.User.Get)
	})

	route.Delete(DELETE, func(ctx *fiber.Ctx) error {
		return deleteUser(ctx, u.application.User.Delete)
	})
}
