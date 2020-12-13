package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"strconv"
)

const north = "north"
const south = "south"
const east = "east"
const west = "west"

func main() {
	elements := utils.ReadFile("/day-12/input.txt")

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
	directions := make(map[string]int)
	orientation := east
	oppositeOrientation := map[string]string {
		south: north,
		north: south,
		west: east,
		east: west,
	}
	for _, element := range elements {
		newDirection := element[0:1]
		value, _ := strconv.Atoi(element[1:])
		switch newDirection {
		case "N": increaseValue(directions, value, south, north)
		case "S": increaseValue(directions, value, north, south)
		case "E": increaseValue(directions, value, west, east)
		case "W": increaseValue(directions, value, east, west)
		case "L": orientation = updateDirection(orientation, "L", value)
		case "R": orientation = updateDirection(orientation, "R", value)
		case "F": increaseValue(directions, value, oppositeOrientation[orientation], orientation)
		}
	}
	return directions[north] + directions[south] + directions[east] + directions[west]
}

func findSolutionTwo(elements []string) int {
	waypoint := map[string]int {
		east: 10,
		north: 1,
	}
	directions := make(map[string]int)
	for _, element := range elements {
		newDirection := element[0:1]
		value, _ := strconv.Atoi(element[1:])
		switch newDirection {
		case "N": increaseValue(waypoint, value, south, north)
		case "S": increaseValue(waypoint, value, north, south)
		case "E": increaseValue(waypoint, value, west, east)
		case "W": increaseValue(waypoint, value, east, west)
		case "L": updateWaypoint(waypoint, value, "L")
		case "R": updateWaypoint(waypoint, value, "R")
		case "F": increaseValueWithWaypoint(waypoint, directions, value)
		}
	}
	return directions[north] + directions[south] + directions[east] + directions[west]
}

func increaseValueWithWaypoint(waypoint map[string]int, directions map[string]int, value int) {
	if waypoint[north] != 0 {
		increaseValue(directions, waypoint[north]*value, south, north)
	} else {
		increaseValue(directions, waypoint[south]*value, north, south)
	}
	if waypoint[east] != 0 {
		increaseValue(directions, waypoint[east]*value, west, east)
	} else {
		increaseValue(directions, waypoint[west]*value, east, west)
	}
}

func updateWaypoint(waypoint map[string]int, value int, turn string) {
	newNorth := updateDirection(north, turn, value)
	newEast := updateDirection(east, turn, value)
	newWest := updateDirection(west, turn, value)
	newSouth := updateDirection(south, turn, value)
	northValue := waypoint[north]
	eastValue := waypoint[east]
	westValue := waypoint[west]
	southValue := waypoint[south]
	waypoint[newNorth] = northValue
	waypoint[newEast] = eastValue
	waypoint[newWest] = westValue
	waypoint[newSouth] = southValue
}

func updateDirection(currentOrientation string, turn string, value int) string {
	degreeMapper := map[int]int {
		90: 1,
		180: 2,
		270: 3,
	}
	orientations := [4]string {
		north,
		east,
		south,
		west,
	}
	orientationIndex := -1
	for i, or := range orientations {
		if or == currentOrientation {
			orientationIndex = i
			break
		}
	}
	if turn == "R" {
		return orientations[(degreeMapper[value] + orientationIndex) % 4]
	} else {
		newIndex := orientationIndex - degreeMapper[value]
		if newIndex < 0 {
			newIndex += 4
		}
		return orientations[newIndex]
	}
}

func increaseValue(mapper map[string]int, value int, aOrientation string, oppositeOrientation string) {
	aOrientationValue := mapper[aOrientation]
	anotherOrientationValue := mapper[oppositeOrientation]
	if aOrientationValue > 0 {
		if aOrientationValue >= value {
			mapper[aOrientation] = aOrientationValue - value
		} else {
			delete(mapper, aOrientation)
			mapper[oppositeOrientation] = value - aOrientationValue
		}
	} else {
		mapper[oppositeOrientation] = value + anotherOrientationValue
	}
}
