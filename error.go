package qerror

import (
	"bytes"
	"runtime"
	"strconv"
	"time"
)

type BaseError struct {
	dt         time.Time
	stacktrace []CallerInfo
}

type CallerInfo struct {
	File     string
	Line     int
	FuncName string
}

func New(offset int) *BaseError {
	e := &BaseError{
		dt: time.Now(),
	}

	var (
		pc   uintptr
		ok   bool
		info CallerInfo
	)
	for i := 1 + offset; ; i++ {
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

func (e *BaseError) Error() string {
	buf := &bytes.Buffer{}

	buf.WriteString("DateTime: ")
	buf.WriteString(e.dt.String())
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

func (e *BaseError) Stacktrace() []CallerInfo {
	return e.stacktrace
}

func (e *BaseError) Dt() time.Time {
	return e.dt
}
