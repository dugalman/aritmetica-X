package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	UserID    uint    `gorm:"primaryKey;foreignKey:UserID"`
	Operation string  `gorm:"not null" json:"operation"`
	Result    float64 `gorm:"not null" json:"result"`
}

var dsn = "root:root@tcp(localhost:3306)/aritmetica?charset=utf8mb4&parseTime=True&loc=Local"
var Database *gorm.DB

func init() {
	var err error
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error en la conexión a la base de datos:", err)
		panic(err)
	}

	fmt.Println("Conexión a la base de datos exitosa")

	// Realizar migraciones aquí
	// if err := Database.AutoMigrate(&User{}, &UserOperation{}); err != nil {
	// 	fmt.Println("Error al realizar migraciones:", err)
	// 	panic(err)
	// }

	// fmt.Println("Migraciones completadas exitosamente")
}
