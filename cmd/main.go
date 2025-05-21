package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/handlers"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/repositories"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/services"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/pkg/config"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/pkg/db"
)

func main() {

	config := config.NewEnvConfig()
	db := db.Init(config, db.DBMigrator)

	app := fiber.New()

	server := app.Group("/api")

	evacuationZoneRepository := repositories.NewEvacuationZoneRepository(db)
	vehicleRepository := repositories.NewVehicleRepository(db)

	evacuationZoneService := services.NewEvacuationZoneService(evacuationZoneRepository)
	vehicleService := services.NewVehicleService(vehicleRepository)

	handlers.NewEvacuationZoneHandler(server.Group("/evacuation-zones"), evacuationZoneService)
	handlers.NewVehicleHandler(server.Group("/vehicles"), vehicleService)

	app.Listen(fmt.Sprintf(":%s", config.AppPort))

}
