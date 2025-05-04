package mapper

import (
	"taskManager/internal/domain/user"
	"taskManager/internal/infrastruct/database/entity"
)

func MapUserFromEntityToDomain(ent *entity.UserEntity) *user.UserModel {
	return &user.UserModel{
		Id:       ent.Id,
		Name:     ent.Name,
		Email:    ent.Email,
		Password: ent.Password,
	}
}
