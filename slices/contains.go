package slices

func Contains[T comparable](subjects []T, target T) bool {
	contains := false
	for _, subj := range subjects {
		if subj == target {
			contains = true
		}
	}
	return contains
}
