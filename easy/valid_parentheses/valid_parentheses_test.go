package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Every close bracket has a corresponding open bracket of the same type.

//Example 1:
//
//Input: s = "()"
//Output: true
//Example 2:
//
//Input: s = "()[]{}"
//Output: true
//Example 3:
//
//Input: s = "(]"
//Output: false

func Test_validParentheses(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedResult bool
	}{
		{
			name:           "Case 1",
			input:          "()",
			expectedResult: true,
		},
		{
			name:           "Case 2",
			input:          "()[]{}",
			expectedResult: true,
		},
		{
			name:           "Case 3",
			input:          "(]",
			expectedResult: false,
		},
		{
			name:           "Case 4",
			input:          "",
			expectedResult: false,
		},
		{
			name:           "Case 5",
			input:          "()[](]",
			expectedResult: false,
		},
		{
			name:           "Case 7",
			input:          "(",
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validParentheses(tt.input)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}

func validParentheses(input string) bool {
	// if the string isn't of even length,
	// it can't be valid, so we can return early
	if len(input)%2 != 0 || len(input) == 0 {
		return false
	}

	// set up stack and map
	var st []rune
	open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	// loop over string
	for _, r := range input {

		// if the current character is in the open map,
		// put its closer into the stack and continue
		if closer, ok := open[r]; ok {
			st = append(st, closer)
			continue
		}

		// otherwise, we're dealing with a closer
		// check to make sure the stack isn't empty
		// and whether the top of the stack matches
		// the current character
		l := len(st) - 1
		if l < 0 || r != st[l] {
			return false
		}

		// take the last element off the stack
		st = st[:l]
	}

	// if the stack is empty, return true, otherwise false
	return len(st) == 0
}
