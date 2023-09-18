package ds

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Login    string `gorm:"type:varchar(255);unique" json:"login"`
	Password string `gorm:"type:varchar(255)" json:"-"`
}
