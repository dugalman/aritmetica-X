package main

import (
	"encoding/gob"
	"examen_client/users"

	"fmt"
	"net"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
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

// PayloadResponse representa la estructura de la respuesta del servidor
type PayloadResponse struct {
	MsgUser   string
	Operation OperationResponse
}

// UserRequest representa la estructura de la solicitud del cliente
type UserRequest struct {
	Payload   users.User
	Operation OperationRequest
	Path      string
}

// UserResponse representa la estructura de la respuesta del servidor
type UserResponse struct {
	Request   UserRequest
	Result    PayloadResponse
	ErrorCode int
}

func main() {
	conn, err := net.Dial(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error al conectar al servidor:", err)
		return
	}
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	//request := OperationRequest{
	//	Num1: 10,
	//	Num2: 5,
	//	Op:   '+',
	//}
	//err = encoder.Encode(request)
	//if err != nil {
	//	fmt.Println("Error al enviar la solicitud:", err)
	//	return
	//}

	//var response OperationResponse
	//err = decoder.Decode(&response)
	//if err != nil {
	//	fmt.Println("Error al recibir la respuesta:", err)
	//	return
	//}

	//if response.ErrorCode != 0 {
	//	fmt.Printf("Error: %d\n", response.ErrorCode)
	//} else {
	//	fmt.Printf("Resultado de %d %c %d = %.2f\n", response.Request.Num1, response.Request.Op, response.Request.Num2, response.Result)
	//}

	// USERS
	userRequest := UserRequest{
		Payload: users.User{
			Email:    "test_01@test.com",
			Password: "123qwe",
		},
		Operation: OperationRequest{
			Num1: 10,
			Num2: 5,
			Op:   '-',
		},
		Path: "sing-in",
	}

	err = encoder.Encode(userRequest)
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return
	}

	var userResponse UserResponse
	err = decoder.Decode(&userResponse)
	if err != nil {
		fmt.Println("Error al recibir la respuesta:", err)
		return
	}

	if userResponse.ErrorCode != 0 {
		fmt.Printf("Error: %d\n", userResponse.ErrorCode)
	} else {
		fmt.Printf("Status: %v\n", userResponse.Result.MsgUser)

		if userRequest.Path == "sing-in" {
			if userResponse.Result.Operation.ErrorCode != 0 {
				fmt.Printf("Error: %d\n", userResponse.Result.Operation.ErrorCode)
			} else {
				fmt.Printf("Resultado de %d %c %d = %.2f\n", userResponse.Result.Operation.Request.Num1, userResponse.Result.Operation.Request.Op, userResponse.Result.Operation.Request.Num2, userResponse.Result.Operation.Result)
			}
		}
	}
}
