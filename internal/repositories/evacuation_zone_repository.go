package repositories

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"gorm.io/gorm"
)

type EvacuationZoneRepository interface {
	Create(evacuationZone models.EvacuationZone) (models.EvacuationZone, error)
}

type evacuationZoneRepository struct {
	db *gorm.DB
}

func NewEvacuationZoneRepository(db *gorm.DB) EvacuationZoneRepository {
	return &evacuationZoneRepository{db: db}
}

func (r *evacuationZoneRepository) Create(evacuationZone models.EvacuationZone) (models.EvacuationZone, error) {
	if err := r.db.Create(&evacuationZone).Error; err != nil {
		return models.EvacuationZone{}, err
	}

	return evacuationZone, nil
}