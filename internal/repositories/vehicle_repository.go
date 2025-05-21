package repositories

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	Create(vehicle models.Vehicle) (models.Vehicle, error)
	FindAll() ([]models.Vehicle, error)
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

func (r *vehicleRepository) FindAll() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	if err := r.db.Find(&vehicles).Error; err != nil {
		return []models.Vehicle{}, err
	}

	return vehicles, nil
}
