package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/services"
)

type VehicleHandler interface {
	CreateVehicle(c *fiber.Ctx) error
}

type vehicleHandler struct {
	vehicleService services.VehicleService
}

func NewVehicleHandler(router fiber.Router, vehicleService services.VehicleService) VehicleHandler {
	handler := &vehicleHandler{vehicleService: vehicleService}

	router.Post("/", handler.CreateVehicle)

	return handler
}

func (h *vehicleHandler) CreateVehicle(c *fiber.Ctx) error {
	var input models.VehicleInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "format invalid",
			"error":   err.Error(),
		})
	}

	vehicleInput := models.Vehicle{
		VehicleID: input.VehicleID,
		Latitude:  input.LocationCoordinates.Latitude,
		Longitude: input.LocationCoordinates.Longitude,
		Capacity:  input.Capacity,
		Type:      input.Type,
		Speed:     input.Speed,
	}

	vehicle, err := h.vehicleService.CreateVehicle(vehicleInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	vehicleResponse := models.VehicleInput{
		VehicleID: vehicle.VehicleID,
		LocationCoordinates: models.LocationCoordinates{
			Latitude:  vehicle.Latitude,
			Longitude: vehicle.Longitude,
		},
		Capacity: vehicle.Capacity,
		Type:     vehicle.Type,
		Speed:    vehicle.Speed,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    vehicleResponse,
	})
}
