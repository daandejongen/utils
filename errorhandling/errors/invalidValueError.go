package errors

import "fmt"

type invalidValueError[T any] struct {
	invalidValue T
}

func (err invalidValueError[T]) Error() string {
	return fmt.Sprintf("invalid value for type %T: %+v", *new(T), err.invalidValue)
}

func NewInvalidValueError[T any](invalidValue T) invalidValueError[T] {
	return invalidValueError[T]{invalidValue}
}
