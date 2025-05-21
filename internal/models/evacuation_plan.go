package models

import (
	"gorm.io/gorm"
)

type EvacuationPlan struct {
	gorm.Model
	ZoneID         string  `gorm:"not null" json:"zone_id"`
	VehicleID      string  `gorm:"not null" json:"vehicle_id"`
	NumberOfPeople int     `gorm:"not null" json:"number_of_people"`
	ETA            float64 `gorm:"not null" json:"eta"`
}
