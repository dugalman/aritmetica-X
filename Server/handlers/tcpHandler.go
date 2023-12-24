package handlers

import (
	"encoding/gob"
	"examen_server/operations"
	"fmt"
	"net"
)

func HandleClient(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	for {
		var request operations.OperationRequest
		err := decoder.Decode(&request)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Conexion finalizada del lado del cliente")
			} else {
				fmt.Println("Error al decodificar la solicitud:", err)
			}
			return
		}

		response := operations.Operation(request)
		err = encoder.Encode(response)
		if err != nil {
			fmt.Println("Error al enviar la respuesta:", err)
			return
		}
	}
}
