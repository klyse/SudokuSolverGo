package main

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"
)

type Sudoku struct {
	s [9][9]int
}

func (e *Sudoku) Parse(byteArray []byte) {
	var fullStr = string(byteArray)
	fullStr = strings.Replace(fullStr, " ", "", -1)
	fullStr = strings.Replace(fullStr, ".", "0", -1)
	fullStr = strings.Replace(fullStr, "\r\n", "\n", -1)
	fullStr = strings.Replace(fullStr, "\r", "\n", -1)
	split := strings.Split(fullStr, "\n")

	for i, v := range split {
		for y, r := range v {
			e.s[i][y], _ = strconv.Atoi(string(r))
		}
	}
}

func (e Sudoku) Solve() ([9][9]int, error) {
	return SolveRec(e.s, 0)
}

func ToString(e [9][9]int) string {
	outString := ""

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			c := "."
			if e[x][y] != 0 {
				c = fmt.Sprintf("%d", e[x][y])
			}
			outString = fmt.Sprintf("%v%v", outString, c)
			if (y+1)%3 == 0 {
				outString = outString + string(" ")
			}
		}
		outString = outString + string("\n")
	}
	return outString
}

func GetPossibleCombinations(s [9][9]int, cnt int) []int {
	x := cnt % 9
	y := cnt / 9

	var possNr = make([]int, 0, 9)

	if s[x][y] != 0 {
		possNr = append(possNr, s[x][y])
		return possNr
	}

	var nrAr = make([]bool, 10)

	sqBeginX := x / 3 * 3
	sqBeginY := y / 3 * 3

	for i1 := sqBeginX; i1 < sqBeginX+3; i1++ {
		for i2 := sqBeginY; i2 < sqBeginY+3; i2++ {
			nrAr[s[i1][i2]] = true
		}
	}

	for i1 := 0; i1 < 9; i1++ {
		nrAr[s[x][i1]] = true
	}

	for i1 := 0; i1 < 9; i1++ {
		nrAr[s[i1][y]] = true
	}

	for i := 1; i < len(nrAr); i++ {
		if !nrAr[i] {
			possNr = append(possNr, i)
		}
	}

	return possNr
}

func SolveRec(s [9][9]int, cnt int) ([9][9]int, error) {
	comb := GetPossibleCombinations(s, cnt)
	if len(comb) <= 0 {
		return [9][9]int{}, errors.New("no combination left")
	}

	var exec [9][9]int
	copier.Copy(&exec, &s)

	x := cnt % 9
	y := cnt / 9

	var newCnt = cnt + 1

	for _, c := range comb {
		exec[x][y] = c

		if newCnt == 9*9 {
			return exec, nil
		}

		newSolved, err := SolveRec(exec, newCnt)

		if err == nil {
			return newSolved, nil
		}
	}

	if cnt == 0 {
		panic("cannot solve sudoku")
	}

	return [9][9]int{}, errors.New("out of possibilities")
}
