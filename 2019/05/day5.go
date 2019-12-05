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
		} else if opcode == 5 {
			param1, param2 := getParams(intcode, i)
			if param1 != 0 {
				step = param2 - i
			} else {
				step = 3
			}
		} else if opcode == 6 {
			param1, param2 := getParams(intcode, i)
			if param1 == 0 {
				step = param2 - i
			} else {
				step = 3
			}
		} else if opcode == 7 {
			param1, param2 := getParams(intcode, i)
			intcode[intcode[i+3]] = 0
			if param1 < param2 {
				intcode[intcode[i+3]] = 1
			}
			step = 4
		} else if opcode == 8 {
			param1, param2 := getParams(intcode, i)
			intcode[intcode[i+3]] = 0
			if param1 == param2 {
				intcode[intcode[i+3]] = 1
			}
			step = 4
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
	intcode := getIntCode()
	res := process(intcode, 5)
	fmt.Println(res[len(res)-1])
}

func main() {
	if len(os.Args) != 2 {
		return
	} else if os.Args[1] == "1" {
		part1()
	} else if os.Args[1] == "2" {
		part2()
	} else if os.Args[1] == "test" {
		type testcase struct {
			intcode []int
			input   int
			res     int
		}
		tests := []testcase{
			{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 8, 1},
			{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 5, 0},
			{[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 0, 1},
			{[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 8, 0},
			{[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 15, 0},

			{[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 8, 1},
			{[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 5, 0},
			{[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 0, 1},
			{[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 8, 0},
			{[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 15, 0},

			{[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, -10, 1},
			{[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 0, 0},
			{[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 5, 1},
			{[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 10, 1},

			{[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, -10, 1},
			{[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 0, 0},
			{[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 10, 1},
		}
		for _, test := range tests {
			fmt.Println(test)
			res := process(test.intcode, test.input)
			fmt.Println(res[len(res)-1] == test.res)
		}
	}
}
