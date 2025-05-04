package mapper

import (
	"taskManager/internal/application/user/dto/request"
	"taskManager/internal/application/user/dto/response"
	"taskManager/internal/domain/user"
)

func MapUserFromRequestToModel(dto *request.RequestCreateUser) *user.UserModel {
	return &user.UserModel{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func MapUserFromModelToResponse(model *user.UserModel) *response.ResponseUser {
	return &response.ResponseUser{
		Id:    model.Id,
		Name:  model.Name,
		Email: model.Email,
	}
}
