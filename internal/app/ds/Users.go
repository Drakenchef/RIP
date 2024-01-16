package ds

type Users struct {
	Login    string `gorm:"type:varchar(255);unique" json:"login"`
	Password string `gorm:"type:varchar(255)" json:"-"`
	IsAdmin  bool   `json:"is_admin"`
}
