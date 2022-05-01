package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	res, err := PostfixToInfix("4 2 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "(4-2)*3+5", res)
	}

	res, err = PostfixToInfix("2 3 4 + *")
	if assert.Nil(t, err) {
		assert.Equal(t, "2*(3+4)", res)
	}

	res, err = PostfixToInfix("2 3 - 4 + 5 6 7 * + *")
	if assert.Nil(t, err) {
		assert.Equal(t, "((2-3)+4)*(5+6*7)", res)
	}

	_, err = PostfixToInfix("")
	assert.EqualError(t, err, "empty input")

	_, err = PostfixToInfix("@ 2 + 1*")
	assert.EqualError(t, err, "invalid input")
}

func ExamplePostfixToInfix() {
	input := "4 2 + 3 * 5 +"
	res, _ := PostfixToInfix(input)

	fmt.Println("Usage example of PostfixToInfix()")
	fmt.Println(input, " -> ", res)
}
