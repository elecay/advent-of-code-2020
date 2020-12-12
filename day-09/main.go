package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"math"
	"strconv"
)

func main() {
	elements := utils.ReadFile("/day-09/input.txt")

	const preamble = 25

	solutionOne := findSolutionOne(elements, preamble)
	if solutionOne != -1 {
		fmt.Print("Solution 1: ")
		fmt.Println(solutionOne)
	} else {
		fmt.Println("Solution 1 not found")
	}

	solutionTwo := findSolutionTwo(elements, findSolutionOne(elements, preamble))
	if solutionTwo != -1 {
		fmt.Print("Solution 2: ")
		fmt.Println(solutionTwo)
	} else {
		fmt.Println("Solution 2 not found")
	}
}

func findSolutionOne(elements []string, preamble int) int {
	numbers := buildInitialList(elements, preamble)
	mapper := buildMapper(numbers)
	remainingNumbers := elements[preamble:]

	for i := 0; i < len(elements); i++ {
		number, _ := strconv.Atoi(remainingNumbers[i])
		if !mapper[number] {
			return number
		}
		mapper = rebuildMapper(numbers, number, i % preamble)
	}
	return -1
}

func findSolutionTwo(elements []string, upTo int) int {
	for i := 0; i < len(elements) - 1; i++ {
		var numbers []int
		number, _ := strconv.Atoi(elements[i])
		totalSum := number
		numbers = append(numbers, number)
		for j := i + 1; j < len(elements); j++ {
			innerNumber, _ := strconv.Atoi(elements[j])
			numbers = append(numbers, innerNumber)
			totalSum += innerNumber
			if totalSum == upTo {
				smaller, bigger := findSmallerAndBigger(numbers)
				return smaller + bigger
			} else if totalSum > upTo {
				break
			}
		}
	}
	return -1
}

func findSmallerAndBigger(numbers []int) (int, int) {
	smaller := math.MaxInt64
	bigger := -1
	for _, number := range numbers {
		if number < smaller {
			smaller = number
		}
		if number > bigger {
			bigger = number
		}
	}
	return smaller, bigger
}

func buildInitialList(elements []string, preamble int) []int {
	var numbers []int
	for i := 0; i < preamble; i++ {
		number, _ := strconv.Atoi(elements[i])
		numbers = append(numbers, number)
	}
	return numbers
}

func rebuildMapper(numbers []int, number int, position int) map[int]bool {
	numbers[position] = number
	return buildMapper(numbers)
}

func buildMapper(numbers []int) map[int]bool {
	mapper := make(map[int]bool)
	for _, number := range numbers {
		for _, toAddNumber := range numbers {
			if number != toAddNumber {
				mapper[number + toAddNumber] = true
			}
		}
	}
	return mapper
}
