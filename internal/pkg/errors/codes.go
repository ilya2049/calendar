package errors

type Code int

const (
	UnexpectedCode         Code = 0
	NotFoundCode           Code = 1
	InvalidInputCode       Code = 2
	InvalidInputSchemaCode Code = 3
)

func filterCode(code Code) Code {
	switch code {
	case UnexpectedCode,
		NotFoundCode,
		InvalidInputCode,
		InvalidInputSchemaCode:
		return code

	default:
		return UnexpectedCode
	}
}
