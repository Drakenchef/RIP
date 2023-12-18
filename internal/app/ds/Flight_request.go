package ds

import (
	"time"
)

type FlightRequest struct {
	ID             uint             `json:"id" gorm:"primary_key"`
	DateCreate     time.Time        `json:"date_create"`
	DateFormation  time.Time        `json:"date_formation"`
	DateCompletion time.Time        `json:"date_completion;default:'0001-01-01'"`
	Status         string           `gorm:"type:varchar(255)" json:"status"`
	AMS            string           `gorm:"type:varchar(255)" json:"ams"`
	UserID         uint             `json:"user_id"`
	ModerID        uint             `json:"moder_id"`
	ModerLogin     string           `json:"moder_login"`
	UserLogin      string           `json:"user_login"`
	User           Users            `gorm:"foreignKey:UserID" json:"-"`
	PlanetsRequest []PlanetsRequest `json:"planets_request" gorm:"foreignkey:FRID"`
	Result         string           `gorm:"type:varchar(255)" json:"result;default:'не запущен'"`
}
