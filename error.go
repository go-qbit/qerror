package qerror

import (
	"fmt"
	"runtime"
	"time"
	"bytes"
	"strconv"
)

type QBitError interface {
	error
	Stacktrace() []CallerInfo
	Dt() time.Time
}

type qbitError struct {
	dt          time.Time
	message     string
	messageArgs []interface{}
	stacktrace  []CallerInfo
}

type CallerInfo struct {
	File     string
	Line     int
	FuncName string
}

func New(message string, a ...interface{}) QBitError {
	e := &qbitError{
		dt:          time.Now(),
		message:     message,
		messageArgs: a,
	}

	var (
		pc   uintptr
		ok   bool
		info CallerInfo
	)
	for i := 1; ; i++ {
		pc, info.File, info.Line, ok = runtime.Caller(i)
		if !ok {
			break
		}

		f := runtime.FuncForPC(pc)
		if f != nil {
			info.FuncName = f.Name()
		}

		e.stacktrace = append(e.stacktrace, info)
	}

	return e
}

func (e *qbitError) Error() string {
	buf := &bytes.Buffer{}

	fmt.Fprintf(buf, e.message, e.messageArgs...)
	buf.WriteString("\nStacktrace:")
	for _, caller := range e.stacktrace {
		buf.WriteByte('\n')
		buf.WriteByte('\t')
		buf.WriteString(caller.FuncName)
		buf.WriteString(" at ")
		buf.WriteString(caller.File)
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(caller.Line))
	}

	return buf.String()
}

func (e *qbitError) Stacktrace() []CallerInfo {
	return e.stacktrace
}

func (e *qbitError) Dt() time.Time {
	return e.dt
}
