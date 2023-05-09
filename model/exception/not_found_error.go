package exception

type NotFoundError struct {
	Error string
}

func NewNotFound(error string) NotFoundError {
	return NotFoundError{Error: error}
}
