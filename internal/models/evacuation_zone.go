package models

type EvacuationZone struct {
	ZoneID         string  `gorm:"primaryKey;unique"  `
	Latitude       float64 `gorm:"not null" json:"latitude"`
	Longitude      float64 `gorm:"not null" json:"longitude"`
	NumberOfPeople int     `gorm:"not null" json:"number_of_people"`
	UrgencyLevel   int     `gorm:"not null" json:"urgency_level"`
}

type LocationCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type EvacuationZoneInput struct {
	ZoneID              string              `json:"ZoneID"`
	LocationCoordinates LocationCoordinates `json:"LocationCoordinates"`
	NumberOfPeople      int                 `json:"NumberOfPeople"`
	UrgencyLevel        int                 `json:"UrgencyLevel"`
}
