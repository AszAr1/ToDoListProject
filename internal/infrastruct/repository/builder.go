package repository

import (
	"taskManager/internal/domain/user"
	"taskManager/internal/infrastruct/database"
)

type Repository struct {
	User user.IUserRepository
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
