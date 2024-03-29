package models

import "examen_server/db"

type User struct {
	Id       uint
	Username string `gorm:"type:varchar(100)" json:"username"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	// Operation string `gorm:"type:varchar(100)" json:"operation"`
}

func Migrations() {
	db.Database.AutoMigrate(&User{})
}
