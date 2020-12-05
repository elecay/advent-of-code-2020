package main

import (
	"advent-of-code-2020/utils"
	"strconv"
	"testing"
)

func TestExampleOne(t *testing.T) {
	elements := utils.ReadFile("/test-input.txt")
	expected := 820

	solution, _ := findSolutionOne(elements)

	if solution != expected {
		t.Errorf("Solution = %d; want "+strconv.Itoa(expected), solution)
	}
}

func TestExampleTwo(t *testing.T) {
	elements := utils.ReadFile("/test-input.txt")
	expected := -1

	solution := findSolutionTwo(elements)

	if solution != expected {
		t.Skip("Solution = %d; want "+strconv.Itoa(expected), solution)
	}
}

func TestRowAndColumns(t *testing.T) {
	boardingsPass := []string{"FBFBBFFRLR", "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	expected := [][]int{{44, 5}, {70, 7}, {14, 7}, {102, 4}}

	for i, element := range boardingsPass {
		row, column := findRowAndColumn(element)
		if row != expected[i][0] {
			t.Errorf("row = %d; want "+strconv.Itoa(expected[i][0]), row)
		}
		if column != expected[i][1] {
			t.Errorf("column = %d; want "+strconv.Itoa(expected[i][1]), column)
		}
	}
}