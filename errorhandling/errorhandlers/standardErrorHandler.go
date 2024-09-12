package errorhandlers

import (
	"log"
	"strings"
)

type standardErrorHandler struct{}

func NewErrorLogger() ErrorHandler {
	return &standardErrorHandler{}
}

func (handler standardErrorHandler) Handle(errs ...error) {
	log.Println(handler.combineErrors(errs...))
}

func (handler standardErrorHandler) combineErrors(errs ...error) string {
	builder := strings.Builder{}
	for _, err := range errs {
		if err != nil {
			builder.WriteString("\n")
			builder.WriteString(err.Error())
		}
	}
	return builder.String()
}
