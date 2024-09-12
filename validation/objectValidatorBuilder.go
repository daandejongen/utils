package validation

import "github.com/daandejongen/utils/errorhandling/errorhandlers"

type ObjectValidatorBuilder[T any] interface {
	Build() ObjectValidator[T]
	WithErrorHandler(handler errorhandlers.ErrorHandler) ObjectValidatorBuilder[T]
}

type objectValidatorBuilder[T any] struct {
	validator *objectValidator[T]
}

func NewObjectValidatorBuilder[T any]() ObjectValidatorBuilder[T] {
	builder := &objectValidatorBuilder[T]{
		&objectValidator[T]{},
	}
	return builder
}

func (builder *objectValidatorBuilder[T]) Build() ObjectValidator[T] {
	return builder.validator
}

func (builder *objectValidatorBuilder[T]) WithErrorHandler(handler errorhandlers.ErrorHandler) ObjectValidatorBuilder[T] {
	builder.validator.errorHandler = handler
	return builder
}
