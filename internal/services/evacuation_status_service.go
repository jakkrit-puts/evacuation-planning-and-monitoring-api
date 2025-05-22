package services

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/repositories"
)

type EvacuationStatusService interface {
	GetEvacuationStatusList() ([]models.EvacuationStatus, error)
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
