package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/services"
)

type EvacuationPlanHandler interface {
	GeneratesPlan(c *fiber.Ctx) error
}

type evacuationPlanHandler struct {
	evacuationPlanService services.EvacuationPlanService
}

func NewEvacuationPlanHandler(router fiber.Router, evacuationPlanService services.EvacuationPlanService) EvacuationPlanHandler {
	handler := &evacuationPlanHandler{evacuationPlanService: evacuationPlanService}

	router.Post("/plan", handler.GeneratesPlan)

	return handler
}

func (h *evacuationPlanHandler) GeneratesPlan(c *fiber.Ctx) error {
	fmt.Println("plan gen")
	fmt.Println("get zones -> urgency levels first, vehicle")
	fmt.Println("Distance Calculation by Haversine")
	fmt.Println("Travel Time Estimation")
	return nil
}
