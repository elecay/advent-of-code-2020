package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strings"
)

func main() {
	elements := utils.ReadFile("/day-11/input.txt")

	solutionOne := findSolutionOne(elements)
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

func findSolutionOne(elements []string) int {
	matrix := buildMatrix(elements)
	hasChange := false
	for {
		hasChange, matrix = updateState(matrix, hasAdjacentFree, hasNAdjacentOccupied)
		if !hasChange {
			break
		}
	}
	return countSeats(matrix)
}

func findSolutionTwo(elements []string) int {
	matrix := buildMatrix(elements)
	hasChange := false
	for {
		hasChange, matrix = updateState(matrix, hasAdjacentFreeR, hasNAdjacentOccupiedR)
		if !hasChange {
			break
		}
	}
	return countSeats(matrix)
}

func updateState(matrix [][]string, adjacentFree func([][]string, int, int) bool,
	adjacentOccupied func(matrix [][]string, row int, column int) bool) (bool, [][]string) {

	hasChangeState := false
	matrixUpdated := copyMatrix(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if isFloor(matrix, i, j) {
				continue
			} else if isEmpty(matrix, i, j) && adjacentFree(matrix, i, j) {
				matrixUpdated[i][j] = "#"
				hasChangeState = true
			} else if isOccupied(matrix, i, j) && adjacentOccupied(matrix, i, j) {
				matrixUpdated[i][j] = "L"
				hasChangeState = true
			}
		}
	}
	return hasChangeState, matrixUpdated
}

func copyMatrix(matrix [][]string) [][]string {
	newMatrix := make([][]string, len(matrix), len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		newMatrix[i] = make([]string, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}
	return newMatrix
}

func countSeats(matrix [][]string) int {
	counter := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if isOccupied(matrix, i, j) {
				counter++
			}
		}
	}
	return counter
}

func hasAdjacentFree(matrix [][]string, row int, column int) bool {
		return isEmpty(matrix, row-1, column-1) &&
			isEmpty(matrix, row-1, column) &&
			isEmpty(matrix, row-1, column+1) &&
			isEmpty(matrix, row, column-1) &&
			isEmpty(matrix, row, column+1) &&
			isEmpty(matrix, row+1, column-1) &&
			isEmpty(matrix, row+1, column) &&
			isEmpty(matrix, row+1, column+1)
}

func hasAdjacentFreeR(matrix [][]string, row int, column int) bool {
	return recursiveCalc(matrix, row-1, column-1, -1, -1, isEmpty) &&
		recursiveCalc(matrix, row-1, column, -1, 0, isEmpty) &&
		recursiveCalc(matrix, row-1, column+1, -1, 1, isEmpty) &&
		recursiveCalc(matrix, row, column-1, 0, -1, isEmpty) &&
		recursiveCalc(matrix, row, column+1, 0, 1, isEmpty) &&
		recursiveCalc(matrix, row+1, column-1, 1, -1, isEmpty) &&
		recursiveCalc(matrix, row+1, column, 1, 0, isEmpty) &&
		recursiveCalc(matrix, row+1, column+1, 1, 1, isEmpty)
}

func recursiveCalc(matrix [][]string, row int, column int, offsetRow int, offsetColumn int,
	seatType func([][]string, int, int)bool) bool {
	if isFloor(matrix, row, column) {
		return recursiveCalc(matrix, row+offsetRow, column+offsetColumn, offsetRow, offsetColumn, seatType)
	} else {
		return seatType(matrix, row, column)
	}
}


func hasNAdjacentOccupied(matrix [][]string, row int, column int) bool {
	occupied := []bool{
		isOccupied(matrix, row-1, column-1),
		isOccupied(matrix, row-1, column),
		isOccupied(matrix, row-1, column+1),
		isOccupied(matrix, row, column-1),
		isOccupied(matrix, row, column+1),
		isOccupied(matrix, row+1, column-1),
		isOccupied(matrix, row+1, column),
		isOccupied(matrix, row+1, column+1),
	}
	counter := 0
	for _, seatOccupied := range occupied {
		if seatOccupied {
			counter++
		}
	}
	return counter >= 4
}

func hasNAdjacentOccupiedR(matrix [][]string, row int, column int) bool {
	occupied := []bool{
		recursiveCalc(matrix, row-1, column-1, -1, -1, isOccupied),
		recursiveCalc(matrix, row-1, column, -1, 0, isOccupied),
		recursiveCalc(matrix, row-1, column+1, -1, 1, isOccupied),
		recursiveCalc(matrix, row, column-1, 0, -1, isOccupied),
		recursiveCalc(matrix, row, column+1, 0, 1, isOccupied),
		recursiveCalc(matrix, row+1, column-1, 1, -1, isOccupied),
		recursiveCalc(matrix, row+1, column, 1, 0, isOccupied),
		recursiveCalc(matrix, row+1, column+1, 1, 1, isOccupied),
	}
	counter := 0
	for _, seatOccupied := range occupied {
		if seatOccupied {
			counter++
		}
	}
	return counter >= 5
}

func isOccupied(matrix [][]string, row int, column int) bool {
	if isValidPosition(matrix, row, column) {
		return matrix[row][column] == "#"
	}
	return false
}

func isEmpty(matrix [][]string, row int, column int) bool {
	if isValidPosition(matrix, row, column) {
		return matrix[row][column] == "L" || isFloor(matrix, row, column)
	}
	return true
}

func isFloor(matrix [][]string, row int, column int) bool {
	if isValidPosition(matrix, row, column) {
		return matrix[row][column] == "."
	}
	return false
}

func isValidPosition(matrix [][]string, row int, column int) bool {
	return column >= 0 && column <= (len(matrix[0])-1) && row >= 0 && row <= (len(matrix)-1)
}

func buildMatrix(elements []string) [][]string {
	matrix := make([][]string, len(elements), len(elements[0]))
	for i, element := range elements {
		matrix[i] = strings.Split(element, "")
	}
	return matrix
}
