package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	positionX := 0
	positionY := 0
	directions := []byte{'E', 'S', 'W', 'N'}
	direction := 0
	var actions map[byte]func(int)
	actions = map[byte]func(int){
		'E': func(val int) {
			positionX += val
		},
		'S': func(val int) {
			positionY -= val
		},
		'W': func(val int) {
			positionX -= val
		},
		'N': func(val int) {
			positionY += val
		},
		'F': func(val int) {
			actions[directions[direction]](val)
		},
		'R': func(val int) {
			direction = (direction + val/90) % 4
		},
		'L': func(val int) {
			direction = (4 + direction - val/90) % 4
		},
	}
	for scanner.Scan() {
		line := scanner.Text()
		action := line[0]
		value, _ := strconv.Atoi(line[1:])
		actions[action](value)
	}
	fmt.Println(int(math.Abs(float64(positionX)) + math.Abs(float64(positionY))))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go [1|2]")
		os.Exit(1)
	}
	if os.Args[1] == "1" {
		part1()
	}
}
