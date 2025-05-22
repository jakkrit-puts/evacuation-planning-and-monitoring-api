package repositories

import (
	"errors"

	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/pkg/errs"
	"gorm.io/gorm"
)

type EvacuationZoneRepository interface {
	Create(evacuationZone models.EvacuationZone) (models.EvacuationZone, error)
	FindUrgentZones() ([]models.EvacuationZone, error)
	FindZoneByID(id string) (models.EvacuationZone, error)
	CheckExists(id string) (bool, error)
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

func (r *evacuationZoneRepository) FindUrgentZones() ([]models.EvacuationZone, error) {
	var zones []models.EvacuationZone

	err := r.db.Where("number_of_people > 0").Order("urgency_level DESC").Find(&zones).Error
	return zones, err
}

func (r *evacuationZoneRepository) FindZoneByID(id string) (models.EvacuationZone, error) {
	var zone models.EvacuationZone

	if err := r.db.First(&zone, "zone_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.EvacuationZone{}, errs.ErrDataNotFound
		}
		return models.EvacuationZone{}, err
	}

	return zone, nil
}

func (r *evacuationZoneRepository) CheckExists(id string) (bool, error) {
	var count int64
	err := r.db.Model(&models.EvacuationZone{}).
		Where("zone_id = ?", id).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
