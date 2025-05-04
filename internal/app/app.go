package app

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"taskManager/internal/application"
	"taskManager/internal/infrastruct"
	"taskManager/internal/infrastruct/database"
	"taskManager/internal/presentation"
)

type App struct {
	db     *sql.DB
	router fiber.Router
}

func NewApp(
	db *sql.DB,
	router fiber.Router,
) *App {
	return &App{
		db:     db,
		router: router,
	}
}

func (app App) Build() {
	databaseInstance := database.NewDatabase(app.db)
	infrastructModule := infrastruct.BuildInfrastructModule(databaseInstance)
	applicationModule := application.NewApplicationModule(infrastructModule)
	presentationModule := presentation.BuildPresentation(app.router, applicationModule)

	presentationModule.User.BuildUserPresentation()
}
