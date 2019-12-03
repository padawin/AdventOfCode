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

func (c coordinate) toString() string {
	return fmt.Sprintf("%d-%d", c.x, c.y)
}

func (c coordinate) distanceFromOrigin() float64 {
	return math.Abs(float64(c.x)) + math.Abs(float64(c.y))
}

func processPath(path []string, callback func(coord coordinate)) {
	currentCoordinateWire := coordinate{0, 0}
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
			coord := coordinate{
				currentCoordinateWire.x + s*step.x,
				currentCoordinateWire.y + s*step.y,
			}
			callback(coord)
		}
		currentCoordinateWire.x += distance * step.x
		currentCoordinateWire.y += distance * step.y
	}
}

func part1() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	path1 := strings.Split(string(line), ",")
	line, _, _ = reader.ReadLine()
	path2 := strings.Split(string(line), ",")
	usedCoordinatesWire := make(map[string]bool)
	closestCoordinate := coordinate{0, 0}
	processPath(path1, func(coord coordinate) {
		key := coord.toString()
		usedCoordinatesWire[key] = true
	})
	processPath(path2, func(coord coordinate) {
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
