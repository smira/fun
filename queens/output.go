package queens

// PrintSolution prints a solution to the standard output.
func PrintSolution(solution Solution) {
	for _, y := range solution {
		for x := 0; x < Dimension; x++ {
			if x == y {
				print("Q")
			} else {
				print(".")
			}
		}
		println()
	}
	println()
}
