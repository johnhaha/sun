package main

import (
	"sun/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	api := app.Group("api")
	api.Get("/info", func(c *fiber.Ctx) error {
		apps, err := handlers.GetAppInfo()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "get app info failed",
				"success": false,
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "get app info ok",
			"success": true,
			"data":    apps,
		})
	})

	app.Get("", func(c *fiber.Ctx) error {
		return c.SendFile("assets/web/index.html")
	})

	app.Static("", "assets/web")

	app.Listen(":8008")
}
