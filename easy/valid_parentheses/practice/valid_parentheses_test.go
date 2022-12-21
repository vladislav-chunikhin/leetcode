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
		{
			name:           "Case 8",
			input:          "({[]})",
			expectedResult: true,
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
	if len(input)%2 != 0 || len(input) == 0 {
		return false
	}

	var closedBrackets []rune
	closedBracketByOpenMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, bracket := range input {
		if closed, ok := closedBracketByOpenMap[bracket]; ok {
			closedBrackets = append(closedBrackets, closed)
			continue
		}

		lastElement := len(closedBrackets) - 1
		if lastElement < 0 || closedBrackets[lastElement] != bracket {
			return false
		}

		closedBrackets = closedBrackets[:lastElement]
	}

	return len(closedBrackets) == 0
}
