package palindrome_number

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//Given an integer x, return true if x is a palindrome, and false otherwise.

//Example 1:
//
//Input: x = 121
//Output: true
//Explanation: 121 reads as 121 from left to right and from right to left.

//Example 2:
//
//Input: x = -121
//Output: false
//Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore, it is not a palindrome.

//Example 3:
//
//Input: x = 10
//Output: false
//Explanation: Reads 01 from right to left. Therefore, it is not a palindrome.

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name           string
		number         int
		expectedResult bool
	}{
		{
			name:           "Case 1",
			number:         121,
			expectedResult: true,
		},
		{
			name:           "Case 2",
			number:         -121,
			expectedResult: false,
		},
		{
			name:           "Case 3",
			number:         10,
			expectedResult: false,
		},
		{
			name:           "Case 4",
			number:         1221,
			expectedResult: true,
		},
		{
			name:           "Case 5",
			number:         5,
			expectedResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.number)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	origin := x
	var reversed int
	for x != 0 {
		chunk := x % 10
		reversed = reversed*10 + chunk
		x = x / 10
	}

	return origin == reversed

}
