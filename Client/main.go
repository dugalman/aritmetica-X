package main

import (
	"encoding/gob"
	"examen_client/handlers"
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
		Num1: -10,
		Num2: 0,
		Op:   operations.SQR,
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
		handlers.HandleError(response.ErrorCode)
	} else {
		if response.Request.Op == operations.LOG {
			fmt.Printf("Resultado del Log_%d( %d ) = %.2f\n", response.Request.Num2, response.Request.Num1, response.Result)
		} else if response.Request.Op == operations.SQR {
			fmt.Printf("Resultado de la raiz cuadrada de %d = %.2f\n", response.Request.Num1, response.Result)
		} else {
			fmt.Println(response)
			fmt.Printf("Resultado de %d %s %d = %.2f\n", response.Request.Num1, operations.OpSymbol(response.Request.Op), response.Request.Num2, response.Result)
		}
	}
}
