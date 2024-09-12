package validation

import(
	"ddj/utils/errorhandling/errorhandlers"
	"ddj/utils/errorhandling/errors"
)

type ObjectValidator[T any] interface {
	Validate(object T, requirements []func(T) bool) T
}

type objectValidator[T any] struct {
	errorHandler errorhandlers.ErrorHandler
}

func (validator objectValidator[T]) Validate(object T, requirements []func(T) bool) T {
	for _, requirement := range requirements {
		if !requirement(object) {
			validator.errorHandler.Handle(errors.NewInvalidValueError(object))
			object = *new(T)
		}
	}
	return object
}
