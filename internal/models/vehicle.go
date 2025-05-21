package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	Type      string  `gorm:"type:varchar(50);not null" json:"type"`
	Capacity  int     `gorm:"not null" json:"capacity"`
	Latitude  float64 `gorm:"not null" json:"latitude"`
	Longitude float64 `gorm:"not null" json:"longitude"`
	Speed     float64 `gorm:"not null" json:"speed"`
}
