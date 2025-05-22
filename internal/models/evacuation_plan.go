package models

type EvacuationPlan struct {
	ZoneID         string `gorm:"primaryKey;not null" json:"ZoneID"`
	VehicleID      string `gorm:"not null" json:"VehicleID"`
	NumberOfPeople int    `gorm:"not null" json:"NumberOfPeople"`
	ETA            string `gorm:"not null" json:"ETA"`
}

type EvacuationPlanResponse struct {
	ZoneID         string `json:"ZoneID"`
	VehicleID      string `json:"VehicleID"`
	NumberOfPeople int    `json:"NumberOfPeople"`
	ETA            string `json:"ETA"`
}

type EvacuationUpdateInput struct {
	ZoneID      string `json:"ZoneID"`
	VehicleID   string `json:"VehicleID"`
	PeopleMoved int    `json:"PeopleMoved"`
}
