package repositories

import (
	"errors"

	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/pkg/errs"
	"gorm.io/gorm"
)

type EvacuationStatusRepository interface {
	FindAll() ([]models.EvacuationStatus, error)
	FindByID(id string) (models.EvacuationStatus, error)
	Create(evacuationStatus models.EvacuationStatus) (models.EvacuationStatus, error)
	CheckExists(id string) bool
	Update(status models.EvacuationStatus) (models.EvacuationStatus, error)
}

type evacuationStatusRepository struct {
	db *gorm.DB
}

func NewEvacuationStatusRepository(db *gorm.DB) EvacuationStatusRepository {
	return &evacuationStatusRepository{db: db}
}

func (r *evacuationStatusRepository) Create(evacuationStatus models.EvacuationStatus) (models.EvacuationStatus, error) {
	if err := r.db.Create(&evacuationStatus).Error; err != nil {
		return models.EvacuationStatus{}, err
	}

	return evacuationStatus, nil
}

func (r *evacuationStatusRepository) FindAll() ([]models.EvacuationStatus, error) {
	var evacuationStatus []models.EvacuationStatus
	if err := r.db.Find(&evacuationStatus).Error; err != nil {
		return []models.EvacuationStatus{}, err
	}

	return evacuationStatus, nil
}

func (r *evacuationStatusRepository) FindByID(id string) (models.EvacuationStatus, error) {
	var evacuationStatus models.EvacuationStatus

	if err := r.db.First(&evacuationStatus, "zone_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.EvacuationStatus{}, errs.ErrDataNotFound
		}
		return models.EvacuationStatus{}, err
	}

	return evacuationStatus, nil
}

func (r *evacuationStatusRepository) CheckExists(id string) bool {
	var count int64
	err := r.db.Model(&models.EvacuationStatus{}).
		Where("zone_id = ?", id).
		Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func (r *evacuationStatusRepository) Update(evacuationStatus models.EvacuationStatus) (models.EvacuationStatus, error) {
	err := r.db.Save(evacuationStatus).Error
	return evacuationStatus, err
}
