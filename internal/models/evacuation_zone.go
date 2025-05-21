package models

import "gorm.io/gorm"

type EvacuationZone struct {
	gorm.Model
	ZoneID         string  `gorm:"type:varchar(10);unique;not null" json:"zone_id"`
	Latitude       float64 `gorm:"not null" json:"latitude"`
	Longitude      float64 `gorm:"not null" json:"longitude"`
	NumberOfPeople int     `gorm:"not null" json:"number_of_people"`
	UrgencyLevel   int     `gorm:"not null" json:"urgency_level"`
}
