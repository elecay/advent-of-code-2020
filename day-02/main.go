package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	elements := utils.ReadFile("/day-02/input.txt")

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
	validPasswordsCounter := 0
	for _, element := range elements {
		min, max, letter, letters := parseValues(element)
		letterAppearance := strings.Count(letters, letter)
		if letterAppearance >= min && letterAppearance <= max {
			validPasswordsCounter++
		}
	}
	return validPasswordsCounter
}

func findSolutionTwo(elements []string) int {
	validPasswordsCounter := 0
	for _, element := range elements {
		min, max, letter, letters := parseValues(element)
		if isValidPassword(letters, letter, min, max) {
			validPasswordsCounter++
		}
	}
	return validPasswordsCounter
}

func isValidPassword(letters string, letter string, min int, max int) bool {
	firstPosition := string(letters[min-1])
	secondPosition := string(letters[max-1])
	return (firstPosition == letter && secondPosition != letter) || (secondPosition == letter && firstPosition != letter)
}

func parseValues(element string) (int, int, string, string) {
	vars := strings.Split(element, " ")
	minMax := strings.Split(vars[0], "-")
	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		log.Fatal(err)
	}
	letter := string(vars[1][0])
	letters := vars[2]
	return min, max, letter, letters
}
