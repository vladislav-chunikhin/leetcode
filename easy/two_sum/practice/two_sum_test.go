package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//You can return the answer in any order.

//Example 1:
//
//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

//Example 2:
//
//Input: nums = [3,2,4], target = 6
//Output: [1,2]

//Example 3:
//
//Input: nums = [3,3], target = 6
//Output: [0,1]

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name           string
		numbs          []int
		target         int
		expectedResult []int
	}{
		{
			name:           "Case 1",
			numbs:          []int{2, 7, 11, 15},
			target:         9,
			expectedResult: []int{0, 1},
		},
		{
			name:           "Case 2",
			numbs:          []int{3, 2, 4},
			target:         6,
			expectedResult: []int{1, 2},
		},
		{
			name:           "Case 3",
			numbs:          []int{3, 3},
			target:         6,
			expectedResult: []int{0, 1},
		},
		{
			name:           "Case 4",
			numbs:          []int{5, 2, 7, 8},
			target:         9,
			expectedResult: []int{1, 2},
		},
		{
			name:           "Case 5",
			numbs:          []int{5},
			target:         5,
			expectedResult: []int{0},
		},
		{
			name:           "Case 6",
			numbs:          []int{9, 2, 3},
			target:         5,
			expectedResult: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := twoSum(tt.numbs, tt.target)
			require.Equal(t, tt.expectedResult, actualResult)
		})
	}
}

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{0}
	}
	store := make(map[int]int)
	for index, num := range nums {
		dif := target - num
		if secondIndex, ok := store[dif]; !ok {
			store[num] = index
		} else {
			return []int{secondIndex, index}
		}
	}
	return []int{}
}
