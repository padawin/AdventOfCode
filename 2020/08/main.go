package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	command  string
	argument int
	executed bool
}

func readInstruction(line string) instruction {
	values := strings.Split(line, " ")
	instr := instruction{}
	instr.command = values[0]
	instr.argument, _ = strconv.Atoi(values[1])
	return instr
}

func runProgram(instructions []instruction) (int, bool) {
	cursor := 0
	accumulator := 0
	infiniteLoop := false
	for cursor < len(instructions) {
		if instructions[cursor].executed {
			infiniteLoop = true
			break
		}
		instructions[cursor].executed = true
		if instructions[cursor].command == "acc" {
			accumulator += instructions[cursor].argument
			cursor += 1
		} else if instructions[cursor].command == "jmp" {
			cursor += instructions[cursor].argument
		} else if instructions[cursor].command == "nop" {
			cursor += 1
		}
	}
	return accumulator, infiniteLoop
}

func part1() {
	instructions := []instruction{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		instr := readInstruction(line)
		instructions = append(instructions, instr)
	}

	accumulator, _ := runProgram(instructions)
	fmt.Println(accumulator)
}

func part2() {
	instructions := []instruction{}
	scanner := bufio.NewScanner(os.Stdin)
	candidatesToChange := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		instr := readInstruction(line)
		if instr.command == "jmp" || instr.command == "nop" {
			candidatesToChange = append(candidatesToChange, len(instructions))
		}
		instructions = append(instructions, instr)
	}

	for _, indexToChange := range candidatesToChange {
		old := instructions[indexToChange].command
		if instructions[indexToChange].command == "nop" {
			instructions[indexToChange].command = "jmp"
		} else {
			instructions[indexToChange].command = "nop"
		}

		accumulator, infiniteLoop := runProgram(instructions)
		if !infiniteLoop {
			fmt.Println(accumulator)
			break
		}
		instructions[indexToChange].command = old
		for index, _ := range instructions {
			instructions[index].executed = false
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go [1|2]")
		os.Exit(1)
	}
	if os.Args[1] == "1" {
		part1()
	} else {
		part2()
	}
}
