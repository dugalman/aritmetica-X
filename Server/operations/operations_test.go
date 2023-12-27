package operations_test

import (
	"examen_server/operations"
	"math"
	"testing"
)

func TestOperationSum(t *testing.T) {
	request := operations.OperationRequest{Num1: 3, Num2: 4, Op: operations.SUM}
	response := operations.Operation(request)

	expectedResult := float64(9)
	expectedErrorCode := operations.OK

	if response.Result != expectedResult || response.ErrorCode != expectedErrorCode {
		t.Errorf("Error en la suma. Se esperaba resultado: %v, código de error: %v, pero se obtuvo resultado: %v, código de error: %v",
			expectedResult, expectedErrorCode, response.Result, response.ErrorCode)
	}
}

func TestOperationMinus(t *testing.T) {
	request := operations.OperationRequest{Num1: 10, Num2: 4, Op: operations.MINUS}
	response := operations.Operation(request)

	expectedResult := float64(6)
	expectedErrorCode := operations.OK

	if response.Result != expectedResult || response.ErrorCode != expectedErrorCode {
		t.Errorf("Error en la resta. Se esperaba resultado: %v, código de error: %v, pero se obtuvo resultado: %v, código de error: %v",
			expectedResult, expectedErrorCode, response.Result, response.ErrorCode)
	}
}

func TestOperationDivisionByZero(t *testing.T) {
	request := operations.OperationRequest{Num1: 5, Num2: 0, Op: operations.DIV}
	response := operations.Operation(request)

	expectedErrorCode := operations.DivisionByZero

	if response.ErrorCode != expectedErrorCode {
		t.Errorf("Error en la división por cero. Se esperaba código de error: %v, pero se obtuvo código de error: %v",
			expectedErrorCode, response.ErrorCode)
	}
}

func TestOperationMultiplication(t *testing.T) {
	request := operations.OperationRequest{Num1: 5, Num2: 3, Op: operations.MULT}
	response := operations.Operation(request)

	expectedResult := float64(15)
	expectedErrorCode := operations.OK

	if response.Result != expectedResult || response.ErrorCode != expectedErrorCode {
		t.Errorf("Error en la multiplicación. Se esperaba resultado: %v, código de error: %v, pero se obtuvo resultado: %v, código de error: %v",
			expectedResult, expectedErrorCode, response.Result, response.ErrorCode)
	}
}

func TestOperationSin(t *testing.T) {
	request := operations.OperationRequest{Num1: 30, Op: operations.SIN}
	response := operations.Operation(request)

	expectedResult := math.Sin(float64(30))
	expectedErrorCode := operations.OK

	if response.Result != expectedResult || response.ErrorCode != expectedErrorCode {
		t.Errorf("Error en la función seno. Se esperaba resultado: %v, código de error: %v, pero se obtuvo resultado: %v, código de error: %v",
			expectedResult, expectedErrorCode, response.Result, response.ErrorCode)
	}
}
