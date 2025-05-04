package infrastruct

import (
	"taskManager/internal/infrastruct/database"
	"taskManager/internal/infrastruct/repository"
)

type InfrastructModule struct {
	Repository *repository.Repository
}

func BuildInfrastructModule(db *database.Database) *InfrastructModule {
	return &InfrastructModule{
		Repository: repository.NewRepository(db),
	}
}
