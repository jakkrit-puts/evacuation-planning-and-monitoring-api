package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/handlers"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/repositories"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/services"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/pkg/cache"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/pkg/config"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/pkg/db"
)

func main() {

	config := config.NewEnvConfig()
	db := db.Init(config, db.DBMigrator)
	redisClient := cache.InitRedis()
	// _ = redisClient

	app := fiber.New()

	server := app.Group("/api")

	evacuationZoneRepository := repositories.NewEvacuationZoneRepository(db)
	vehicleRepository := repositories.NewVehicleRepository(db)
	evacuationPlanRepository := repositories.NewEvacuationPlanRepository(db)
	evacuationStatusRepository := repositories.NewEvacuationStatusRepository(db)

	evacuationZoneService := services.NewEvacuationZoneService(evacuationZoneRepository)
	vehicleService := services.NewVehicleService(vehicleRepository)
	evacuationPlanService := services.NewEvacuationPlanService(evacuationPlanRepository)
	evacuationStatusService := services.NewEvacuationStatusService(evacuationStatusRepository)

	handlers.NewEvacuationZoneHandler(server.Group("/evacuation-zones"), evacuationZoneService)
	handlers.NewVehicleHandler(server.Group("/vehicles"), vehicleService)
	// handlers.NewEvacuationPlanHandler(server.Group("/evacuations"), evacuationPlanService, evacuationZoneService, vehicleService, evacuationStatusService)
	handlers.NewEvacuationPlanRedisHandler(server.Group("/evacuations"), evacuationPlanService, evacuationZoneService, vehicleService, evacuationStatusService, redisClient)

	app.Listen(fmt.Sprintf(":%s", config.AppPort))

}
