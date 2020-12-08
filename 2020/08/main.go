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

func runProgram(instructions []instruction) int {
	cursor := 0
	accumulator := 0
	for {
		if instructions[cursor].executed {
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
	return accumulator
}

func part1() {
	instructions := []instruction{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		instr := readInstruction(line)
		instructions = append(instructions, instr)
	}

	accumulator := runProgram(instructions)
	fmt.Println(accumulator)
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
