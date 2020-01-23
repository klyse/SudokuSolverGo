package main

import (
	"strconv"
	"strings"
)

type Sudoku struct {
	s [9][9]int
}

func (e Sudoku) Parse(byteArray []byte) Sudoku {
	var fullStr = string(byteArray)
	fullStr = strings.Replace(fullStr, " ", "", -1)
	fullStr = strings.Replace(fullStr, ".", "0", -1)
	split := strings.Split(fullStr, "\r\n")

	for i, v := range split {
		for y, r := range v {
			e.s[i][y], _ = strconv.Atoi(string(r))
		}
	}
	return e
}
