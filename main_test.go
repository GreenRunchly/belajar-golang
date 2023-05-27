package main

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCobaAssertionSkip(test *testing.T) {
	hasil := runtime.GOOS
	if hasil != "darwin" {
		fmt.Println("Harus berjalan di mac!")
		test.Skip() // di Skip testingnya
	}
	fmt.Println("Opening Game Center!")
}

func TestCobaAssertion(test *testing.T) {
	hasil := ApaCoba("iya")
	assert.Equal(test, "iya", hasil, "harus iya")
}

func TestCobaAssertionRequire(test *testing.T) {
	hasil := ApaCoba("iya")
	require.Equal(test, "iya", hasil, "harus iya")
}

func ApaCoba(apa string) string {
	return apa
}
