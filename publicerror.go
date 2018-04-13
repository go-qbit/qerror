package qerror

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type logger interface {
	Print(v ...interface{})
}

var Logger logger = log.New(os.Stderr, "", log.LstdFlags)

type PublicError interface {
	error
	PublicError() string
}

type publicError struct {
	err           error
	publicMessage string
}

func ToPublic(err error, publicMessage string) PublicError {
	if err == nil {
		return nil
	}

	perr, ok := err.(PublicError)
	if ok {
		return perr
	}

	return &publicError{err, publicMessage}
}

func (e *publicError) Error() string {
	return e.err.Error()
}

func (e *publicError) PublicError() string {
	if e.err != nil {
		Logger.Print(e.err.Error())
	}

	return e.publicMessage
}

type TextPublicError struct {
	*BaseError
	message     string
	messageArgs []interface{}
}

func PublicErrorf(message string, a ...interface{}) *TextPublicError {
	return &TextPublicError{New(1), message, a}
}

func (e *TextPublicError) Error() string {
	buf := &bytes.Buffer{}
	buf.WriteString(fmt.Sprintf(e.message, e.messageArgs...))
	buf.WriteByte('\n')
	buf.WriteString(e.BaseError.Error())

	return buf.String()
}

func (e *TextPublicError) PublicError() string {
	return fmt.Sprintf(e.message, e.messageArgs...)
}
