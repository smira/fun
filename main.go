package main

import (
	"log"
	"os"

	"github.com/smira/fun/queens"
)

// Run me with:
// * go run github.com/smira/fun queens
// * go run github.com/smira/fun twosum

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: fun queens|twosum")
	}

	switch os.Args[1] {
	case "queens":
		solutions := queens.FindSolution(nil)

		log.Printf("Found %d solutions", len(solutions))

		for _, solution := range solutions {
			queens.PrintSolution(solution)
		}
	}
}
