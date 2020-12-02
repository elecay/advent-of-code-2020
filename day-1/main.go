package main

import (
	"advent-of-code-2020/utils"
	"fmt"
)

func main() {
	elements := utils.ReadFile("/day-1/input.txt")
	const target = 2020

	solutionOne := findSolutionOne(elements, target)
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

func findSolutionOne(elements []int, target int) int {
	mapper := make(map[int]bool)
	for _, element := range elements {
		diffValue := target - element
		_, exists := mapper[diffValue]
		if exists {
			solution := diffValue * element
			return solution
		} else {
			mapper[element] = true
		}
	}
	return -1
}

func findSolutionTwo(elements []int) int {
	const target = 2020
	for index, element := range elements {
		possibleSolution := findSolutionOne(utils.RemoveByIndex(elements, index), target-element)
		if possibleSolution != -1 {
			return possibleSolution * element
		}
	}
	return -1
}
