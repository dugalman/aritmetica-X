package models

type User struct {
	UserID     uint            `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Num1       float64         `json:"num1"`
	Num2       float64         `json:"num2"`
	Op         int             `json:"op"`
	Username   string          `gorm:"type:varchar(100)" json:"username"`
	Password   string          `gorm:"type:varchar(100)" json:"password"`
	Operations []UserOperation `gorm:"foreignkey:UserID" json:"operations"`
}

type UserOperation struct {
	Id        uint    `gorm:"type:varchar(100)" json:"id"`
	UserID    uint    `gorm:"not null" json:"user_id"`
	Operation string  `gorm:"not null" json:"operation"`
	Result    float64 `gorm:"not null" json:"result"`
}
