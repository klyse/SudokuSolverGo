package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"
)

type Sudoku struct {
	s [9][9]int
}

func (e Sudoku) Parse(byteArray []byte) {
	var fullStr = string(byteArray)
	fullStr = strings.Replace(fullStr, " ", "", -1)
	fullStr = strings.Replace(fullStr, ".", "0", -1)
	split := strings.Split(fullStr, "\r\n")

	for i, v := range split {
		for y, r := range v {
			e.s[i][y], _ = strconv.Atoi(string(r))
		}
	}
}

func (e Sudoku) Solve() {
	SolveRec(e.s, 0)
}

func GetPossibleCombinations(s [9][9]int, cnt int) []int {
	x := cnt % 9
	y := cnt / 9

	var nrAr = make([]bool, 10)

	for i := 0; i < 9; i++ {
		nrAr[s[x][i]] = true
	}

	for i := 0; i < 9; i++ {
		nrAr[s[x][i]] = true
	}

	sqBeginX := x / 3 * 3
	sqBeginY := y / 3 * 3

	for i1 := sqBeginX; i1 < sqBeginX+3; i1++ {
		for i2 := sqBeginY; i2 < sqBeginY+3; i2++ {
			nrAr[s[i1][i2]] = true
		}
	}

	var possNr = make([]int, 0, 9)

	possNr = append(possNr, 0)

	for i := 1; i < len(nrAr); i++ {
		if !nrAr[i] {
			possNr = append(possNr, i)
		}
	}

	return possNr
}

func SolveRec(s [9][9]int, cnt int) {
	var exec [][]int
	copier.Copy(&s, &exec)

	comb := GetPossibleCombinations(s, cnt)
	fmt.Println(len(comb))
}
