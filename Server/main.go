package main

import (
	"encoding/gob"
	"fmt"
	"math"
	"net"
)

// OperationRequest representa la estructura de la solicitud del cliente
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
)

type OperationRequest struct {
	Num1 float64
	Num2 float64
	Op   OperationType
}

// OperationResponse representa la estructura de la respuesta del servidor
type OperationResponse struct {
	Request   OperationRequest
	Result    float64
	ErrorCode int
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	for {
		var request OperationRequest
		err := decoder.Decode(&request)
		if err != nil {
			fmt.Println("Error al decodificar la solicitud:", err)
			return
		}

		var response OperationResponse

		switch request.Op {
		case SUM:
			response.Result = float64(request.Num1 + request.Num2)
		case MINUS:
			response.Result = float64(request.Num1 - request.Num2)
		case DIV:
			response.Result = float64(request.Num1) / float64(request.Num2)
		case MULT:
			response.Result = float64(request.Num1) * float64(request.Num2)
		case SIN:
			response.Result = math.Sin(float64(request.Num1))
		case LOG:
			response.Result = math.Log(float64(request.Num1))
		case EXP:
			response.Result = math.Exp(float64(request.Num1))
		case SQR:
			response.Result = math.Sqrt(float64(request.Num1))

		// Puedes agregar más operaciones aquí según sea necesario

		default:
			response.ErrorCode = -1 // Código de error para operación no válida
		}

		response.Request = request
		err = encoder.Encode(response)
		if err != nil {
			fmt.Println("Error al enviar la respuesta:", err)
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
