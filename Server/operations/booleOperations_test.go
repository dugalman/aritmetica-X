package operations_test

import (
	"examen_server/operations"
	"testing"
)

func TestOperationAND(t *testing.T) {
	request := operations.OperationRequest{Input1: true, Input2: true, Op: operations.AND}
	response := operations.Operation(request)

	expectedErrorCode := operations.OK

	if response.Result.Value != true || response.ErrorCode != expectedErrorCode {
		t.Errorf("Error en la operación AND. Se esperaba resultado: %v, código de error: %v, pero se obtuvo resultado: %v, código de error: %v",
			true, expectedErrorCode, response.Result.Value, response.ErrorCode)
	}
}

func TestOperationOR(t *testing.T) {
	request := operations.OperationRequest{Input1: true, Input2: false, Op: operations.OR}
	response := operations.Operation(request)

	//expectedResult := true
	expectedErrorCode := operations.OK

	if response.Result.Value != true || response.ErrorCode != expectedErrorCode {
		t.Errorf("Error en la operación OR. Se esperaba resultado: %v, código de error: %v, pero se obtuvo resultado: %v, código de error: %v",
			true, expectedErrorCode, response.Result.Value, response.ErrorCode)
	}
}

func TestOperationNOT(t *testing.T) {
	request := operations.OperationRequest{Input1: true, Op: operations.NOT}
	response := operations.Operation(request)

	//expectedResult := false
	expectedErrorCode := operations.OK

	if response.Result.Value != false || response.ErrorCode != expectedErrorCode {
		t.Errorf("Error en la operación NOT. Se esperaba resultado: %v, código de error: %v, pero se obtuvo resultado: %v, código de error: %v",
			false, expectedErrorCode, response.Result.Value, response.ErrorCode)
	}
}

func TestOperationXOR(t *testing.T) {
	request := operations.OperationRequest{Input1: true, Input2: false, Op: operations.XOR}
	response := operations.Operation(request)

	//expectedResult := true
	expectedErrorCode := operations.OK

	if response.Result.Value != true || response.ErrorCode != expectedErrorCode {
		t.Errorf("Error en la operación XOR. Se esperaba resultado: %v, código de error: %v, pero se obtuvo resultado: %v, código de error: %v",
			true, expectedErrorCode, response.Result.Value, response.ErrorCode)
	}
}

func TestOperationNAND(t *testing.T) {
	request := operations.OperationRequest{Input1: true, Input2: true, Op: operations.NAND}
	response := operations.Operation(request)

	//expectedResult := false
	expectedErrorCode := operations.OK

	if response.Result.Value != false || response.ErrorCode != expectedErrorCode {
		t.Errorf("Error en la operación NAND. Se esperaba resultado: %v, código de error: %v, pero se obtuvo resultado: %v, código de error: %v",
			false, expectedErrorCode, response.Result.Value, response.ErrorCode)
	}
}
