package models

import (
	"github.com/google/uuid"
)

type EvacuationZone struct {
	ZoneID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"zone_id"`
	Latitude       float64   `gorm:"not null" json:"latitude"`
	Longitude      float64   `gorm:"not null" json:"longitude"`
	NumberOfPeople int       `gorm:"not null" json:"number_of_people"`
	UrgencyLevel   int       `gorm:"not null" json:"urgency_level"`
}
