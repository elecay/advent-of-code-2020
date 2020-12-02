package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"log"
	"strconv"
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

func findSolutionOne(elements []string, target int) int {
	mapper := make(map[int]bool)
	for _, element := range elements {
		elementAsInt, err := strconv.Atoi(element)
		if err != nil {
			log.Fatal(err)
		}
		diffValue := target - elementAsInt
		_, exists := mapper[diffValue]
		if exists {
			solution := diffValue * elementAsInt
			return solution
		}
		mapper[elementAsInt] = true
	}
	return -1
}

func findSolutionTwo(elements []string) int {
	const target = 2020
	for index, element := range elements {
		elementAsInt, err := strconv.Atoi(element)
		if err != nil {
			log.Fatal(err)
		}
		possibleSolution := findSolutionOne(utils.RemoveByIndex(elements, index), target-elementAsInt)
		if possibleSolution != -1 {
			return possibleSolution * elementAsInt
		}
	}
	return -1
}
