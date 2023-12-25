package main

import (
	"encoding/gob"
	"examen_client/operations"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error al conectar al servidor:", err)
		return
	}
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	request := operations.OperationRequest{
		Input1: true,
		Input2: true,
		Op:     operations.NAND,
	}

	err = encoder.Encode(request)
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return
	}

	var response operations.OperationResponse
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println("Error al recibir la respuesta:", err)
		return
	}

	if response.ErrorCode != 0 {
		fmt.Printf("Error: %d\n", response.ErrorCode)
	} else {
		if response.Request.Op == operations.NOT {
			fmt.Printf("Resultado de NOT_( %v ) = %.v\n", response.Request.Input1, response.Result.Value)
		} else {
			fmt.Printf("Resultado de %v %v %v = %.v\n", response.Request.Input1, operations.OpSymbol(response.Request.Op), response.Request.Input2, response.Result.Value)
		}

	}
}
