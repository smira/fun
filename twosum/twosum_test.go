package twosum_test

import (
	"strconv"
	"testing"

	"golang.org/x/exp/rand"

	"github.com/smira/fun/twosum"
)

func generateTestCase(N int) (input twosum.Input, targets []int) {
	input = make([]int, N)

	for i := 0; i < N; i++ {
		input[i] = i + 1
	}

	rand.Shuffle(N, func(i, j int) {
		input[i], input[j] = input[j], input[i]
	})

	targets = make([]int, 0, 2*N)

	for i := 3; i <= N+N-1; i++ {
		targets = append(targets, i)
	}

	return input, targets
}

func TestSolvers(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name   string
		solver twosum.Solver
	}{
		{"BruteForceSolver", twosum.BruteForceSolver},
		{"SortedSolver", twosum.SortedSolver},
		{"HashSolver", twosum.HashSolver},
		{"SplitSolver", twosum.SplitSolver},
	} {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			input, targets := generateTestCase(100)

			for _, target := range targets {
				a, b := test.solver(input, target)
				if a+b != target {
					t.Errorf("Expected %d + %d = %d, got %d + %d = %d", a, b, target, a, b, a+b)
				}
			}
		})
	}
}

func BenchmarkSolvers(b *testing.B) {
	for _, N := range []int{10, 1000, 10000, 100000} {
		b.Run("N="+strconv.Itoa(N), func(b *testing.B) {
			input, targets := generateTestCase(N)

			for _, test := range []struct {
				name   string
				solver twosum.Solver
			}{
				{"BruteForceSolver", twosum.BruteForceSolver},
				{"SortedSolver", twosum.SortedSolver},
				{"HashSolver", twosum.HashSolver},
				{"SplitSolver", twosum.SplitSolver},
			} {
				test := test

				if N >= 1000000 && test.name == "BruteForceSolver" {
					continue
				}

				b.Run(test.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						test.solver(input, targets[i%len(targets)])
					}
				})
			}
		})
	}
}
