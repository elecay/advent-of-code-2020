package main

import (
	"advent-of-code-2020/utils"
	"strconv"
	"testing"
)

func TestExampleOne(t *testing.T) {
	elements := utils.ReadFile("/test-input.txt")
	expected := 4

	solution := findSolutionOne(elements)

	if solution != expected {
		t.Errorf("Solution = %d; want "+strconv.Itoa(expected), solution)
	}
}

func TestExampleTwo(t *testing.T) {
	elements := utils.ReadFile("/test-input.txt")
	expected := 32

	solution := findSolutionTwo(elements)

	if solution != expected {
		t.Errorf("Solution = %d; want "+strconv.Itoa(expected), solution)
	}
}

func TestExampleTwo_2(t *testing.T) {
	elements := utils.ReadFile("/test-input-2.txt")
	expected := 126

	solution := findSolutionTwo(elements)

	if solution != expected {
		t.Errorf("Solution = %d; want "+strconv.Itoa(expected), solution)
	}
}