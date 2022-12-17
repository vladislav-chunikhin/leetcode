package longest_common_prefix

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

//Write a function to find the longest common prefix string amongst an array of strings.
//
//If there is no common prefix, return an empty string "".

//Example 1:
//
//Input: strs = ["flower","flow","flight"]
//Output: "fl"
//
//Example 2:
//
//Input: strs = ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		expectedResult string
	}{
		{
			name:           "Case 1",
			input:          []string{"flower", "flow", "flight"},
			expectedResult: "fl",
		},
		{
			name:           "Case 2",
			input:          []string{"dog", "racecar", "car"},
			expectedResult: "",
		},
		{
			name:           "Case 3",
			input:          []string{"dog"},
			expectedResult: "dog",
		},
		{
			name:           "Case 4",
			input:          []string{},
			expectedResult: "",
		},
		{
			name:           "Case 5",
			input:          []string{"a"},
			expectedResult: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix(tt.input)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}

func longestCommonPrefix(input []string) string {
	if len(input) < 1 {
		return ""
	}
	if len(input) == 1 {
		return input[0]
	}

	prefix := input[0]

	for i := 1; i < len(input); i++ {
		for strings.Index(input[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return prefix
			}
		}
	}

	return prefix
}
