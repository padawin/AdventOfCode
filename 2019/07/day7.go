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

func process(intcode []int, input []int) []int {
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
			intcode[intcode[i+1]] = input[0]
			input = input[1:]
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
	max := 0
	var sequence [5]int
	intcode := getIntCode()
	sequenceVals := []int{0, 1, 2, 3, 4}
	for i, v1 := range sequenceVals {
		cp := make([]int, len(sequenceVals))
		copy(cp, sequenceVals)
		t1 := append(cp[:i], sequenceVals[i+1:]...)
		for i, v2 := range t1 {
			cp := make([]int, len(t1))
			copy(cp, t1)
			t2 := append(cp[:i], t1[i+1:]...)
			for i, v3 := range t2 {
				cp := make([]int, len(t2))
				copy(cp, t2)
				t3 := append(cp[:i], t2[i+1:]...)
				for i, v4 := range t3 {
					cp := make([]int, len(t3))
					copy(cp, t3)
					t4 := append(cp[:i], t3[i+1:]...)
					for _, v5 := range t4 {
						output := 0
						s := [5]int{v1, v2, v3, v4, v5}
						for _, input := range s {
							res := process(intcode, []int{input, output})
							output = res[len(res)-1]
						}
						if output > max {
							max = output
							sequence = s
						}
					}
				}
			}
		}
	}
	fmt.Println(max, sequence)
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
