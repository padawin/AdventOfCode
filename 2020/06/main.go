package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	answers := map[rune]struct{}{}
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			count += len(answers)
			answers = map[rune]struct{}{}
			continue
		}
		for _, answer := range line {
			answers[answer] = struct{}{}
		}
	}
	count += len(answers)
	fmt.Println(count)
}

func part2() {
	answers := map[rune]int{}
	count := 0
	groupSize := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for _, c := range answers {
				if c == groupSize {
					count++
				}
			}
			answers = map[rune]int{}
			groupSize = 0
			continue
		}
		for _, answer := range line {
			answers[answer]++
		}
		groupSize++
	}
	for _, c := range answers {
		if c == groupSize {
			count++
		}
	}
	fmt.Println(count)
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
