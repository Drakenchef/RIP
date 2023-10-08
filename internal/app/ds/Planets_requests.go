package ds

import "gorm.io/gorm"

type PlanetsRequest struct {
	gorm.Model
	ID           uint `gorm:"auto_increment:false""`
	FRID         uint `json:"-" gorm:"primaryKey;auto_increment:false"`
	PlanetID     uint `json:"-" gorm:"primaryKey;auto_increment:false"`
	FlightNumber uint `json:"flight_number"`
}

//lab2comm
//FlightRequest FlightRequest `gorm:"foreignKey:FRID" json:"-"`
//Planet        Planet        `gorm:"foreignKey:PlanetID" json:"-"`
