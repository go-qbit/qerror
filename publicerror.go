package qerror

type PublicError interface {
	error
	PublicError() string
}

type publicError struct {
	err           error
	publicMessage string
}

func ToPublic(err error, publicMessage string) PublicError {
	return &publicError{err, publicMessage}
}

func (e *publicError) Error() string {
	return e.err.Error()
}

func (e *publicError) PublicError() string {
	return e.publicMessage
}
