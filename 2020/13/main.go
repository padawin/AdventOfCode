package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	current_time, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	bus_ids := strings.Split(scanner.Text(), ",")
	next_bus := 0
	next_bus_time := -1
	for _, bus_id := range bus_ids {
		id, err := strconv.Atoi(bus_id)
		if err != nil {
			continue
		}
		bus_time := id - (current_time % id)
		if next_bus_time == -1 || bus_time < next_bus_time {
			next_bus_time = bus_time
			next_bus = id
		}
	}
	fmt.Println(next_bus * next_bus_time)
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
