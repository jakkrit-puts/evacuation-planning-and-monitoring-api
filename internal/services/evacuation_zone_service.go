package services

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/repositories"
)

type EvacuationZoneService interface {
	CreateEvacuationZone(brand models.EvacuationZone) (models.EvacuationZone, error)
}

type evacuationZoneService struct {
	evacuationZoneRepository repositories.EvacuationZoneRepository
}

func NewBrandService(evacuationZoneRepository repositories.EvacuationZoneRepository) EvacuationZoneService {
	return &evacuationZoneService{evacuationZoneRepository}
}

func (s *evacuationZoneService) CreateEvacuationZone(evacuationZone models.EvacuationZone) (models.EvacuationZone, error) {
	return s.evacuationZoneRepository.Create(evacuationZone)
}
