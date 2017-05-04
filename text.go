package qerror

import (
	"bytes"
	"fmt"
)

type TextError struct {
	*BaseError
	message     string
	messageArgs []interface{}
}

func Errorf(message string, a ...interface{}) *TextError {
	return &TextError{New(1), message, a}
}

func (e *TextError) Error() string {
	buf := &bytes.Buffer{}
	buf.WriteString(fmt.Sprintf(e.message, e.messageArgs...))
	buf.WriteByte('\n')
	buf.WriteString(e.BaseError.Error())

	return buf.String()
}
