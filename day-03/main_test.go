package main

import (
	"advent-of-code-2020/utils"
	"strconv"
	"testing"
)

func TestExampleOne(t *testing.T) {
	elements := utils.ReadFile("/test-input.txt")
	expected := 7

	solution := findSolutionOne(elements)

	if solution != expected {
		t.Errorf("Solution = %d; want "+strconv.Itoa(expected), solution)
	}
}

func TestExampleTwo(t *testing.T) {
	elements := utils.ReadFile("/test-input.txt")
	expected := 336

	solution := findSolutionTwo(elements)

	if solution != expected {
		t.Errorf("Solution = %d; want "+strconv.Itoa(expected), solution)
	}
}
