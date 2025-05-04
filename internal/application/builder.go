package application

import (
	"taskManager/internal/application/user"
	"taskManager/internal/infrastruct"
)

type ApplicationModule struct {
	User *user.UserApplication
}

func NewApplicationModule(inf *infrastruct.InfrastructModule) *ApplicationModule {
	return &ApplicationModule{
		User: user.BuildUserApplication(inf.Repository.User),
	}
}
