package handlers

import (
	"fmt"
)

// Códigos de Error Personalizados
const (
	OK                 int = 0
	OverflowSum        int = 1
	InvalidInputParam  int = 2
	DivisionByZero     int = 3
	InvalidLogarithm   int = 4
	NegativeSquareRoot int = 5
	InvalidOperation   int = 6
)

func HandleError(errorCode int) {
	switch errorCode {
	case DivisionByZero:
		fmt.Println("Error: No se puede dividir por cero.")
	case InvalidLogarithm:
		fmt.Println("Error: Logaritmo no válido para números no positivos.")
	case NegativeSquareRoot:
		fmt.Println("Error: No se puede calcular la raíz cuadrada de un número negativo.")
	case OverflowSum:
		fmt.Println("Error: Overflow en la suma.")
	case InvalidInputParam:
		fmt.Println("Error: Parámetro de entrada no válido.")
	case InvalidOperation:
		fmt.Println("Error: Operación no válida.")
	default:
		fmt.Printf("Error desconocido: %d\n", errorCode)
	}
}
