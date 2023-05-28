package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/smira/fun/queens"
	"github.com/smira/fun/twosum"
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
	case "twosum":
		fmt.Print("enter numbers: ")

		reader := bufio.NewReader(os.Stdin)

		s, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		var input twosum.Input

		for _, elem := range strings.Split(string(s), " ") {
			elem = strings.TrimSpace(elem)
			if elem == "" {
				continue
			}

			n, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatal(err)
			}

			input = append(input, n)
		}

		fmt.Print("enter target: ")

		var target int

		if _, err := fmt.Scanf("%d\n", &target); err != nil {
			log.Fatal(err)
		}

		a, b := twosum.BruteForceSolver(input, target)

		if a == 0 && b == 0 {
			fmt.Print("no solution found\n")
		} else {
			fmt.Printf("%d + %d = %d\n", a, b, target)
		}
	default:
		log.Fatal("usage: fun queens|twosum")
	}
}
