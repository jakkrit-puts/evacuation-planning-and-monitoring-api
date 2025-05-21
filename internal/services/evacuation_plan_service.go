package services

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/repositories"
)

type EvacuationPlanService interface {
	CreateEvacuationPlan(evacuationPlan models.EvacuationPlan) (models.EvacuationPlan, error)
}

type evacuationPlanService struct {
	evacuationPlanRepository repositories.EvacuationPlanRepository
}

func NewEvacuationPlanService(evacuationPlanRepository repositories.EvacuationPlanRepository) EvacuationPlanService {
	return &evacuationPlanService{evacuationPlanRepository}
}

func (s *evacuationPlanService) CreateEvacuationPlan(evacuationPlan models.EvacuationPlan) (models.EvacuationPlan, error) {
	return s.evacuationPlanRepository.Create(evacuationPlan)
}
