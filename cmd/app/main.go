package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"taskManager/internal/app"
	"taskManager/internal/infrastruct/config"
	"taskManager/pkg"
)

func main() {
	cfg := config.MustLoadAppConfig()
	db := pkg.MustConnectDatabase(cfg.Database.Url)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(fmt.Sprintf("Error closing database connection: %v", err))
		}
	}(db)

	server := fiber.New(fiber.Config{
		ReadTimeout:       cfg.Timeout,
		IdleTimeout:       cfg.IdleTimeout,
		WriteTimeout:      cfg.Timeout,
		EnablePrintRoutes: true,
	})

	api := server.Group("/api")

	app.NewApp(db, api).Build()

	err := server.Listen(cfg.HttpServer.Address)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
