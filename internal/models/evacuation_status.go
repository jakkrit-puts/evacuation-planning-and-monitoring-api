package models

import (
	"time"
)

type EvacuationStatus struct {
	ZoneID          string    `gorm:"primaryKey" json:"zone_id"`
	TotalEvacuated  int       `gorm:"not null;default:0" json:"total_evacuated"`
	RemainingPeople int       `gorm:"not null" json:"remaining_people"`
	LastVehicleUsed *string   `gorm:"" json:"last_vehicle_used,omitempty"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
