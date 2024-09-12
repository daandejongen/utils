package errorhandlers

type emptyErrorHandler struct {}

func (handler emptyErrorHandler) Handle(...error) {
	//can be used as mock implementation when there is no need/wish to handle errors
}

func NewEmptyErrorHandler() ErrorHandler {
	return emptyErrorHandler{}
}