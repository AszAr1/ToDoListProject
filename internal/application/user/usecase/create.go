package usecase

import (
	"net/http"
	"taskManager/internal/application/user/dto/request"
	"taskManager/internal/application/user/mapper"
	"taskManager/internal/domain/lib"
	"taskManager/internal/domain/user"
)

type CreateUserUseCase struct {
	rep user.IUserRepository
}

func NewCreateUserUseCase(rep user.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{rep: rep}
}

func (receiver CreateUserUseCase) Invoke(dto *request.RequestCreateUser) *lib.Response {
	model := mapper.MapUserFromRequestToModel(dto)

	if err := receiver.rep.Create(*model); err != nil {
		return &lib.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &lib.Response{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
	}
}
