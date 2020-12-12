package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strings"
)

func main() {
	elements := utils.ReadFile("/day-06/input.txt")

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
	answers, _ := parseData(elements)
	total := 0
	for _, answer := range answers {
		total += len(answer)
	}
	return total
}

func findSolutionTwo(elements []string) int {
	answers, peopleCounter := parseData(elements)
	total := 0
	for i, answer := range answers {
		for _, v := range answer {
			if v == peopleCounter[i] {
				total++
			}
		}
	}
	return total
}

func parseData(elements []string) ([]map[string]int, map[int]int) {
	var answers []map[string]int
	peopleCounter := make(map[int]int)
	answerYes := make(map[string]int)
	count := 0
	group := 0

	for _, element := range elements {
		if element == "" {
			peopleCounter[group] = count
			answers = append(answers, answerYes)
			answerYes = make(map[string]int)
			count = 0
			group++
		} else {
			count++
			for _, letter := range strings.Split(element, "") {
				answerYes[letter] = answerYes[letter] + 1
			}
		}
	}
	answers = append(answers, answerYes)
	peopleCounter[group] = count
	return answers, peopleCounter
}
