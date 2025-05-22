package repositories

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"gorm.io/gorm"
)

type EvacuationStatusRepository interface {
	FindAll() ([]models.EvacuationStatus, error)
}

type evacuationStatusRepository struct {
	db *gorm.DB
}

func NewEvacuationStatusRepository(db *gorm.DB) EvacuationStatusRepository {
	return &evacuationStatusRepository{db: db}
}

func (r *evacuationStatusRepository) FindAll() ([]models.EvacuationStatus, error) {
	var evacuationStatus []models.EvacuationStatus
	if err := r.db.Find(&evacuationStatus).Error; err != nil {
		return []models.EvacuationStatus{}, err
	}

	return evacuationStatus, nil
}
