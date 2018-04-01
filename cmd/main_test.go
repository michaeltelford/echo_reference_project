package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPortColonPrefix(t *testing.T) {
	assert.Equal(t, ":8000", addPortColonPrefix("8000"))
	assert.Equal(t, ":8000", addPortColonPrefix(":8000"))
}
