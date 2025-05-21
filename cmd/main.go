package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/pkg/config"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/pkg/db"
)

func main() {

	config := config.NewEnvConfig()
	db := db.Init(config, db.DBMigrator)
	_ = db

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(fmt.Sprintf(":%s", config.AppPort))

}