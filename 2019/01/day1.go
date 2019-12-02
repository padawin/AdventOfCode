package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() {
	reader := bufio.NewReader(os.Stdin)
	requiredFuel := 0
	for {
		line, _, _ := reader.ReadLine()
		sLine := string(line)
		if sLine == "" {
			break
		}
		mass, _ := strconv.Atoi(sLine)
		requiredFuel += mass/3 - 2
	}
	fmt.Println(requiredFuel)
}

func part2() {
	reader := bufio.NewReader(os.Stdin)
	requiredFuel := 0
	for {
		line, _, _ := reader.ReadLine()
		sLine := string(line)
		if sLine == "" {
			break
		}
		mass, _ := strconv.Atoi(sLine)
		for {
			requiredFuelModule := mass/3 - 2
			if requiredFuelModule < 0 {
				break
			}
			requiredFuel += requiredFuelModule
			mass = requiredFuelModule
		}
	}
	fmt.Println(requiredFuel)
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
