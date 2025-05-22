package repositories

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"gorm.io/gorm"
)

type EvacuationPlanRepository interface {
	Create(evacuationPlan models.EvacuationPlan) (models.EvacuationPlan, error)
	CheckExists(zoneID string, vehicleID string) (bool, error)
	DeleteAll() error
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

func (r *evacuationPlanRepository) CheckExists(zoneID string, vehicleID string) (bool, error) {
	var count int64
	err := r.db.Model(&models.EvacuationPlan{}).
		Where("zone_id = ? AND vehicle_id = ?", zoneID, vehicleID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *evacuationPlanRepository) DeleteAll() error {
	return r.db.Exec("DELETE FROM evacuation_plans").Error
}
