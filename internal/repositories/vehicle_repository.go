package repositories

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	Create(vehicle models.Vehicle) (models.Vehicle, error)
}

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) VehicleRepository {
	return &vehicleRepository{db: db}
}

func (r *vehicleRepository) Create(vehicle models.Vehicle) (models.Vehicle, error) {
	if err := r.db.Create(&vehicle).Error; err != nil {
		return models.Vehicle{}, err
	}

	return vehicle, nil
}
