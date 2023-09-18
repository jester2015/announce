package app

import (
	v1 "absent.com/absentapi/routers/v1"
	"os"

	"absent.com/absentapi/config"
	"absent.com/absentapi/database"
	"absent.com/absentapi/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupAndRunApp() error {
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	//Setup new database
	database.CreateDatabase()

	//Setup Migrations
	err = database.SqlLiteClient.AutoMigrate(&models.Announcement{})
	if err != nil {
		return err
	}

	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	v1.SetupRoutes(app)

	// attach swagger
	//config.AddSwaggerRoutes(app)

	// get the port and start
	port := os.Getenv("PORT")
	err = app.Listen(":" + port)
	if err != nil {
		return err
	}

	return nil
}
