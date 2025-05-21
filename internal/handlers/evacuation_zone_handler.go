package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/services"
)

type EvacuationZoneHandler interface {
	CreateZone(c *fiber.Ctx) error
}

type evacuationZoneHandler struct {
	evacuationZoneService services.EvacuationZoneService
}

func NewEvacuationZoneHandler(router fiber.Router, evacuationZoneService services.EvacuationZoneService) EvacuationZoneHandler {
	handler := &evacuationZoneHandler{evacuationZoneService: evacuationZoneService}

	router.Post("/", handler.CreateZone)

	return handler
}

func (h *evacuationZoneHandler) CreateZone(c *fiber.Ctx) error {
	var input models.EvacuationZoneInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "format invalid",
			"error":   err.Error(),
		})
	}

	evaZoneInput := models.EvacuationZone{
		ZoneID:         input.ZoneID,
		Latitude:       input.LocationCoordinates.Latitude,
		Longitude:      input.LocationCoordinates.Longitude,
		NumberOfPeople: input.NumberOfPeople,
		UrgencyLevel:   input.UrgencyLevel,
	}

	evaZone, err := h.evacuationZoneService.CreateEvacuationZone(evaZoneInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	evaZoneResponse := models.EvacuationZoneInput{
		ZoneID: evaZone.ZoneID,
		LocationCoordinates: models.LocationCoordinates{
			Latitude:  evaZone.Latitude,
			Longitude: evaZone.Longitude,
		},
		NumberOfPeople: evaZone.NumberOfPeople,
		UrgencyLevel:   evaZone.UrgencyLevel,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    evaZoneResponse,
	})
}
