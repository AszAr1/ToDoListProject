package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"taskManager/internal/infrastruct/config"
	"taskManager/pkg"
)

func main() {
	cfg := config.MustLoadAppConfig()
	db := pkg.MustNewDbInstance(cfg.Database.Url)
	defer db.Instance.Close()

	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
		WriteTimeout: cfg.Timeout,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	err := app.Listen(cfg.HttpServer.Address)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
