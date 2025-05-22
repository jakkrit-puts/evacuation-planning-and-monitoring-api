package services

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/repositories"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/pkg/utils"
)

type EvacuationPlanService interface {
	CreateEvacuationPlan(evacuationPlan models.EvacuationPlan) (models.EvacuationPlan, error)
	GenerateEvacuationPlan(zones []models.EvacuationZone, vehicles []models.Vehicle) []models.EvacuationPlanResponse
	ClearPlan() error
}

type evacuationPlanService struct {
	evacuationPlanRepository repositories.EvacuationPlanRepository
}

func NewEvacuationPlanService(evacuationPlanRepository repositories.EvacuationPlanRepository) EvacuationPlanService {
	return &evacuationPlanService{evacuationPlanRepository}
}

func (s *evacuationPlanService) CreateEvacuationPlan(evacuationPlan models.EvacuationPlan) (models.EvacuationPlan, error) {
	ISexists, err := s.evacuationPlanRepository.CheckExists(evacuationPlan.ZoneID, evacuationPlan.VehicleID)
	if err != nil {
		return models.EvacuationPlan{}, err
	}

	if ISexists {
		return evacuationPlan, nil
	}

	return s.evacuationPlanRepository.Create(evacuationPlan)
}

func (s *evacuationPlanService) GenerateEvacuationPlan(zones []models.EvacuationZone, vehicles []models.Vehicle) []models.EvacuationPlanResponse {
	var plans []models.EvacuationPlanResponse

	sort.Slice(zones, func(i, j int) bool {
		return zones[i].UrgencyLevel > zones[j].UrgencyLevel
	})

	for _, z := range zones {
		remaining := z.NumberOfPeople

		sort.Slice(vehicles, func(i, j int) bool {
			di := utils.Haversine(z.Latitude, z.Longitude, vehicles[i].Latitude, vehicles[i].Longitude)
			dj := utils.Haversine(z.Latitude, z.Longitude, vehicles[j].Latitude, vehicles[j].Longitude)
			return di < dj
		})

		for _, v := range vehicles {
			if remaining <= 0 {
				break
			}

			distance := utils.Haversine(z.Latitude, z.Longitude, v.Latitude, v.Longitude)
			timeHours := distance / float64(v.Speed)
			eta := int(timeHours * 60)

			toEvacuate := v.Capacity
			if remaining < v.Capacity {
				toEvacuate = remaining
			}

			plans = append(plans, models.EvacuationPlanResponse{
				ZoneID:         z.ZoneID,
				VehicleID:      v.VehicleID,
				ETA:            fmt.Sprintf("%v minutes", strconv.Itoa(eta)),
				NumberOfPeople: toEvacuate,
			})

			remaining -= toEvacuate
		}
	}

	return plans
}

func (s *evacuationPlanService) ClearPlan() error {
	return s.evacuationPlanRepository.DeleteAll()
}
