package errors

import "fmt"

type invalidValueError[T any] struct {
	invalidValue T
	reason string
}

func (err invalidValueError[T]) Error() string {
	return fmt.Sprintf("invalid value for object of type %T: %+v.\nreason: %s", *new(T), err.invalidValue, err.reason)
}

func NewInvalidValueError[T any](invalidValue T, reason string) invalidValueError[T] {
	return invalidValueError[T]{invalidValue, reason}
}
