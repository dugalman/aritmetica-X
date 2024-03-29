package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"

	"examen_client/utils"
)

// CreateUser representa los datos de un nuevo usuario
type CreateUser struct {
	Username string
	Password string
}

// OperationRequest representa la estructura de la solicitud del cliente
type OperationRequest struct {
	Num1 float64
	Num2 float64
	Op   utils.OperationType // Usa el tipo definido en utils
	User CreateUser
}

// OperationResponse representa la estructura de la respuesta del servidor
type OperationResponse struct {
	Result    float64
	ErrorCode int
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error al conectar al servidor:", err)
		return
	}
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	gob.Register(CreateUser{})
	gob.Register(OperationRequest{})
	gob.Register(OperationResponse{})

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese su nombre de usuario para registrarse o iniciar sesion: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Ingrese su contraseña: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// Crear una solicitud de registro de usuario
	registerRequest := CreateUser{
		Username: username,
		Password: password,
	}

	// Enviar la solicitud de registro al servidor
	err = encoder.Encode(registerRequest)
	if err != nil {
		fmt.Println("Error al enviar la solicitud de registro:", err)
		return
	}

	var num1, num2 float64
	var operation utils.OperationType

	// El menú de opciones
	fmt.Println("Seleccione el tipo de operacion")
	fmt.Println("1.SUM")
	fmt.Println("2.MINUS")
	fmt.Println("3.DIV")
	fmt.Println("4.MULT")
	fmt.Println("5.SIN")
	fmt.Println("6.LOG")
	fmt.Println("7.EXP")
	fmt.Println("8.SQR")

	// El bucle para el caso de que ingresen un número no válido
	for {
		fmt.Print("Ingrese el tipo de operación: ")
		_, err = fmt.Scanf("%d\n", &operation)
		if err != nil || operation < 1 || operation > 8 {
			fmt.Println("Tipo de operación no válida.")
			continue
		}
		break
	}

	// Input del primer número
	for {
		fmt.Print("Ingrese el primer número: ")
		_, err := fmt.Scan(&num1)
		if err != nil {
			fmt.Println("Error: Por favor, ingrese un número válido.")
			continue
		}
		break
	}

	// Comprobar si la operación requiere un segundo número
	if operation != utils.SIN && operation != utils.LOG && operation != utils.EXP && operation != utils.SQR {
		for {
			fmt.Print("Ingrese el segundo número: ")
			_, err := fmt.Scan(&num2)
			if err != nil {
				fmt.Println("Error: Por favor, ingrese un número válido.")
				continue
			}
			break
		}
	} else {
		// Establecer num2 a 0 para operaciones unarias
		num2 = 0
	}

	// Verificar en base a la operación y el segundo número para no dividir o multiplicar por cero
	if operation == utils.DIV && num2 == 0 {
		fmt.Println("Error: No se puede dividir entre cero")
		return
	}
	if operation == utils.MULT && (num1 == 0 || num2 == 0) {
		fmt.Println("Error: No se puede multiplicar entre cero")
		return
	}

	request := OperationRequest{
		Num1: num1,
		Num2: num2,
		Op:   operation,
	}

	// Enviar la solicitud de operación al servidor
	err = encoder.Encode(request)
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return
	}

	var response OperationResponse
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println("Error al recibir la respuesta:", err)
		return
	}

	if response.ErrorCode == 1 {
		fmt.Printf("Error: %d\n", response.ErrorCode)
	}

	if response.ErrorCode != 0 {
		fmt.Printf("Error: %d\n", response.ErrorCode)
	} else {
		fmt.Printf("Resultado de %g %s %g = %g\n", request.Num1, utils.OperationSymbol(request.Op), request.Num2, response.Result)
	}
}
