package errorhandling

type ErrorHandler interface {
	Handle(...error)
}
