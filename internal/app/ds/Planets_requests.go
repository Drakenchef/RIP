package ds

import "gorm.io/gorm"

type PlanetsRequest struct {
	gorm.Model
	ID           uint `gorm:"auto_increment:false""`
	FRID         uint `json:"fr_id" gorm:"primaryKey;auto_increment:false"`
	PlanetID     uint `json:"planet_id" gorm:"primaryKey;auto_increment:false"`
	FlightNumber uint `json:"flight_number"`
}
