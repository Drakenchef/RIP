package ds

import (
	"gorm.io/gorm"
	"time"
)

type FlightRequest struct {
	gorm.Model
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
	Status    string    `gorm:"type:varchar(255)" json:"status"`
	AMS       string    `gorm:"type:varchar(255)" json:"ams"`
	UserID    uint      `json:"-"`
	User      Users     `gorm:"foreignKey:UserID" json:"-"`
}
