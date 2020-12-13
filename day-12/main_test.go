package main

import (
	"advent-of-code-2020/utils"
	"strconv"
	"testing"
)

func TestExampleOne(t *testing.T) {
	elements := utils.ReadFile("/test-input.txt")
	expected := 25


	solution := findSolutionOne(elements)

	if solution != expected {
		t.Errorf("Solution = %d; want "+strconv.Itoa(expected), solution)
	}
}

func TestExampleTwo(t *testing.T) {
	elements := utils.ReadFile("/test-input.txt")
	expected := 286

	solution := findSolutionTwo(elements)

	if solution != expected {
		t.Errorf("Solution = %d; want "+strconv.Itoa(expected), solution)
	}
}

func TestUpdateDirections(t *testing.T) {
	if updateDirection(north, "R", 90) != east {
		t.Errorf("north, \"R\", 90 want east")
	}
	if updateDirection(north, "R", 180) != south {
		t.Errorf("north, \"R\", 180 want south")
	}
	if updateDirection(north, "R", 270) != west {
		t.Errorf("north, \"R\", 180 want west")
	}
	if updateDirection(north, "L", 90) != west {
		t.Errorf("north, \"L\", 90 want west")
	}
	if updateDirection(north, "L", 180) != south {
		t.Errorf("north, \"L\", 180 want south")
	}
	if updateDirection(north, "L", 270) != east {
		t.Errorf("north, \"L\", 270 want east")
	}

	if updateDirection(east, "R", 90) != south {
		t.Errorf("east, \"R\", 90 want south")
	}
	if updateDirection(east, "R", 180) != west {
		t.Errorf("east, \"R\", 180 want west")
	}
	if updateDirection(east, "R", 270) != north {
		t.Errorf("east, \"R\", 180 want north")
	}
	if updateDirection(east, "L", 90) != north {
		t.Errorf("east, \"L\", 90 want north")
	}
	if updateDirection(east, "L", 180) != west {
		t.Errorf("east, \"L\", 180 want west")
	}
	if updateDirection(east, "L", 270) != south {
		t.Errorf("east, \"L\", 270 want south")
	}

	if updateDirection(south, "R", 90) != west {
		t.Errorf("south, \"R\", 90 want west")
	}
	if updateDirection(south, "R", 180) != north {
		t.Errorf("south, \"R\", 180 want north")
	}
	if updateDirection(south, "R", 270) != east {
		t.Errorf("south, \"R\", 180 want east")
	}
	if updateDirection(south, "L", 90) != east {
		t.Errorf("south, \"L\", 90 want east")
	}
	if updateDirection(south, "L", 180) != north {
		t.Errorf("south, \"L\", 180 want north")
	}
	if updateDirection(south, "L", 270) != west {
		t.Errorf("south, \"L\", 270 want west")
	}

	if updateDirection(west, "R", 90) != north {
		t.Errorf("west, \"R\", 90 want north")
	}
	if updateDirection(west, "R", 180) != east {
		t.Errorf("west, \"R\", 180 want east")
	}
	if updateDirection(west, "R", 270) != south {
		t.Errorf("west, \"R\", 180 want south")
	}
	if updateDirection(west, "L", 90) != south {
		t.Errorf("west, \"L\", 90 want south")
	}
	if updateDirection(west, "L", 180) != east {
		t.Errorf("west, \"L\", 180 want east")
	}
	if updateDirection(west, "L", 270) != north {
		t.Errorf("west, \"L\", 270 want north")
	}
}
