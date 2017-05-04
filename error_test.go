package qerror

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	err := New(0)

	assert.Regexp(t, `^DateTime: .+?\nStacktrace:\n\tgithub.com/go-qbit/qerror.TestNew at .+?/github.com/go-qbit/qerror/error_test.go:\d+\n`+
		`\ttesting.tRunner at .+?/testing/testing.go:\d+`, err.Error())

	stacktrace := err.Stacktrace()
	assert.Len(t, stacktrace, 3)
	assert.Regexp(t, `.+/github.com/go-qbit/qerror/error_test.go$`, stacktrace[0].File)
	assert.NotZero(t, stacktrace[0].Line)
	assert.Equal(t, "github.com/go-qbit/qerror.TestNew", stacktrace[0].FuncName)

	assert.NotZero(t, err.Dt())
}

func TestErrorf(t *testing.T) {
	err := Errorf("test %s", "message")

	assert.Regexp(t, `^test message\nDateTime: .+?\nStacktrace:\n\tgithub.com/go-qbit/qerror.TestErrorf at .+?/github.com/go-qbit/qerror/error_test.go:\d+\n`+
		`\ttesting.tRunner at .+?/testing/testing.go:\d+`, err.Error())

	stacktrace := err.Stacktrace()
	assert.Len(t, stacktrace, 3)
	assert.Regexp(t, `.+/github.com/go-qbit/qerror/error_test.go$`, stacktrace[0].File)
	assert.NotZero(t, stacktrace[0].Line)
	assert.Equal(t, "github.com/go-qbit/qerror.TestErrorf", stacktrace[0].FuncName)
}
