package main

import (
	"encoding/gob"
	"examen_client/utils"
	"fmt"
	"net"
	"strconv"
)

// OperationRequest representa la estructura de la solicitud del cliente
type OperationRequest struct {
	Num1 float64
	Num2 float64
	Op   utils.OperationType // Usa el tipo definido en utils
}

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

	var num1Str, num2Str string
	var num1, num2 float64
	var operation utils.OperationType

	//El menu de opciones
	fmt.Println("Seleccione el tipo de operacion")
	fmt.Println("1.SUM")
	fmt.Println("2.MINUS")
	fmt.Println("3.DIV")
	fmt.Println("4.MULT")
	fmt.Println("5.SIN")
	fmt.Println("6.LOG")
	fmt.Println("7.EXP")
	fmt.Println("8.SQR")

	//El for para el caso que ingresen un numero no valido
	for {
		_, err = fmt.Scanf("%d\n", &operation)
		if err != nil {
			fmt.Println("Error al leer el tipo de operacion", operation)
		}
		if operation < 1 || operation > 8 {
			fmt.Println("Tipo de operación no válida.")
			fmt.Println("Seleccione el tipo de operacion")
			fmt.Println("1.SUM")
			fmt.Println("2.MINUS")
			fmt.Println("3.DIV")
			fmt.Println("4.MULT")
			fmt.Println("5.SIN")
			fmt.Println("6.LOG")
			fmt.Println("7.EXP")
			fmt.Println("8.SQR")
			continue
		}
		break

	}

	for {
		fmt.Println("Ingrese el primer número:")
		_, err := fmt.Scan(&num1Str)
		if err != nil {
			fmt.Println("Error: Por favor, ingrese un número válido.")
			continue
		}

		num1, err = strconv.ParseFloat(num1Str, 64)
		if err != nil {
			fmt.Println("Error: Por favor, ingrese un número válido.")
			continue
		}
		break
	}

	// Comprueba si la operación requiere un segundo número
	if operation != utils.SIN && operation != utils.LOG && operation != utils.EXP && operation != utils.SQR {
		for {
			fmt.Println("Ingrese el segundo número:")
			_, err := fmt.Scan(&num2Str)
			if err != nil {
				fmt.Println("Error: Por favor, ingrese un número válido.")
				continue
			}

			num2, err = strconv.ParseFloat(num2Str, 64)
			if err != nil {
				fmt.Println("Error: Por favor, ingrese un número válido.")
				continue
			}
			break
		}
	} else {
		// Establece num2 a 0 para operaciones unarias
		num2 = 0
	}

	//Verifico en base a la operacion y el segundo numero para no dividir o multiplicar por cero
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
		fmt.Printf("Resultado de %g %s %g = %g\n", response.Request.Num1, utils.OperationSymbol(response.Request.Op), response.Request.Num2, response.Result)
	}

}
