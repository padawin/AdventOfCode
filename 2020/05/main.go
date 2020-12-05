package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func bisect(directions string, count int) int {
	min := 0
	max := count
	iter := int(math.Log2(float64(count)))
	for i := 0; i < iter; i++ {
		var middle int
		middle = min + (max-min)/2
		if directions[i] == 'F' || directions[i] == 'L' {
			max = middle
		} else if directions[i] == 'B' || directions[i] == 'R' {
			min = middle
		}
	}
	return min
}

func calculateSeat(code string) (int, int, int) {
	row := bisect(code, 128)
	seat := bisect(code[7:], 8)
	id := row*8 + seat
	return row, seat, id
}

func part1() {
	maxID := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		_, _, id := calculateSeat(line)
		if id > maxID {
			maxID = id
		}
	}
	fmt.Println(maxID)
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
