// Package queens implements a solver for the eight queens puzzle.
package queens

// Dimension is the size of the board.
const Dimension = 8

// Solution is a (pontentially partial) solution to the eight queens puzzle.
//
// Number in each position is the column number of the queen in that row.
type Solution []int

// FindSolution finds all solutions for the eight queens puzzle starting from position in.
//
// FindSolution is a recursive function that tries to place a queen on each next row of the board.
// It takes into account the queens already placed on the board and skips positions that are
// under attack by them.
// It returns a slice of all solutions found.
func FindSolution(in Solution) []Solution {
	// If we have a complete solution, return it.
	if len(in) == Dimension {
		return []Solution{in}
	}

	// Otherwise, try to place a queen on each next row, and recursively call itself
	// to place the next queen.
	var result []Solution

loop:
	for candidate := 0; candidate < Dimension; candidate++ {
		for x, y := range in {
			// Check if the candidate is under attack by any of the queens already placed.
			if y == candidate {
				continue loop
			}

			if candidate-y == x-len(in) || candidate-y == len(in)-x {
				continue loop
			}
		}

		// before calling recursively, append the candidate to the solution, and copy the solution.
		result = append(result, FindSolution(append(append(Solution(nil), in...), candidate))...)
	}

	return result
}
