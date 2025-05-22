package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/services"
)

type EvacuationPlanHandler interface {
	GeneratesPlan(c *fiber.Ctx) error
}

type evacuationPlanHandler struct {
	evacuationPlanService services.EvacuationPlanService
	evacuationZoneService services.EvacuationZoneService
	vehicleService        services.VehicleService
}

func NewEvacuationPlanHandler(
	router fiber.Router,
	evacuationPlanService services.EvacuationPlanService,
	evacuationZoneService services.EvacuationZoneService,
	vehicleService services.VehicleService,
) EvacuationPlanHandler {
	handler := &evacuationPlanHandler{
		evacuationPlanService: evacuationPlanService,
		evacuationZoneService: evacuationZoneService,
		vehicleService:        vehicleService,
	}

	router.Post("/plan", handler.GeneratesPlan)

	return handler
}

func (h *evacuationPlanHandler) GeneratesPlan(c *fiber.Ctx) error {

	var zones []models.EvacuationZone
	var vehicles []models.Vehicle

	zones, err := h.evacuationZoneService.GetUrgentZones()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	vehicles, errNew := h.vehicleService.GetVehicles()
	if errNew != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errNew.Error(),
		})
	}

	plans := h.evacuationPlanService.GenerateEvacuationPlan(zones, vehicles)

	var savedPlans []models.EvacuationPlan
	var newPlan models.EvacuationPlan
	for _, p := range plans {

		newPlan = models.EvacuationPlan{
			ZoneID:         p.ZoneID,
			VehicleID:      p.VehicleID,
			NumberOfPeople: p.NumberOfPeople,
			ETA:            p.ETA,
		}

		saved, err := h.evacuationPlanService.CreateEvacuationPlan(newPlan)
		if err != nil {
			fmt.Printf("failed to save plan: %+v, error: %v\n", p, err)
			continue
		}
		savedPlans = append(savedPlans, saved)
	}

	return c.JSON(savedPlans)
}
