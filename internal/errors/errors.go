package errors

import "errors"

var (
	IncorrectParamsCount = errors.New("incorrect request parameters count")
	UnknownType          = errors.New("unknown type")
)
