package main

import (
	"adventcode/utils"
	"strconv"
	"testing"
)

func TestExampleOne(t *testing.T) {
	elements := utils.ReadFile("/test-input.txt")
	expected := 514579

	solution := findSolutionOne(elements, 2020)

	if solution != expected {
		t.Errorf("Solution = %d; want " + strconv.Itoa(expected), solution)
	}
}

func TestExampleTwo(t *testing.T) {
	elements := utils.ReadFile("/test-input.txt")
	expected := 241861950

	solution := findSolutionTwo(elements)

	if solution != expected {
		t.Errorf("Solution = %d; want " + strconv.Itoa(expected), solution)
	}
}
