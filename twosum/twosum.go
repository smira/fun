// Package twosum implements a solver for the two sum problem.
package twosum

import "sort"

// Input is the input for the two sum problem.
type Input []int

// Solver is the type of the function for solving the two sum problem.
//
// It takes the input and the target sum, and returns two numbers from the input
// that sum up to the target.
type Solver func(in Input, target int) (a int, b int)

// BruteForceSolver is a brute force solver for the two sum problem.
func BruteForceSolver(in Input, target int) (a int, b int) {
	for i, x := range in {
		for j, y := range in {
			if i == j {
				continue
			}

			if x+y == target {
				return x, y
			}
		}
	}

	return 0, 0
}

// SortedSolver solves the two sum problem by sorting the input and then using
// two pointers to find the solution.
func SortedSolver(in Input, target int) (a int, b int) {
	// Sort the input.
	in = append(Input(nil), in...)
	sort.Ints(in)

	// Use two pointers to find the solution.
	i := 0
	j := len(in) - 1

	for i < j {
		sum := in[i] + in[j]

		if sum == target {
			return in[i], in[j]
		}

		if sum < target {
			i++
		} else {
			j--
		}
	}

	return 0, 0
}

// HashSolver solves the two sum problem by using a hash table.
func HashSolver(in Input, target int) (a int, b int) {
	// Create a hash table.
	hash := make(map[int]int, len(in))

	// Fill the hash table.
	for _, x := range in {
		hash[x] = x
	}

	// Find the solution.
	for _, x := range in {
		y := target - x

		if _, ok := hash[y]; ok {
			return x, y
		}
	}

	return 0, 0
}

// SplitSolver solves the two sum problem by splitting the input into two
// halves, and then solving the problem for each half.
func SplitSolver(in Input, target int) (a int, b int) {
	mid := target / 2

	left, right := make(Input, 0, len(in)/2), make(Input, 0, len(in)/2)

	for _, x := range in {
		if x <= mid {
			left = append(left, x)
		} else if x < target {
			right = append(right, x)
		}
	}

	for _, a := range left {
		for _, b := range right {
			if a+b == target {
				return a, b
			}
		}
	}

	return 0, 0
}
