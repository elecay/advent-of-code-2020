package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	elements := utils.ReadFile("/day-10/input.txt")

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
	adapters := getSortedAdapters(elements)
	oneJoltDiff := 0
	threeJoltDiff := 1
	lastAdapter := 0

	for _, adapter := range adapters {
		if (adapter - lastAdapter) == 1 {
			oneJoltDiff++
		} else if (adapter - lastAdapter) == 3 {
			threeJoltDiff++
		}
		lastAdapter = adapter
	}
	return oneJoltDiff * threeJoltDiff
}

func findSolutionTwo(elements []string) int {
	adapters := getSortedAdapters(elements)
	lastAdapter := 0
	previousAdapter := 0
	groupSize := 0
	counterMultiplier := 1

	for i := 0; i < len(adapters); i++ {
		adapter := adapters[i]
		if (adapter - lastAdapter) <= 3 {
			groupSize++
			previousAdapter = adapter
		} else {
			if groupSize == 2 {
				if (adapter - lastAdapter) == 4 {
					counterMultiplier *= 3
				} else if (adapter - lastAdapter) == 5 {
					counterMultiplier *= 2
				}
			} else if groupSize == 3 {
				if (adapter - lastAdapter) == 4 {
					counterMultiplier *= 7
				} else if (adapter - lastAdapter) == 5 {
					counterMultiplier *= 6
				} else if (adapter - lastAdapter) == 6 {
					counterMultiplier *= 4
				}
			}
			lastAdapter = previousAdapter
			i--
			groupSize = 0
		}
	}
	return counterMultiplier
}

func getSortedAdapters(elements []string) []int {
	var adapters []int
	for _, element := range elements {
		adapter, _ := strconv.Atoi(element)
		adapters = append(adapters, adapter)
	}
	sort.Ints(adapters)
	return adapters
}
