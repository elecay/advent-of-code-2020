package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	elements := utils.ReadFile("/day-7/input.txt")

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
	bagsMapper := buildMapper(elements)
	result := make(map[string]bool)
	countBags("shiny gold", bagsMapper, result)
	return len(result) - 1
}

func findSolutionTwo(elements []string) int {
	bagsMapper := buildMapperTwo(elements)
	return countBagsTwo("shiny gold", bagsMapper)
}

func buildMapper(elements []string) map[string][]string {
	bagMapper := make(map[string][]string)
	for _, element := range elements {
		re, _ := regexp.Compile(`(\w+ (\w+)? bag(s)?)`)
		res := re.FindAllStringSubmatch(element, -1)
		var originBag string
		for i, r := range res {
			bag := strings.Split(r[0], " ")
			if i == 0 {
				originBag = bag[0] + " " + bag[1]
			} else {
				currentBag := bag[0] + " " + bag[1]
				bagMapper[currentBag] = append(bagMapper[currentBag], originBag)
			}
		}
	}
	return bagMapper
}

func buildMapperTwo(elements []string) map[string][]string {
	bagMapper := make(map[string][]string)
	for _, element := range elements {
		re, _ := regexp.Compile(`((\d )?\w+ (\w+)? bag(s)?)`)
		res := re.FindAllStringSubmatch(element, -1)
		var originBag string
		for i, r := range res {
			bag := strings.Split(r[0], " ")
			if i == 0 {
				originBag = bag[0] + " " + bag[1]
			} else {
				currentBag := bag[1] + " " + bag[2]
				bags := bag[0]
				bagsAsInt, _ := strconv.Atoi(bags)
				for i := 0; i < bagsAsInt; i++ {
					bagMapper[originBag] = append(bagMapper[originBag], currentBag)
				}
			}
		}
	}
	return bagMapper
}

func countBags(bag string, bagsMapper map[string][]string, result map[string]bool) {
	result[bag] = true
	for _, newBag := range bagsMapper[bag] {
		countBags(newBag, bagsMapper, result)
	}
}

func countBagsTwo(bag string, bagsMapper map[string][]string) int {
	total := len(bagsMapper[bag])
	for _, newBag := range bagsMapper[bag] {
		total += countBagsTwo(newBag, bagsMapper)
	}
	return total
}
