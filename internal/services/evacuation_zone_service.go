package services

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/repositories"
)

type EvacuationZoneService interface {
	CreateEvacuationZone(evacuationZone models.EvacuationZone) (models.EvacuationZone, error)
	GetUrgentZones() ([]models.EvacuationZone, error)
	FindZoneByID(id string) (models.EvacuationZone, error)
}

type evacuationZoneService struct {
	evacuationZoneRepository repositories.EvacuationZoneRepository
}

func NewEvacuationZoneService(evacuationZoneRepository repositories.EvacuationZoneRepository) EvacuationZoneService {
	return &evacuationZoneService{evacuationZoneRepository}
}

func (s *evacuationZoneService) CreateEvacuationZone(evacuationZone models.EvacuationZone) (models.EvacuationZone, error) {
	return s.evacuationZoneRepository.Create(evacuationZone)
}

func (s *evacuationZoneService) GetUrgentZones() ([]models.EvacuationZone, error) {
	return s.evacuationZoneRepository.FindUrgentZones()
}

func (s *evacuationZoneService) FindZoneByID(id string) (models.EvacuationZone, error) {
	return s.evacuationZoneRepository.FindZoneByID(id)
}
