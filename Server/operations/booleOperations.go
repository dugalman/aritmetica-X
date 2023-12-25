package operations

import (
	"errors"
	"fmt"
)

// OperationType representa el tipo de operacion a ejecutar
type OperationType int
type booleType bool

const (
	AND  OperationType = 1
	OR   OperationType = 2
	NOT  OperationType = 3
	XOR  OperationType = 4
	NAND OperationType = 5
)

// TODO: Códigos de Error Personalizados
const (
	OK                 int = 0
	OverflowSum        int = 1
	InvalidInputParam  int = 2
	DivisionByZero     int = 3
	InvalidLogarithm   int = 4
	NegativeSquareRoot int = 5
	InvalidOperation   int = 6
	InvalidBoolean     int = 7
)

// OperationRequest representa la estructura de la solicitud del cliente
type OperationRequest struct {
	Input1 booleType
	Input2 booleType
	Op     OperationType
}

// OperationResult representa el resultado de la operación
type OperationResult struct {
	Value booleType
}

// OperationResponse representa la estructura de la respuesta del servidor
type OperationResponse struct {
	Request   OperationRequest
	Result    OperationResult
	ErrorCode int
}

func Operation(request OperationRequest) OperationResponse {
	var response OperationResponse

	//Validation
	if !isBoolean(request.Input1, request.Input2) {
		response.ErrorCode = InvalidBoolean
	}

	fmt.Println(request)
	response.Request = request

	switch request.Op {
	case AND:
		andResult, _ := andOperation(request.Input1, request.Input2)
		response.Result.Value = andResult
		response.ErrorCode = OK
	case OR:
		orResult, _ := orOperation(request.Input1, request.Input2)
		response.Result.Value = orResult
		response.ErrorCode = OK
	case NOT:
		notResult, _ := notOperation(request.Input1)
		response.Result.Value = notResult
		response.ErrorCode = OK
	case XOR:
		xorResult, _ := xorOperation(request.Input1, request.Input2)
		response.Result.Value = xorResult
		response.ErrorCode = OK
	case NAND:
		nandResult, _ := nandOperation(request.Input1, request.Input2)
		response.Result.Value = nandResult
		response.ErrorCode = OK
	default:
		response.ErrorCode = InvalidOperation
	}

	return response
}

func andOperation(operands ...booleType) (booleType, error) {
	if len(operands) < 2 {
		return false, errors.New("AND requiere al menos dos operandos booleanos")
	}
	var result booleType
	result = true
	for _, operand := range operands {
		result = result && operand
	}
	return result, nil
}

func orOperation(operands ...booleType) (booleType, error) {
	if len(operands) < 2 {
		return false, errors.New("OR requiere al menos dos operandos booleanos")
	}
	var result booleType
	result = false
	for _, operand := range operands {
		result = result || operand
	}
	return result, nil
}

func notOperation(operands ...booleType) (booleType, error) {
	if len(operands) != 1 {
		return false, errors.New("NOT requiere exactamente un operando booleano")
	}
	return !operands[0], nil
}

func xorOperation(operands ...booleType) (booleType, error) {
	if len(operands) < 2 {
		return false, errors.New("XOR requiere al menos dos operandos booleanos")
	}
	var result booleType
	result = false
	for _, operand := range operands {
		result = result != operand
	}
	return result, nil
}

func nandOperation(operands ...booleType) (booleType, error) {
	if len(operands) < 2 {
		return false, errors.New("NAND requiere al menos dos operandos booleanos")
	}
	result, err := andOperation(operands...)
	if err != nil {
		return false, err
	}
	return !result, nil
}

func isBoolean(values ...booleType) booleType {
	var result booleType
	for _, value := range values {
		result = value == true || value == false
	}
	return result
}
