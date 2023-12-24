package operations

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

// opSymbol retorna el simbolo de la operacion correspondiente
func OpSymbol(op OperationType) string {
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
		return "sqrt"
	default:
		return "operacion no valida"
	}
}
