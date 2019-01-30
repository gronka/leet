package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	expected := []int{11, 15, 2, 7}

	solution := Solve(input, target)
	if !reflect.DeepEqual(solution, expected) {
		t.Errorf("Test failed %v; got %v, expected %v",
			input,
			solution,
			expected,
		)
	}

}
