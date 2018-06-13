package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
)

func TestExitOnError(t *testing.T) {
	assert := assert.New(t)

	err := errors.New("Test Error")
	assert.PanicsWithValue(err, func () {ExitOnError("Test", err)})

	err = nil
	assert.NotPanics(func () {ExitOnError("Test", err)})
}
