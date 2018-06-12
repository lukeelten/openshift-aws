package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEncodeProjectid(t *testing.T) {
	assert := assert.New(t)

	input := "Test"
	expected := "test"
	assert.Equal(expected, EncodeProjectId(input))


	input = "Test Project"
	expected = "testproject"
	assert.Equal(expected, EncodeProjectId(input))

	input = ""
	expected = "placeholder-id"
	assert.Equal(expected, EncodeProjectId(input))

	input = "Complicated Project Name"
	expected = "complicatedprojectname"
	assert.Equal(expected, EncodeProjectId(input))
}
