package errors

import "fmt"

type InvalidUserInputError struct {
	invalidUserInput string
}

func (err InvalidUserInputError) Error() string {
	return fmt.Sprintf("invalid user input: %s", err.invalidUserInput)
}

func NewInvalidUserInputError(invalidUserInput string) InvalidUserInputError {
	return InvalidUserInputError{invalidUserInput}
}
