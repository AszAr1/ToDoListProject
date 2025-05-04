package user

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"taskManager/internal/application/user"
	"taskManager/internal/application/user/dto/request"
	"taskManager/internal/domain/lib"
	"taskManager/internal/infrastruct/middleware"
)

func create(ctx *fiber.Ctx, uc user.ICreateUserUseCase) error {
	parsedUser := ctx.Locals(middleware.ParseData).(request.RequestCreateUser)
	res := uc.Invoke(&parsedUser)
	if res.Message != "" {
		return ctx.Status(res.Code).JSON(res)
	}
	return ctx.Status(res.Code).JSON(res)
}

func getUser(ctx *fiber.Ctx, uc user.IGetUserUseCase) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(lib.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid id",
		})
	}

	res := uc.Invoke(id)
	if res.Message != "" {
		return ctx.Status(res.Code).JSON(res)
	}

	return ctx.Status(res.Code).JSON(res)
}

func getAllUsers(ctx *fiber.Ctx, uc user.IGetAllUsersUseCase) error {
	res := uc.GetAll()

	if res.Message != "" {
		return ctx.Status(res.Code).JSON(res)
	}

	return ctx.Status(res.Code).JSON(res)
}

func deleteUser(ctx *fiber.Ctx, uc user.IDeleteUserUseCase) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(lib.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid id",
		})
	}

	res := uc.Invoke(id)
	if res.Message != "" {
		return ctx.Status(res.Code).JSON(res)
	}

	return ctx.Status(res.Code).JSON(res)
}
