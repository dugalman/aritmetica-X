package controller

import (
	"examen_server/db"
	"examen_server/models"
	"examen_server/utils"
	"fmt"
	"net"
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
	if err := db.Database.Where("username = ?", username).First(&existingUser).Error; err == nil {
		// El usuario ya existe
		return fmt.Errorf("el usuario '%s' ya está registrado", username)
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
		Op:         op,
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
