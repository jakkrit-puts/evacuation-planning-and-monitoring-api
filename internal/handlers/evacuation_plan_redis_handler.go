package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/services"
	"github.com/redis/go-redis/v9"
)

type EvacuationPlanRedisHandler interface {
	GeneratesPlan(c *fiber.Ctx) error
	GetEvacuationStatus(c *fiber.Ctx) error
	UpdateEvacuationStatus(c *fiber.Ctx) error
	ClearEvacuations(c *fiber.Ctx) error
}

type evacuationPlanRedisHandler struct {
	evacuationPlanService   services.EvacuationPlanService
	evacuationZoneService   services.EvacuationZoneService
	vehicleService          services.VehicleService
	evacuationStatusService services.EvacuationStatusService
	redisClient             *redis.Client
}

func NewEvacuationPlanRedisHandler(
	router fiber.Router,
	evacuationPlanService services.EvacuationPlanService,
	evacuationZoneService services.EvacuationZoneService,
	vehicleService services.VehicleService,
	evacuationStatusService services.EvacuationStatusService,
	redisClient *redis.Client,
) EvacuationPlanRedisHandler {
	handler := &evacuationPlanRedisHandler{
		evacuationPlanService:   evacuationPlanService,
		evacuationZoneService:   evacuationZoneService,
		vehicleService:          vehicleService,
		evacuationStatusService: evacuationStatusService,
		redisClient:             redisClient,
	}

	router.Post("/plan", handler.GeneratesPlan)
	router.Get("/status", handler.GetEvacuationStatus)
	router.Put("/update", handler.UpdateEvacuationStatus)
	router.Delete("/clear", handler.ClearEvacuations)

	return handler
}

func (h *evacuationPlanRedisHandler) GeneratesPlan(c *fiber.Ctx) error {

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
			fmt.Printf("failed save plan: %v, error: %v\n", p, err)
			continue
		}

		exists := h.evacuationStatusService.Exists(p.ZoneID)
		if !exists {
			zone, _ := h.evacuationZoneService.FindZoneByID(p.ZoneID)
			status := models.EvacuationStatus{
				ZoneID:          p.ZoneID,
				TotalEvacuated:  0,
				RemainingPeople: zone.NumberOfPeople,
				LastVehicleUsed: nil,
			}
			_, _ = h.evacuationStatusService.CreateEvacuationStatus(status)
		}

		savedPlans = append(savedPlans, saved)
	}

	return c.JSON(savedPlans)
}

func (h *evacuationPlanRedisHandler) GetEvacuationStatus(c *fiber.Ctx) error {

	key := "handler::GetEvacuationStatus"

	if reponseJson, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("...get data from redis")
		c.Set("content-type", "application/json")
		return c.SendString(reponseJson)
	}

	// service
	statuses, err := h.evacuationStatusService.GetEvacuationStatusList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// SET
	if data, err := json.Marshal(statuses); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	fmt.Println("...get data from DB")
	return c.JSON(statuses)

}

func (h *evacuationPlanRedisHandler) UpdateEvacuationStatus(c *fiber.Ctx) error {
	var input models.EvacuationUpdateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if input.PeopleMoved <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "PeopleMoved gt than 0",
		})
	}

	status, err := h.evacuationStatusService.FindZoneByID(input.ZoneID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if input.PeopleMoved > status.RemainingPeople {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "PeopleMoved exceeds",
		})
	}

	status.TotalEvacuated += input.PeopleMoved
	status.RemainingPeople -= input.PeopleMoved
	status.LastVehicleUsed = &input.VehicleID

	result, errUpdate := h.evacuationStatusService.Update(status)
	if errUpdate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errUpdate.Error(),
		})
	}

	updatedStatuses, _ := h.evacuationStatusService.GetEvacuationStatusList()
	if data, err := json.Marshal(updatedStatuses); err == nil {
		h.redisClient.Set(context.Background(), "handler::GetEvacuationStatus", string(data), time.Second*30)
	}

	return c.JSON(fiber.Map{
		"message": "updated successfully",
		"data":    result,
	})
}

func (h *evacuationPlanRedisHandler) ClearEvacuations(c *fiber.Ctx) error {
	if err := h.evacuationPlanService.ClearPlan(); err != nil {
		return err
	}

	zones, err := h.evacuationZoneService.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	for _, zone := range zones {
		status := models.EvacuationStatus{
			ZoneID:          zone.ZoneID,
			TotalEvacuated:  0,
			RemainingPeople: zone.NumberOfPeople,
			LastVehicleUsed: nil,
		}

		h.evacuationStatusService.Update(status)
	}

	return c.JSON(fiber.Map{
		"message": "clear successfully",
	})
}
