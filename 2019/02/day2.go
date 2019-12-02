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

func part1() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	input := strings.Split(string(line), ",")
	var intcode []int
	for _, val := range input {
		intVal, _ := strconv.Atoi(val)
		intcode = append(intcode, intVal)
	}
	intcode[1] = 12
	intcode[2] = 2
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
	fmt.Println(intcode[0])
}

func main() {
	if len(os.Args) != 2 {
		return
	} else if os.Args[1] == "1" {
		part1()
	}
}
