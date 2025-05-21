package models

import (
	"time"

	"github.com/google/uuid"
)

type EvacuationStatus struct {
	ZoneID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"zone_id"`
	TotalEvacuated  int            `gorm:"not null;default:0" json:"total_evacuated"`
	RemainingPeople int            `gorm:"not null" json:"remaining_people"`
	LastVehicleUsed *uuid.UUID     `gorm:"type:uuid" json:"last_vehicle_used,omitempty"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}
