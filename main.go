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
	dat, err := ioutil.ReadFile("./sudokus/s.txt")
	check(err)

	s := Sudoku.New(dat)

	fmt.Println(reflect.TypeOf(dat))

	fmt.Print(string(dat) + "\n")

	fmt.Println(reflect.TypeOf(dat))

}
