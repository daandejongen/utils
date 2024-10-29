package chainable

func (subjects Chainable[T]) Where(condition func(T) bool) Chainable[T] {
	output := []T{}
	for _, subject := range subjects {
		if condition(subject) {
			output = append(output, subject)
		}
	}
	return output
}
