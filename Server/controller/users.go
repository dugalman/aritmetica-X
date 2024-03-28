package controller

import (
	"examen_server/db"
	"examen_server/models"
	"examen_server/utils"
	"fmt"
	"net"
)

func Register(conn net.Conn, username, password string) error {

	// encoder := gob.NewEncoder(conn)
	user := models.User{Username: username, Password: password}

	var existingUser models.User
	if err := db.Database.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		// Ya existe un usuario con el mismo correo electrónico
		fmt.Println("El Usuario ya está en uso por otro usuario.")
		return err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {

		fmt.Println("Error al hashear la contraseña")
		return err
	}

	user.Password = hashedPassword

	if err := db.Database.Save(&user).Error; err != nil {

		fmt.Println("Error al crear el usuario en la base de datos:", err)
		fmt.Println("Error al crear el usuario en la base de datos:", err)
		return err
	}

	return nil

}

// func Login(conn net.Conn, password, username string) {
// 	user := models.User{Username: username, Password: password}
// }
