package controller

import (
	"examen_server/db"
	"examen_server/models"
	"examen_server/utils"
	"fmt"
	"net"

	"golang.org/x/crypto/bcrypt"
)

func Register(conn net.Conn, username, password string) error {
	user := models.User{Username: username}

	// Buscar el usuario en la base de datos por nombre de usuario
	var existingUser models.User
	if err := db.Database.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		// Ya existe un usuario con el mismo nombre de usuario
		// Verificar si la contraseña coincide
		if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password)); err == nil {
			fmt.Println("El usuario ya está registrado. Puede continuar operando.")
			return nil // No es un error, simplemente indicamos que el usuario ya está registrado
		} else {
			fmt.Println("La contraseña proporcionada no coincide.")
			return err // Devuelve un error indicando que la contraseña no coincide
		}
	}

	// Hashear la contraseña
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("Error al hashear la contraseña:", err)
		return err
	}

	user.Password = hashedPassword

	// Guardar el nuevo usuario en la base de datos
	if err := db.Database.Save(&user).Error; err != nil {
		fmt.Println("Error al crear el usuario en la base de datos:", err)
		return err
	}

	return nil
}

// func Login(conn net.Conn, username, password string) error {
// 	// Buscar el usuario en la base de datos por nombre de usuario
// 	var user models.User
// 	if err := db.Database.Where("username = ?", username).First(&user).Error; err != nil {
// 		// Si hay un error al buscar el usuario, devolver un error
// 		fmt.Println("Error al buscar el usuario:", err)
// 		return err
// 	}

// 	// Verificar si la contraseña ingresada coincide con la contraseña almacenada hasheada
// 	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)); err != nil {
// 		fmt.Println("Los password no coinciden")
// 		fmt.Println("Los password no coinciden")

// 	}

// 	// Si las credenciales son válidas, el inicio de sesión fue exitoso
// 	fmt.Println("Inicio de sesión exitoso.")
// 	return nil
// }
