package errors

import "errors"

var (
	IncorrectParamsCount = errors.New("incorrect request parameters count")
	UnknownMethod        = errors.New("unknown method")
	UnknownType          = errors.New("unknown type")
)
