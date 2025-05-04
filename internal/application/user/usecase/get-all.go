package usecase

import (
	"net/http"
	"taskManager/internal/application/user/dto/response"
	"taskManager/internal/application/user/mapper"
	"taskManager/internal/domain/lib"
	"taskManager/internal/domain/user"
)

type GetAllUsersUseCase struct {
	rep user.IUserRepository
}

func NewGetAllUsersUseCase(rep user.IUserRepository) *GetAllUsersUseCase {
	return &GetAllUsersUseCase{rep: rep}
}

func (u GetAllUsersUseCase) GetAll() *lib.Response {
	res, err := u.rep.GetAll()

	if err != nil {
		return &lib.Response{
			Code:    http.StatusInternalServerError,
			Message: "Fetch all users failed" + err.Error(),
		}
	}

	var users []response.ResponseUser

	for _, user := range res {
		users = append(users, *mapper.MapUserFromModelToResponse(&user))
	}

	return &lib.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    users,
	}
}
