package user

import (
	"taskManager/internal/application/user/dto/request"
	"taskManager/internal/application/user/usecase"
	"taskManager/internal/domain/lib"
	"taskManager/internal/domain/user"
)

type ICreateUserUseCase interface {
	Invoke(dto *request.RequestCreateUser) *lib.Response
}

type IGetUserUseCase interface {
	Invoke(id int) *lib.Response
}

type IDeleteUserUseCase interface {
	Invoke(id int) *lib.Response
}

type IGetAllUsersUseCase interface {
	GetAll() *lib.Response
}

type UserApplication struct {
	Create ICreateUserUseCase
	Get    IGetUserUseCase
	Delete IDeleteUserUseCase
	GetAll IGetAllUsersUseCase
}

func BuildUserApplication(rep user.IUserRepository) *UserApplication {
	return &UserApplication{
		Create: usecase.NewCreateUserUseCase(rep),
		Get:    usecase.NewGetUserUseCase(rep),
		Delete: usecase.NewDeleteUserUseCase(rep),
		GetAll: usecase.NewGetAllUsersUseCase(rep),
	}
}
