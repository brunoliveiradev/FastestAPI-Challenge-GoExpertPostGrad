package util

type CustomError struct {
	Status  int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}
