package errors

type Code int

const (
	CodeUnexpected         Code = 0
	CodeNotFoundCode       Code = 1
	CodeInvalidInput       Code = 2
	CodeInvalidInputSchema Code = 3
)

func filterCode(code Code) Code {
	switch code {
	case CodeUnexpected,
		CodeNotFoundCode,
		CodeInvalidInput,
		CodeInvalidInputSchema:
		return code

	default:
		return CodeUnexpected
	}
}
