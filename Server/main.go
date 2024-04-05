package main

import (
	"encoding/gob"
	"examen_server/controller"
	"examen_server/models"
	"fmt"
	"math"
	"net"
)

// OperationType representa el tipo de operación
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

type OperationRequest struct {
	Num1 float64
	Num2 float64
	Op   OperationType // Tipo de operación
}

// CreateUser representa los datos de un nuevo usuario
type CreateUser struct {
	Username   string
	Password   string
	Operations []UserOperation
}

// UserOperation representa una operación realizada por el usuario
type UserOperation struct {
	Operation string
	Result    float64
}

// OperationResponse representa la estructura de la respuesta del servidor

type OperationResponse struct {
	Result     float64
	ErrorCode  int
	ResultBool bool
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	gob.Register(models.User{})
	gob.Register(CreateUser{})
	gob.Register(OperationRequest{})
	gob.Register(OperationResponse{})
	gob.Register(models.UserOperation{})
	gob.Register(UserOperation{})

	var newUser CreateUser
	var userOperations []models.UserOperation

	for {
		var request OperationRequest
		err := decoder.Decode(&request)
		if err != nil {
			fmt.Println("Error al decodificar la solicitud:", err)
			return
		}
		var response OperationResponse

		// Decodifica la solicitud de creación de usuario solo si es la primera vez que se recibe
		if newUser.Username == "" {
			err := decoder.Decode(&newUser)
			if err != nil {
				fmt.Println("Error al decodificar la solicitud de creación de usuario:", err)
				return
			}

			for _, op := range newUser.Operations {
				userOp := models.UserOperation{Operation: op.Operation, Result: op.Result}
				userOperations = append(userOperations, userOp)
				fmt.Println(userOperations)
			}
		}

		fmt.Println("Este es el request del server", request)

		switch request.Op {
		case SUM:
			response.Result = request.Num1 + request.Num2
		case MINUS:
			response.Result = request.Num1 - request.Num2
		case DIV:
			if request.Num2 == 0 {
				response.ErrorCode = 1 // División por cero
			} else {
				response.Result = request.Num1 / request.Num2
			}
		case MULT:
			response.Result = request.Num1 * request.Num2
		case SIN:
			response.Result = math.Sin(request.Num1)
		case LOG:
			if request.Num1 <= 0 {
				response.ErrorCode = 2 // Logaritmo de número no válido
			} else {
				response.Result = math.Log(request.Num1)
			}
		case EXP:
			response.Result = math.Exp(request.Num1)
		case SQR:
			if request.Num1 < 0 {
				response.ErrorCode = 3 // Raíz cuadrada de número negativo
			} else {
				response.Result = math.Sqrt(request.Num1)
			}
		case AND:
			if request.Num1 == request.Num2 {
				response.ResultBool = true

			} else {
				response.ResultBool = false
			}
		case OR:
			if request.Num1 == request.Num2 {
				response.ResultBool = true

			} else {
				response.ResultBool = false
			}
		case NOT:
			if request.Num1 > 0 {
				response.ResultBool = true

			} else {
				response.ResultBool = false
			}

		case XOR:
			if request.Num1 == request.Num2 {
				response.ResultBool = false

			} else {
				response.ResultBool = true
			}

		case NAND:
			if request.Num1 == request.Num2 {
				response.ResultBool = false
			}
			if request.Num1 != request.Num2 {
				response.ResultBool = true
			}
		default:
			response.ErrorCode = -1 // Operación no válida
		}

		err = encoder.Encode(response)
		if err != nil {
			fmt.Println("Error al enviar la respuesta:", err)
			return
		}

		// Registro de usuario aquí
		err = controller.Register(conn, newUser.Username, newUser.Password, userOperations, request.Num1, request.Num2, int(request.Op))
		if err != nil {
			fmt.Println("Error al registrar usuario:", err)
			return
		}
	}

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor esperando conexiones en el puerto 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}

		go handleClient(conn)
	}
}
