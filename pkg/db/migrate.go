package db

import (
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/internal/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.EvacuationZone{},
		&models.Vehicle{},
		&models.EvacuationPlan{},
		&models.EvacuationStatus{},
	)
}
