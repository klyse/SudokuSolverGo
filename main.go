package main

import (
	"fmt"
	"github.com/bradhe/stopwatch"
	"io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("./sudokus/sudoku.txt")
	if err != nil {
		panic(err)
	}

	sudoku := new(Sudoku)
	sudoku.Parse(dat)

	fmt.Println("Input:")
	fmt.Println(ToString(sudoku.s))
	fmt.Println("____________")

	stopWatch := stopwatch.Start()
	solve, err := sudoku.Solve()
	stopWatch.Stop()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Output (in %dms):\n", stopWatch.Milliseconds())
	fmt.Println(ToString(solve))
}
