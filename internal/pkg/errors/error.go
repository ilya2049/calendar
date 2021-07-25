package errors

import "fmt"

type genericError struct {
	message      string
	code         Code
	details      map[string]string
	wrappedError error
}

func New(err error) *genericError {
	return &genericError{
		message:      err.Error(),
		details:      make(map[string]string),
		wrappedError: err,
		code:         CodeUnexpected,
	}
}

func Wrapf(messageFormat string, args ...interface{}) *genericError {
	return New(fmt.Errorf(messageFormat, args...))
}

func (e *genericError) SetCode(code Code) *genericError {
	e.code = filterCode(code)

	return e
}

func (e *genericError) SetDetail(key, detail string) *genericError {
	if key == "" || detail == "" {
		return e
	}

	e.details[key] = detail

	return e
}

func (e *genericError) SetDetailf(key, detailFormat string, args ...interface{}) *genericError {
	return e.SetDetail(key, fmt.Sprintf(detailFormat, args...))
}

func (e *genericError) Error() string {
	return e.message
}

func (e *genericError) Details() map[string]string {
	detailsCopy := make(map[string]string, len(e.details))

	for key, detail := range e.details {
		detailsCopy[key] = detail
	}

	return detailsCopy
}

func (e *genericError) Code() Code {
	return e.code
}

func (e *genericError) Unwrap() error {
	return e.wrappedError
}
