package errorhandlers

type ErrorHandler interface {
	Handle(...error)
}
