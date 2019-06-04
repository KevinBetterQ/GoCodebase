package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_main(t *testing.T) {
	a := make([]string, 0, 64)
	b := []string{}

	assert.Equal(t, a, b)
	assert.Equal(t, a, []string{})

}
