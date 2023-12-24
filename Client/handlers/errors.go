package handlers

import "fmt"

func HandleError(errorCode int) {
	switch errorCode {
	case -2:
		fmt.Println("Error: No se puede dividir por cero.")
	case -3:
		fmt.Println("Error: Logaritmo no válido para números no positivos.")
	case -4:
		fmt.Println("Error: No se puede calcular la raíz cuadrada de un número negativo.")
	default:
		fmt.Printf("Error desconocido: %d\n", errorCode)
	}
}
