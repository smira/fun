package queens_test

import (
	"testing"

	"github.com/smira/fun/queens"
)

func TestFindSolution(t *testing.T) {
	result := queens.FindSolution(nil)
	if len(result) != 92 {
		t.Errorf("Expected 92 solutions, got %d", len(result))
	}

	t.Logf("result = %v", result)
}

func BenchmarkFindSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queens.FindSolution(nil)
	}
}
