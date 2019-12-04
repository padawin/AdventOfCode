package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func analyzeValue(value int) bool {
	hasDouble := false
	increases := true
	for i := 1; i < 100000; i *= 10 {
		j := i * 10
		charRight := (value / i) % 10
		charLeft := (value / j) % 10
		if charRight == charLeft {
			hasDouble = true
		} else if charLeft > charRight {
			increases = false
		}
	}
	return hasDouble && increases
}

func part1() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	tmp := strings.Split(string(line), "-")
	min, _ := strconv.Atoi(tmp[0])
	max, _ := strconv.Atoi(tmp[1])
	curr := min
	countCandidates := 0
	for curr != max {
		res := analyzeValue(curr)
		if res {
			countCandidates += 1
		}
		curr += 1
	}
	fmt.Println(countCandidates)
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
