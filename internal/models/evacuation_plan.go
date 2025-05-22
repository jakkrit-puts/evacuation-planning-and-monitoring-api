package models

type EvacuationPlan struct {
	ZoneID         string `gorm:"not null" json:"ZoneID"`
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
