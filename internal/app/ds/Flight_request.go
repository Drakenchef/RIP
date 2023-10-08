package ds

import (
	"gorm.io/gorm"
	"time"
)

type FlightRequest struct {
	gorm.Model
	DateCreate     time.Time `json:"date_create"`
	DateFormation  time.Time `json:"date_formation"`
	DateCompletion time.Time `json:"date_completion"`
	Status         string    `gorm:"type:varchar(255)" json:"status"`
	AMS            string    `gorm:"type:varchar(255)" json:"ams"`
	UserID         uint      `json:"-"`
	ModerID        uint      `json:"-"`
	User           Users     `gorm:"foreignKey:UserID" json:"-"`
	Planets        []Planet  `gorm:"many2many:planets_requests;"`
}
