package v1

import (
	"absent.com/absentapi/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	// v1 Routes
	v1 := api.Group("/v1")

	announcement := v1.Group("/announcement")

	announcement.Get("/", handlers.HandleAllAnnouncements)
	announcement.Get("/:id", handlers.HandleGetOneAnnouncement)
	announcement.Post("/", handlers.HandleCreateAnnouncements)
	announcement.Put("/:id", handlers.HandleUpdateAnnouncements)
	announcement.Delete("/:id", handlers.HandleDeleteAnnouncements)

}
