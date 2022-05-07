package lab2

import (
	"log"
	"testing"
	"strings"
	"os"
	"github.com/stretchr/testify/assert"
)

func ExampleComputing() {
	handler := ComputeHandler {
		Input: strings.NewReader("4 2 - 3 * 5 +"),
		Output: os.Stdout,
	}

	err := handler.Compute()
	if err != nil {
		log.Fatal(err)
	}

	res, err := PostfixToInfix("4 2 - 3 * 5 +")
	if err != nil {
		log.Fatal(err)
	}

	if res != "(4-2)*3+5" {
		log.Fatal(err)
	}

	// Output:
	// (4-2)*3+5
}

func TestErrorHandling(t *testing.T) {
	handler := ComputeHandler {
		Input: strings.NewReader("4 2 - zxc * 5 +"),
		Output: os.Stdout,
	}

	err := handler.Compute()
	assert.EqualError(t, err, "invalid input")
}