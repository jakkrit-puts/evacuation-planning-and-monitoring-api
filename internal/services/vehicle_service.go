package services

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/repositories"
)

type VehicleService interface {
	CreateVehicle(brand models.Vehicle) (models.Vehicle, error)
}

type vehicleService struct {
	vehicleRepository repositories.VehicleRepository
}

func NewVehicleService(vehicleRepository repositories.VehicleRepository) VehicleService {
	return &vehicleService{vehicleRepository}
}

func (s *vehicleService) CreateVehicle(vehicle models.Vehicle) (models.Vehicle, error) {
	return s.vehicleRepository.Create(vehicle)
}
    