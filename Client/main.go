package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
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

func operationSymbol(op OperationType) string {
	switch op {
	case SUM:
		return "+"
	case MINUS:
		return "-"
	case DIV:
		return "/"
	case MULT:
		return "*"
	case SIN:
		return "sin"
	case LOG:
		return "log"
	case EXP:
		return "exp"
	case SQR:
		return "sqr"
	default:
		return "?"
	}

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
	var operation OperationType

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

	//Verifico en este caso para que no se ingrese el segundo numero en base a la operacion

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
	if operation != SIN && operation != LOG && operation != EXP && operation != SQR {
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

	if operation == DIV && num2 == 0 {
		fmt.Println("Error: No se puede dividir entre cero")
		return
	}
	if operation == MULT && (num1 == 0 || num2 == 0) {
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
		fmt.Printf("Resultado de %g %s %g = %g\n", response.Request.Num1, operationSymbol(response.Request.Op), response.Request.Num2, response.Result)
	}
}
