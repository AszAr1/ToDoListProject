package usecase

import (
	"net/http"
	"taskManager/internal/domain/lib"
	"taskManager/internal/domain/user"
)

type DeleteUserUseCase struct {
	rep user.IUserRepository
}

func NewDeleteUserUseCase(rep user.IUserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{rep: rep}
}

func (u DeleteUserUseCase) Invoke(id int) *lib.Response {
	res := u.rep.Delete(id)
	if res != nil {
		return &lib.Response{
			Code:    http.StatusOK,
			Message: "Delete failed",
		}
	}

	return &lib.Response{
		Code:    http.StatusOK,
		Message: "Delete successfully",
	}
}
