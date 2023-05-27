package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCobaAssertion(test *testing.T) {
	hasil := ApaCoba("iya")
	assert.Equal(test, "iya", hasil, "harus iya")
}

func ApaCoba(apa string) string {
	return apa
}
