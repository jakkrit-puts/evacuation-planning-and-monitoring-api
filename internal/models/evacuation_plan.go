package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EvacuationPlan struct {
	gorm.Model
	ZoneID         uuid.UUID `gorm:"type:uuid;not null" json:"zone_id"`
	VehicleID      uuid.UUID `gorm:"type:uuid;not null" json:"vehicle_id"`
	NumberOfPeople int       `gorm:"not null" json:"number_of_people"`
	ETA            float64   `gorm:"not null" json:"eta"`

	Zone    EvacuationZone `gorm:"foreignKey:ZoneID" json:"zone"`
	Vehicle Vehicle        `gorm:"foreignKey:VehicleID" json:"vehicle"`
}
