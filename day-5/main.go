package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strings"
)

func main() {
	elements := utils.ReadFile("/day-5/input.txt")

	solutionOne, _ := findSolutionOne(elements)
	if solutionOne != -1 {
		fmt.Print("Solution 1: ")
		fmt.Println(solutionOne)
	} else {
		fmt.Println("Solution 1 not found")
	}

	solutionTwo := findSolutionTwo(elements)
	if solutionTwo != -1 {
		fmt.Print("Solution 2: ")
		fmt.Println(solutionTwo)
	} else {
		fmt.Println("Solution 2 not found")
	}
}

func findSolutionOne(elements []string) (int, map[int]bool) {
	maxId := -1
	ids := make(map[int]bool)

	for _, element := range elements {
		row, column := findRowAndColumn(element)
		possibleMaxId := getId(row, column)
		ids[possibleMaxId] = true

		if possibleMaxId > maxId {
			maxId = possibleMaxId
		}
	}
	return maxId, ids
}

func findSolutionTwo(elements []string) int {
	_, ids := findSolutionOne(elements)
	for i := 1; i < 127; i++ {
		for j := 0; j < 8; j++ {
			id := getId(i, j)
			isBackSeat := ids[id + 1]
			isFrontSeat := ids[id - 1]
			seatExists := ids[id]
			if !seatExists && isBackSeat && isFrontSeat {
				return id
			}
		}
	}
	return -1
}

func getId(row int, column int) int {
	return row * 8 + column
}

func findRowAndColumn(boardingPass string) (int, int) {
	maxRow := 127
	minRow := 0
	maxColumn := 7
	minColumn := 0

	row := -1
	column := -1

	stepsMapper := map[string]func() {
		"F": func() {
			maxRow -= ((maxRow-minRow)/2)+1
			row = maxRow
		},
		"B": func() {
			minRow += ((maxRow-minRow)/2)+1
			row = minRow
		},
		"L": func() {
			maxColumn -= ((maxColumn-minColumn)/2)+1
			column = maxColumn
		},
		"R": func() {
			minColumn += ((maxColumn-minColumn)/2)+1
			column = minColumn
		},
	}

	for _, step := range strings.Split(boardingPass, "") {
		stepsMapper[step]()
	}
	return row, column
}
