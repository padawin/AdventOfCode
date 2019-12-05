package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getParams(intcode []int, operationIndex int) (int, int) {
	modeParam1 := (intcode[operationIndex] / 100) % 10
	modeParam2 := (intcode[operationIndex] / 1000) % 10
	param1 := intcode[operationIndex+1]
	if modeParam1 == 0 {
		param1 = intcode[param1]
	}
	param2 := intcode[operationIndex+2]
	if modeParam2 == 0 {
		param2 = intcode[param2]
	}
	return param1, param2
}

func getIntCode() []int {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	input := strings.Split(string(line), ",")
	var intcode []int
	for _, val := range input {
		intVal, _ := strconv.Atoi(val)
		intcode = append(intcode, intVal)
	}
	return intcode
}

func process(intcode []int, input int) []int {
	var step int
	var output []int
	for i := 0; intcode[i] != 99 && i < len(intcode); i += step {
		instruction := intcode[i]
		opcode := instruction % 100
		if opcode == 1 {
			param1, param2 := getParams(intcode, i)
			intcode[intcode[i+3]] = param1 + param2
			step = 4
		} else if opcode == 2 {
			param1, param2 := getParams(intcode, i)
			intcode[intcode[i+3]] = param1 * param2
			step = 4
		} else if opcode == 3 {
			intcode[intcode[i+1]] = input
			step = 2
		} else if opcode == 4 {
			output = append(output, intcode[intcode[i+1]])
			step = 2
		} else {
			fmt.Printf("Opcode at position %d is not valid: %d\n", i, opcode)
		}
	}
	return output
}

func part1() {
	intcode := getIntCode()
	res := process(intcode, 1)
	fmt.Println(res[len(res)-1])
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
