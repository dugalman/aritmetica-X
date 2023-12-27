package handlers

import (
	"encoding/gob"
	"examen_server/users"
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

func handleClient(w *UserResponse, r UserRequest) {
	switch r.Operation.Op {
	case '+':
		w.Result.Operation.Result = float64(r.Operation.Num1 + r.Operation.Num2)
	case '-':
		w.Result.Operation.Result = float64(r.Operation.Num1 - r.Operation.Num2)
	// Puedes agregar más operaciones aquí según sea necesario

	default:
		w.Result.Operation.ErrorCode = -1 // Código de error para operación no válida
	}
	w.Result.Operation.Request = r.Operation
}

func UserHandler(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	for {
		var request UserRequest
		err := decoder.Decode(&request)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("\nConexion finalizada del lado del cliente")
			} else {
				fmt.Println("\nError al decodificar la solicitud:", err)
			}
			return
		}

		var response UserResponse

		switch request.Path {
		case "sing-in":
			if singInUser(&response, request) {
				handleClient(&response, request)
			}
		case "sing-up":
			singUpUser(&response, request)

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

// singInUser si el usuario ya existe en la "DB" permite el ingreso
func singInUser(w *UserResponse, r UserRequest) bool {
	newUser := getUser(r)
	ok := users.DefaultUserService.VerifyUser(newUser)
	if !ok {
		// Error: User Sign-in Failure
		w.ErrorCode = -20
		return false
	}
	w.Result.MsgUser = "User Sign-in Success"
	fmt.Println("User Sign-in Success")
	return true
}

// singUpUser da de alta un nuevo usuario
func singUpUser(w *UserResponse, r UserRequest) {
	newUser := getUser(r)
	err := users.DefaultUserService.CreateUser(newUser)
	if err != nil {
		// Error: New User Sign-up Failure
		w.ErrorCode = -10
		return
	}
	w.Result.MsgUser = "New User Sign-up"
	return
}

// getUser carga el usuario que viene desde el cliente a User Struct
func getUser(r UserRequest) users.User {
	return users.User{
		Email:    r.Payload.Email,
		Password: r.Payload.Password,
	}
}
