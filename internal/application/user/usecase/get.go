package usecase

import (
	"net/http"
	"taskManager/internal/application/user/mapper"
	"taskManager/internal/domain/lib"
	"taskManager/internal/domain/user"
)

type GetUserUseCase struct {
	rep user.IUserRepository
}

func NewGetUserUseCase(rep user.IUserRepository) *GetUserUseCase {
	return &GetUserUseCase{rep: rep}
}

func (uc GetUserUseCase) Invoke(id int) *lib.Response {
	res, err := uc.rep.GetUser(user.GetUserParams{
		Id:   &id,
		Type: user.ById,
	})

	if err != nil {
		return &lib.Response{
			Code:    http.StatusInternalServerError,
			Message: "Fetch error" + err.Error(),
		}
	}

	return &lib.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    mapper.MapUserFromModelToResponse(res),
	}
}
