package ds

import "gorm.io/gorm"

type PlanetsRequest struct {
	gorm.Model
	FRID          uint          `json:"-" gorm:"primary_key;auto_increment:false"`
	PlanetID      uint          `json:"-" gorm:"primary_key;auto_increment:false"`
	FlightRequest FlightRequest `gorm:"foreignKey:FRID" json:"-"`
	Planet        Planet        `gorm:"foreignKey:PlanetID" json:"-"`
	FlightNumber  uint          `json:"flight_number"`
}

//forcomm
