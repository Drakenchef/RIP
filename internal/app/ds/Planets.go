package ds

import "gorm.io/gorm"

type Planet struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(50)" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Radius      float64 `gorm:"not null" json:"radius"`
	Distance    float64 `gorm:"not null" json:"distance"`
	Gravity     float64 `gorm:"not null" json:"gravity"`
	Image       string  `gorm:"type:varchar(255)" json:"image"`
	Type        string  `gorm:"not null" json:"type"`
	IsDelete    bool    `json:"is_delete"`
	//FlightRequests []FlightRequest `gorm:"many2many:planets_requests"`
}
