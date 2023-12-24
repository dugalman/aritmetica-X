package operations

import (
	"math"
)

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

func Operation(request OperationRequest) OperationResponse {
	var response OperationResponse

	//Validacion de parametros de entrada
	if !isValidNumber(request.Num1, request.Num2) {
		response.ErrorCode = InvalidInputParam
	}

	switch request.Op {
	case SUM:
		sumResult := float64(request.Num1 + request.Num2)
		if sumResult > math.MaxFloat64 {
			response.ErrorCode = OverflowSum
		} else {
			response.Result = sumResult
			response.ErrorCode = OK
		}
	case MINUS:
		minusResult := float64(request.Num1 - request.Num2)
		response.Result = minusResult
		response.ErrorCode = OK

	case DIV:
		if request.Num2 != 0 {
			divResult := float64(request.Num1) / float64(request.Num2)
			response.Result = divResult
			response.ErrorCode = OK
		} else {
			response.ErrorCode = DivisionByZero
		}
	case MULT:
		multResult := float64(request.Num1 * request.Num2)
		response.Result = multResult
		response.ErrorCode = OK
	case SIN:
		sinResult := math.Sin(float64(request.Num1))
		response.Result = sinResult
		response.ErrorCode = OK
	case LOG:
		if request.Num1 > 0 && request.Num2 > 0 {
			logResult := math.Log(float64(request.Num1)) / math.Log(float64(request.Num2))
			response.Result = logResult
			response.ErrorCode = OK
		} else {
			response.ErrorCode = InvalidLogarithm
		}
	case EXP:
		expResult := math.Pow(float64(request.Num1), float64(request.Num2))
		response.Result = expResult
		response.ErrorCode = OK
	case SQR:
		if request.Num1 >= 0 {
			sqrtResult := math.Sqrt(float64(request.Num1))
			response.Result = sqrtResult
			response.ErrorCode = OK
		} else {
			response.ErrorCode = NegativeSquareRoot
		}
	default:
		response.ErrorCode = InvalidOperation
	}
	response.Request = request
	return response
}

func isValidNumber(num1, num2 interface{}) bool {
	_, esNum1 := num1.(float64)
	_, esNum2 := num2.(float64)
	return esNum1 && esNum2
}
