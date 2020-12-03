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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go [1|2]")
		os.Exit(1)
	}
	if os.Args[1] == "1" {
		part1()
	}
}
