package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

// OperationRequest representa la estructura de la solicitud del cliente
type OperationRequest struct {
	Num1 int
	Num2 int
	Op   rune
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
		case '+':
			response.Result = float64(request.Num1 + request.Num2)
		case '-':
			response.Result = float64(request.Num1 - request.Num2)
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
