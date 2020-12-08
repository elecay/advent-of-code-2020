package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strconv"
	"strings"
)

const jmp = "jmp"
const nop = "nop"
const acc = "acc"

func main() {
	elements := utils.ReadFile("/day-8/input.txt")

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

func findSolutionOne(elements []string) (int, bool) {
	opsMapper := parseData(elements, -1)
	return getResult(opsMapper)
}

func findSolutionTwo(elements []string) int {
	change := 1
	for {
		opsMapper := parseData(elements, change)
		accumulator, withLoop := getResult(opsMapper)
		if !withLoop {
			return accumulator
		}
		change++
	}
}

func getResult(opsMapper map[int]map[string]string) (int, bool) {
	executionTracker := make([]bool, len(opsMapper))
	accumulator := 0
	index := 0
	withLoop := false

	for {
		if index >= len(executionTracker) {
			break
		}
		if executionTracker[index] {
			withLoop = true
			break
		}
		executionTracker[index] = true
		accumulator, index = calculate(opsMapper[index], accumulator, index)
	}

	return accumulator, withLoop
}

func calculate(mapper map[string]string, accumulator int, position int) (int, int) {
	accValue := mapper[acc]
	newAccumulator := accumulator
	nextPosition := position + 1

	if accValue != "" {
		op, times := getOpAndTimes(accValue)
		newAccumulator += op * times
	}
	jmpValue := mapper[jmp]
	if jmpValue != "" {
		op, times := getOpAndTimes(jmpValue)
		nextPosition = position + (op * times)
	}
	return newAccumulator, nextPosition
}

func getOpAndTimes(operation string) (int, int) {
	sign := operation[0:1]
	op := 1
	if sign == "-" {
		op *= -1
	}
	times, _ := strconv.Atoi(operation[1:])
	return op, times
}

func parseData(elements []string, changeIndex int) map[int]map[string]string {
	mapper := make(map[int]map[string]string)
	currentIndex := 0
	operation := ""

	for i, element := range elements {
		values := strings.Split(element, " ")
		operation, currentIndex = getOperation(values[0], currentIndex, changeIndex)
		mapper[i] = map[string]string{operation: values[1]}
	}
	return mapper
}

func getOperation(operation string, currentIndex int, changeIndex int) (string, int) {
	if operation == jmp {
		currentIndex++
		if changeIndex == currentIndex {
			return nop, 0
		}
	} else if operation == nop {
		currentIndex++
		if changeIndex == currentIndex {
			return jmp, 0
		}
	}
	return operation, currentIndex + 1
}
