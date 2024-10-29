package chainable

func (subjects Chainable[T]) Contains(target T) bool {
	contains := false
	for _, subj := range subjects {
		if subj == target {
			contains = true
		}
	}
	return contains
}
