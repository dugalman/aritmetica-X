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

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error al conectar al servidor:", err)
		return
	}
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	request := OperationRequest{
		Num1: 10,
		Num2: 5,
		Op:   '+',
	}

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

	if response.ErrorCode != 0 {
		fmt.Printf("Error: %d\n", response.ErrorCode)
	} else {
		fmt.Printf("Resultado de %d %c %d = %.2f\n", response.Request.Num1, response.Request.Op, response.Request.Num2, response.Result)
	}
}
