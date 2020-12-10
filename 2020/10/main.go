package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	adapters := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		adapters = append(adapters, val)
	}
	sort.Ints(adapters)
	reference := 0
	countDiff1 := 0
	countDiff3 := 1
	for _, adapter := range adapters {
		if adapter-reference == 1 {
			countDiff1++
		} else if adapter-reference == 3 {
			countDiff3++
		} else {
			fmt.Println("Invalid difference")
		}
		reference = adapter
	}
	fmt.Println(countDiff1 * countDiff3)
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
