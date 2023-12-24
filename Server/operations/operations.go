package operations

import "math"

// OperationType representa el tipo de operacion a ejecutar
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

// OperationRequest representa la estructura de la solicitud del cliente
type OperationRequest struct {
	Num1 int
	Num2 int
	Op   OperationType
}

// OperationResponse representa la estructura de la respuesta del servidor
type OperationResponse struct {
	Request   OperationRequest
	Result    float64
	ErrorCode int
}

func Operation(request OperationRequest) (float64, int) {
	switch request.Op {
	case SUM:
		return float64(request.Num1 + request.Num2), 0
	case MINUS:
		return float64(request.Num1 - request.Num2), 0
	case DIV:
		if request.Num2 != 0 {
			return float64(request.Num1) / float64(request.Num2), 0
		} else {
			return 0, -2 // Código de error para division
		}
	case MULT:
		return float64(request.Num1 * request.Num2), 0
	case SIN:
		return math.Sin(float64(request.Num1)), 0
	case LOG:
		if request.Num1 > 0 && request.Num2 > 0 {
			return math.Log(float64(request.Num1)) / math.Log(float64(request.Num2)), 0
		} else {
			return 0, -3 // Código de error para logaritmo no valido
		}
	case EXP:
		return math.Pow(float64(request.Num1), float64(request.Num2)), 0
	case SQR:
		if request.Num1 >= 0 {
			return math.Sqrt(float64(request.Num1)), 0
		} else {
			return 0, -4 // Código de error para raiz cuadrada - numero negativo
		}
	default:
		return 0, -1 // Código de error para operación no valida
	}
}
