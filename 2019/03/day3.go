package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type visitedCell struct {
	coord coordinate
	steps int
}

func (c coordinate) toString() string {
	return fmt.Sprintf("%d-%d", c.x, c.y)
}

func (c coordinate) distanceFromOrigin() float64 {
	return math.Abs(float64(c.x)) + math.Abs(float64(c.y))
}

func processPath(path []string, callback func(coord coordinate, steps int)) {
	currentCoordinateWire := coordinate{0, 0}
	totalSteps := 0
	for _, instruction := range path {
		direction := instruction[0]
		distance, _ := strconv.Atoi(instruction[1:])
		step := coordinate{0, 0}
		if direction == 'D' {
			step.y = -1
		} else if direction == 'U' {
			step.y = 1
		} else if direction == 'L' {
			step.x = -1
		} else if direction == 'R' {
			step.x = 1
		}
		for s := 1; s <= distance; s++ {
			totalSteps += 1
			coord := coordinate{
				currentCoordinateWire.x + s*step.x,
				currentCoordinateWire.y + s*step.y,
			}
			callback(coord, totalSteps)
		}
		currentCoordinateWire.x += distance * step.x
		currentCoordinateWire.y += distance * step.y
	}
}

func getPaths() ([]string, []string) {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	path1 := strings.Split(string(line), ",")
	line, _, _ = reader.ReadLine()
	path2 := strings.Split(string(line), ",")
	return path1, path2
}

func part1() {
	path1, path2 := getPaths()
	usedCoordinatesWire := make(map[string]bool)
	closestCoordinate := coordinate{0, 0}
	processPath(path1, func(coord coordinate, _ int) {
		key := coord.toString()
		usedCoordinatesWire[key] = true
	})
	processPath(path2, func(coord coordinate, _ int) {
		key := coord.toString()
		if usedCoordinatesWire[key] == true {
			if closestCoordinate.distanceFromOrigin() == 0 || coord.distanceFromOrigin() < closestCoordinate.distanceFromOrigin() {
				closestCoordinate = coord
			}
		}
	})
	fmt.Println(closestCoordinate.distanceFromOrigin())
}

func part2() {
	path1, path2 := getPaths()
	usedCoordinatesWire := make(map[string]visitedCell)
	minSteps := 0
	processPath(path1, func(coord coordinate, steps int) {
		key := coord.toString()
		_, found := usedCoordinatesWire[key]
		if !found {
			usedCoordinatesWire[key] = visitedCell{coord, steps}
		}
	})
	processPath(path2, func(coord coordinate, steps int) {
		key := coord.toString()
		foundCoord, found := usedCoordinatesWire[key]
		if found {
			if minSteps == 0 || steps+foundCoord.steps < minSteps {
				minSteps = steps + foundCoord.steps
			}
		}
	})
	fmt.Println(minSteps)
}

func main() {
	if len(os.Args) != 2 {
		return
	} else if os.Args[1] == "1" {
		part1()
	} else if os.Args[1] == "2" {
		part2()
	}
}
