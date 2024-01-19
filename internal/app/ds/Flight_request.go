package ds

import "time"

type FlightRequest struct {
	DateCreate     time.Time `json:"date_create"`
	DateFormation  time.Time `json:"date_formation"`
	DateCompletion time.Time `json:"date_completion"`
	Status         string    `gorm:"type:varchar(255)" json:"status"`
	AMS            string    `gorm:"type:varchar(255)" json:"ams"`
	UserID         uint      `json:"-"`
	ModerID        uint      `json:"-"`
	User           Users     `gorm:"foreignKey:UserID" json:"-"`
	Planets        []Planet  `gorm:"many2many:planets_requests;"`
	DateFlight     time.Time `json:"date_flight"`
}
