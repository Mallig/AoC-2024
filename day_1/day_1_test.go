package day_1

import (
	"testing"
)

var partOneAnswer int = 11

// Test SolvePartOne 
func TestSolvePartOne(t *testing.T) {
	
	left := []int{1, 2, 3, 3, 3, 4}
	right := []int{3, 3, 3, 4, 5, 9}

	result, err := SolvePartOne(left, right)
	if result != partOneAnswer || err != nil {
		t.Fatalf("Failed to solve part one\n	expected %d, received %d\n	error message: %e", partOneAnswer, result, err)
	}
}

// Test SolvePartOne returns error if input is not sorted
func TestSolvePartOneUnsortedInput(t *testing.T) {
	
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}
	
	result, err := SolvePartOne(left, right)
	if err == nil || result != 0 {
		t.Fatalf("Unsorted input not caught")
	}
}
