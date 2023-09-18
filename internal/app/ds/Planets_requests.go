package ds

import "gorm.io/gorm"

type PlanetsRequest struct {
	gorm.Model
	FRID          uint          `json:"-"`
	PlanetID      uint          `json:"-"`
	FlightRequest FlightRequest `gorm:"foreignKey:FRID" json:"-"`
	Planet        Planet        `gorm:"foreignKey:PlanetID" json:"-"`
}
