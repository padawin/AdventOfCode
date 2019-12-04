package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func analyzeValue(value int, part1 bool) bool {
	hasDouble := false
	increases := true
	doubles := []int{1}
	for i := 1; i < 100000; i *= 10 {
		j := i * 10
		charRight := (value / i) % 10
		charLeft := (value / j) % 10
		if charRight == charLeft {
			hasDouble = true
			doubles[len(doubles)-1] += 1
		} else if charLeft > charRight {
			increases = false
		} else {
			doubles = append(doubles, 1)
		}
	}
	if hasDouble && increases {
		if part1 {
			return true
		}
		for _, i := range doubles {
			if i == 2 {
				return true
			}
		}
	}
	return false
}

func part1() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	tmp := strings.Split(string(line), "-")
	min, _ := strconv.Atoi(tmp[0])
	max, _ := strconv.Atoi(tmp[1])
	curr := min
	countCandidates := 0
	for curr <= max {
		res := analyzeValue(curr, true)
		if res {
			countCandidates += 1
		}
		curr += 1
	}
	fmt.Println(countCandidates)
}

func part2() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	tmp := strings.Split(string(line), "-")
	min, _ := strconv.Atoi(tmp[0])
	max, _ := strconv.Atoi(tmp[1])
	curr := min
	countCandidates := 0
	for curr <= max {
		res := analyzeValue(curr, false)
		if res {
			countCandidates += 1
		}
		curr += 1
	}
	fmt.Println(countCandidates)
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
