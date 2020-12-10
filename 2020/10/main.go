package main

import (
	"bufio"
	"fmt"
	"math"
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

func calculateMultiplier(numberConsecutive float64) float64 {
	var multiplier float64
	multiplier = 1
	if numberConsecutive > 2 {
		multiplier = math.Pow(2, numberConsecutive-2)
		if numberConsecutive >= 5 {
			multiplier -= math.Pow(2, numberConsecutive-5+1) - 1
		}
	}
	return multiplier
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	adapters := []int{0}
	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		adapters = append(adapters, val)
	}
	sort.Ints(adapters)
	previousVal := 0
	var numberConsecutive float64
	combinations := 1.0
	for _, adapter := range adapters {
		if previousVal+1 == adapter {
			numberConsecutive++
		} else {
			combinations *= calculateMultiplier(numberConsecutive)
			numberConsecutive = 1
		}
		previousVal = adapter
	}
	combinations *= calculateMultiplier(numberConsecutive)
	fmt.Println(combinations)
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
