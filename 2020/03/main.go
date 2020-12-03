package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1() {
	countTrees := 0
	angleX := 3
	var trackWidth int
	posX := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if trackWidth == 0 {
			trackWidth = len(line)
		}
		if line[posX] == '#' {
			countTrees += 1
		}
		posX = (posX + angleX) % trackWidth
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println(countTrees)
}

type vector struct {
	x, y int
}

func part2() {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	res := 1
	angles := []vector{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, angle := range angles {
		countTrees := 0
		var trackWidth int
		pos := vector{0, 0}
		for y, line := range lines {
			if y != pos.y {
				continue
			}
			if trackWidth == 0 {
				trackWidth = len(line)
			}
			if line[pos.x] == '#' {
				countTrees += 1
			}
			pos.x = (pos.x + angle.x) % trackWidth
			pos.y = pos.y + angle.y
		}
		if countTrees > 0 {
			res *= countTrees
		}
	}
	fmt.Println(res)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go [1|2]")
		os.Exit(1)
	}
	if os.Args[1] == "1" {
		part1()
	} else {
		part2()
	}
}
