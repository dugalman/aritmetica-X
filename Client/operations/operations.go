package operations

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

// OperationRequest representa la estructura de la solicitud del cliente
type OperationRequest struct {
	Input1 booleType
	Input2 booleType
	Op     OperationType
}

type OperationResult struct {
	Value booleType
}

// OperationResponse representa la estructura de la respuesta del servidor
type OperationResponse struct {
	Request   OperationRequest
	Result    OperationResult
	ErrorCode int
}

// opSymbol retorna el simbolo de la operacion correspondiente
func OpSymbol(op OperationType) string {
	switch op {
	case AND:
		return "AND"
	case OR:
		return "OR"
	case NOT:
		return "NOT"
	case XOR:
		return "XOR"
	case NAND:
		return "NAND"
	default:
		return "operacion no valida"
	}
}
