package ds

import "github.com/drakenchef/RIP/internal/app/role"

type Users struct {
	//gorm.Model
	ID       uint      `json:"id" gorm:"primary_key"`
	UserName string    `json:"user_name" gorm:"type:varchar(50)"`
	Login    string    `gorm:"type:varchar(255);unique" json:"login"`
	Password string    `gorm:"type:varchar(255)" json:"-"`
	Role     role.Role `json:"role" sql:"type:string"`
}
