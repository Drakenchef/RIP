package ds

type Users struct {
	//gorm.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	Login    string `gorm:"type:varchar(255);unique" json:"login"`
	Password string `gorm:"type:varchar(255)" json:"-"`
	IsAdmin  bool   `json:"is_admin"`
}
