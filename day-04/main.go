package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	elements := utils.ReadFile("/day-04/input.txt")

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
	return validatePassports(elements, validatorOne)
}

func findSolutionTwo(elements []string) int {
	return validatePassports(elements, validatorTwo)
}

func validatePassports(elements []string, validator func(string) bool) int {
	passports := parseData(elements)
	validPassports := 0
	for _, passport := range passports {
		if validator(passport) {
			validPassports++
		}
	}
	return validPassports
}

func validatorOne(passport string) bool {
	for field := range fieldsToCheck() {
		if !strings.Contains(passport, field) {
			return false
		}
	}
	return true
}

func validatorTwo(passport string) bool {
	if !validatorOne(passport) {
		return false
	}
	fieldsToCheck := fieldsToCheck()
	items := strings.Split(passport, " ")[1:]
	for _, item := range items {
		values := strings.Split(item, ":")
		match, _ := regexp.MatchString(fieldsToCheck[values[0]], values[1])
		if !match {
			return false
		}
	}
	return true
}

func fieldsToCheck() map[string]string {
	return map[string]string{
		"byr": "^(19[2-9]\\d|200[0-2])$",
		"iyr": "^(201\\d|2020)$",
		"eyr": "^(202\\d|2030)$",
		"hgt": "^(1[5-8]\\d|19[0-3])cm$|^(59|6[0-9]|7[0-6])in$",
		"hcl": "^#([a-f0-9]{6})$",
		"ecl": "^(amb|blu|brn|gry|grn|hzl|oth)$",
		"pid": "^\\d{9}$",
	}
}

func parseData(elements []string) []string {
	var passports []string
	var passport string
	for _, element := range elements {
		if element == "" {
			passports = append(passports, passport)
			passport = ""
		} else {
			passport += " "
			passport += element
		}
	}
	return append(passports, passport)
}
