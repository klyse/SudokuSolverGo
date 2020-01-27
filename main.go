package main

import (
	"fmt"
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
	fmt.Print(string(dat) + "\n")
	sudoku.Solve()
}
