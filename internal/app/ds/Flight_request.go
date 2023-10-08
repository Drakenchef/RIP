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
	DateApprove    time.Time `json:"date_approve"`
	DateRefuse     time.Time `json:"date_refuse"`
	Status         string    `gorm:"type:varchar(255)" json:"status"`
	AMS            string    `gorm:"type:varchar(255)" json:"ams"`
	UserID         uint      `json:"user_id"`
	ModerID        uint      `json:"moder_id"`
	User           Users     `gorm:"foreignKey:UserID" json:"-"`
	//Planets        []Planet  `gorm:"many2many:planets_requests;"`
}
