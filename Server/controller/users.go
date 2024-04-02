package controller

import (
	"examen_server/db"
	"examen_server/models"
	"examen_server/utils"
	"fmt"
	"net"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type OperationType int

const (
	SUM   OperationType = 1
	MINUS OperationType = 2
	DIV   OperationType = 3
	MULT  OperationType = 4
	SIN   OperationType = 5
	LOG   OperationType = 6
	EXP   OperationType = 7
	SQR   OperationType = 8
	AND   OperationType = 9
	OR    OperationType = 10
	NOT   OperationType = 11
	XOR   OperationType = 12
	NAND  OperationType = 13
)

func Register(conn net.Conn, username, password string, operations []models.UserOperation, num1, num2 float64, op int) error {
	// Verificar si el usuario ya está registrado
	var existingUser models.User
	if err := db.Database.Where("username = ?", username).Preload("Operations").First(&existingUser).Error; err == nil {
		// El usuario ya existe
		// Verificar si la contraseña es correcta
		if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password)); err != nil {
			// La contraseña es incorrecta, no almacenar nada
			fmt.Println("Contraseña incorrecta para el usuario existente.")
			return nil
		}

		// La contraseña es correcta, agregar las nuevas operaciones
		existingUser.Operations = append(existingUser.Operations, operations...)

		// Actualizar el usuario en la base de datos con las nuevas operaciones
		if err := db.Database.Save(&existingUser).Error; err != nil {
			fmt.Println("Error al actualizar las operaciones del usuario:", err)
			return err
		}

		fmt.Println("Operaciones añadidas al usuario existente.")
		return nil
	}

	// Hashear la contraseña
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("Error al hashear la contraseña:", err)
		return err
	}

	// Crear un nuevo usuario
	newUser := models.User{
		Num1:       num1,
		Num2:       num2,
		Op:         strconv.Itoa(op),
		Username:   username, // Usar el nombre de usuario proporcionado
		Password:   hashedPassword,
		Operations: operations,
	}

	// Guardar el nuevo usuario en la base de datos
	if err := db.Database.Create(&newUser).Error; err != nil {
		fmt.Println("Error al crear el usuario en la base de datos:", err)
		return err
	}

	fmt.Println("Usuario registrado exitosamente.")
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
