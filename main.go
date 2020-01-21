package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./sudokus/sudoku.txt")
	check(err)

	fmt.Print(string(dat) + "\n")

	fmt.Println(reflect.TypeOf(dat))

	var sudoku [9][9]int

	var sudokuIdx int = 0
	for i := 0; i < len(dat); i++ {
		if string(dat[i]) == "\n" || string(dat[i]) == " " {
			continue
		}

		sudokuIdx++
		print(string(dat[i]))

		if string(dat[i]) == "." {
			sudoku[sudokuIdx/9][sudokuIdx%9] = 0
		} else{
			sudoku[sudokuIdx/9][sudokuIdx%9] = int(dat[i])
		}
	}

	fmt.Println(sudoku[0][1])
}
