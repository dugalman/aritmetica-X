package utils

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
	AND   OperationType = 9
	OR    OperationType = 10
	NOT   OperationType = 11
	XOR   OperationType = 12
	NAND  OperationType = 13
)

// OperationSymbol devuelve el símbolo de la operación
func OperationSymbol(op OperationType) string {
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
	case AND:
		return "and"
	case OR:
		return "or"
	case NOT:
		return "not"
	case XOR:
		return "xor"
	case NAND:
		return "nand"
	default:
		return "?"
	}
}
