package services

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/repositories"
)

type EvacuationStatusService interface {
	GetEvacuationStatusList() ([]models.EvacuationStatus, error)
	FindZoneByID(id string) (models.EvacuationStatus, error)
	CreateEvacuationStatus(evacuationStatus models.EvacuationStatus) (models.EvacuationStatus, error)
	Exists(id string) bool
	Update(status models.EvacuationStatus) (models.EvacuationStatus, error)
}

type evacuationStatusService struct {
	evacuationStatusRepository repositories.EvacuationStatusRepository
}

func NewEvacuationStatusService(evacuationStatusRepository repositories.EvacuationStatusRepository) EvacuationStatusService {
	return &evacuationStatusService{evacuationStatusRepository}
}

func (s *evacuationStatusService) GetEvacuationStatusList() ([]models.EvacuationStatus, error) {
	return s.evacuationStatusRepository.FindAll()
}

func (s *evacuationStatusService) FindZoneByID(id string) (models.EvacuationStatus, error) {
	return s.evacuationStatusRepository.FindByID(id)
}

func (s *evacuationStatusService) CreateEvacuationStatus(evacuationStatus models.EvacuationStatus) (models.EvacuationStatus, error) {
	return s.evacuationStatusRepository.Create(evacuationStatus)
}

func (s *evacuationStatusService) Exists(id string) bool {
	return s.evacuationStatusRepository.CheckExists(id)
}

func (s *evacuationStatusService) Update(status models.EvacuationStatus) (models.EvacuationStatus, error) {
	return s.evacuationStatusRepository.Update(status)
}
