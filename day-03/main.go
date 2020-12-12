package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strings"
)

func main() {
	elements := utils.ReadFile("/day-03/input.txt")

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
	matrix := getMatrix(elements)
	return countTrees(matrix, 3, 1)
}

func findSolutionTwo(elements []string) int {
	matrix := getMatrix(elements)
	return countTrees(matrix, 1, 1) *
		countTrees(matrix, 3, 1) *
		countTrees(matrix, 5, 1) *
		countTrees(matrix, 7, 1) *
		countTrees(matrix, 1, 2)
}

func countTrees(matrix [][]string, right int, down int) int {
	rows := len(matrix)
	columns := len(matrix[0])

	trees := 0
	currentRow := 0
	currentColumn := 0

	for {
		currentRow += down
		if currentRow >= rows {
			break
		}
		currentColumn = (currentColumn+right)%columns
		if isTree(matrix[currentRow][currentColumn]) {
			trees++
		}
	}
	return trees
}

func isTree(char string) bool {
	return char == "#"
}

func getMatrix(elements []string) [][]string {
	size := len(elements)
	matrix := make([][]string, size)
	for i := 0; i < size; i++ {
		matrix[i] = strings.Split(elements[i], "")
	}
	return matrix
}
