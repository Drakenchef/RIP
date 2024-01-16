package ds

type Planet struct {
	//gorm.Model
	ID          uint    `json:"id" gorm:"primary_key"`
	Name        string  `gorm:"type:varchar(50)" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Radius      float64 `gorm:"not null" json:"radius"`
	Distance    float64 `gorm:"not null" json:"distance"`
	Gravity     float64 `gorm:"not null" json:"gravity"`
	Image       string  `json:"image" gorm:"type:varchar(1000);default:'http://172.18.0.5:9000/amsflights/notfound.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=LPS0EGHZC58LP7DX4FMN%2F20231023%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20231023T151625Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiJMUFMwRUdIWkM1OExQN0RYNEZNTiIsImV4cCI6MTY5ODA3NzY3NiwicGFyZW50IjoibWluaW8ifQ.qGA2Cz3RLMKAKlVOOhjq1lke8oTUN_FhOu9cKJAkpflhSSR5BtL5cqwHx5eJVD61y7L_EesdxJyutyWlXHvvRQ&X-Amz-SignedHeaders=host&versionId=8de62709-768d-4266-9c5b-44e1d22c5442&X-Amz-Signature=7276264afea1dee29e45b98156dc1bce952bfedeb7f6132cf3d7dd2a4b1ad701'"`
	Type        string  `gorm:"not null" json:"type"`
	IsDelete    bool    `json:"is_delete"`
}
