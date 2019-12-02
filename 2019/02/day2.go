package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operations map[int]func(a, b int) int = map[int]func(a, b int) int{
	1: func(a, b int) int {
		return a + b
	},
	2: func(a, b int) int {
		return a * b
	},
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

func process(intcode []int, noun, verb int) int {
	intcode[1] = noun
	intcode[2] = verb
	for i := 0; intcode[i] != 99 && i < len(intcode); i += 4 {
		opcode := intcode[i]
		in1 := intcode[i+1]
		in2 := intcode[i+2]
		output := intcode[i+3]
		operation, ok := operations[opcode]
		if !ok {
			fmt.Printf("Opcode at position %d is not valid: %d\n", i, opcode)
		}
		intcode[output] = operation(intcode[in1], intcode[in2])
	}
	return intcode[0]
}

func part1() {
	intcode := getIntCode()
	res := process(intcode, 12, 2)
	fmt.Println(res)
}

func part2() {
	intcode := getIntCode()
	var currIntcode []int
	currIntcode = make([]int, len(intcode))
	copy(currIntcode, intcode)
	resInitial := process(currIntcode, 0, 0)
	copy(currIntcode, intcode)
	delta := process(currIntcode, 1, 0) - resInitial
	expectedRes := 19690720
	noun := (expectedRes - resInitial) / delta
	verb := expectedRes - (resInitial + (noun * delta))
	fmt.Println(noun*100 + verb)
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
