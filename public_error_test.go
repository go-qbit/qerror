package qerror_test

import (
	"testing"

	"github.com/go-qbit/qerror"
	"github.com/stretchr/testify/assert"
)

func TestToPublic(t *testing.T) {
	pubErr := qerror.PublicErrorf("Publict error")

	assert.Equal(t, pubErr, qerror.ToPublic(pubErr, "Public message"))
}
