package repositories

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"gorm.io/gorm"
)

type EvacuationPlanRepository interface {
	Create(evacuationPlan models.EvacuationPlan) (models.EvacuationPlan, error)
}

type evacuationPlanRepository struct {
	db *gorm.DB
}

func NewEvacuationPlanRepository(db *gorm.DB) EvacuationPlanRepository {
	return &evacuationPlanRepository{db: db}
}

func (r *evacuationPlanRepository) Create(evacuationPlan models.EvacuationPlan) (models.EvacuationPlan, error) {
	if err := r.db.Create(&evacuationPlan).Error; err != nil {
		return models.EvacuationPlan{}, err
	}

	return evacuationPlan, nil
}
