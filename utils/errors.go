package utils

type CustomError struct {
	Message    string
	StatusCode int
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(message string, statusCode int) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: statusCode,
	}
}
