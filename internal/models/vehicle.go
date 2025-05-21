package models

type Vehicle struct {
	VehicleID string  `gorm:"primaryKey;unique" json:"vehicle_id"`
	Type      string  `gorm:"type:varchar(50);not null" json:"type"`
	Capacity  int     `gorm:"not null" json:"capacity"`
	Latitude  float64 `gorm:"not null" json:"latitude"`
	Longitude float64 `gorm:"not null" json:"longitude"`
	Speed     int `gorm:"not null" json:"speed"`
}

type VehicleInput struct {
	VehicleID           string              `json:"VehicleID"`
	Capacity            int                 `json:"Capacity"`
	Type                string              `json:"Type"`
	LocationCoordinates LocationCoordinates `json:"LocationCoordinates"`
	Speed               int                 `json:"Speed"`
}
