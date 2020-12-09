package main

import (
	"advent-of-code-2020/utils"
	"strconv"
	"testing"
)

const preamble = 5

func TestExampleOne(t *testing.T) {
	elements := utils.ReadFile("/test-input.txt")
	expected := 127

	solution := findSolutionOne(elements, preamble)

	if solution != expected {
		t.Errorf("Solution = %d; want "+strconv.Itoa(expected), solution)
	}
}

func TestExampleTwo(t *testing.T) {
	elements := utils.ReadFile("/test-input.txt")
	expected := 62

	solution := findSolutionTwo(elements, findSolutionOne(elements, preamble))

	if solution != expected {
		t.Errorf("Solution = %d; want "+strconv.Itoa(expected), solution)
	}
}
