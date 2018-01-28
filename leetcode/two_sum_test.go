package main

import (
	"testing"
)

type testpair struct {
	values  []int
	target  int
	indices [2]int
}

var tests = []testpair{
	{[]int{2, 7, 11, 14}, 9, [2]int{0, 1}},
	{[]int{3, 2, 4}, 6, [2]int{1, 2}},
	{[]int{0, 3, 15, 6, 19, 5}, 11, [2]int{3, 5}},
	{[]int{12, 19, 48, 29, 24, 99, 77, 89, 20}, 49, [2]int{3, 8}},
}

func TestTwoSum(t *testing.T) {
	for _, pair := range tests {
		i := twoSum(pair.values, pair.target)
		if i != pair.indices {
			t.Error(
				"For", pair.values,
				"expected", pair.indices,
				"got", i,
			)
		}
	}
}
