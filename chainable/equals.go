package chainable

func (chainable Chainable[T]) Equals(other Chainable[T]) bool {
	if len(chainable) != len(other) {
		return false
	}
	for i := range chainable {
		if chainable[i] != other[i] {
			return false
		}
	}
	return true
}