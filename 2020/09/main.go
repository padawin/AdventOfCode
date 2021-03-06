package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func initPossibleChoices(preamble []int) map[int]int {
	result := map[int]int{}
	for i := 0; i < len(preamble)-1; i++ {
		for j := i + 1; j < len(preamble); j++ {
			result[preamble[i]+preamble[j]]++
		}
	}
	return result
}

func removePossibleChoice(possibleChoices map[int]int, preamble []int, choice int) map[int]int {
	for i := 0; i < len(preamble); i++ {
		key := choice + preamble[i]
		possibleChoices[key]--
		if possibleChoices[key] == 0 {
			delete(possibleChoices, key)
		}
	}
	return possibleChoices
}

func addPossibleChoice(possibleChoices map[int]int, preamble []int, choice int) map[int]int {
	for i := 0; i < len(preamble); i++ {
		key := choice + preamble[i]
		possibleChoices[key]++
	}
	return possibleChoices
}

func part1(preambleSize int) {
	scanner := bufio.NewScanner(os.Stdin)
	preamble := []int{}
	possibleChoices := map[int]int{}
	linesScanned := 0
	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		linesScanned++
		if linesScanned < preambleSize {
			preamble = append(preamble, val)
		} else if linesScanned == preambleSize {
			preamble = append(preamble, val)
			possibleChoices = initPossibleChoices(preamble)
		} else {
			if _, found := possibleChoices[val]; !found {
				fmt.Println(val)
				break
			}
			itemToRemove := preamble[0]
			preamble = preamble[1:]
			possibleChoices = removePossibleChoice(possibleChoices, preamble, itemToRemove)
			preamble = append(preamble, val)
			possibleChoices = addPossibleChoice(possibleChoices, preamble, val)
		}
	}
}

func sumInts(ints []int) int {
	res := 0
	for _, val := range ints {
		res += val
	}
	return res
}

func part2(resultPart1 int) {
	contiguous := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		contiguous = append(contiguous, val)
		sum := sumInts(contiguous)
		for sumInts(contiguous) > resultPart1 {
			contiguous = contiguous[1:]
		}
		sum = sumInts(contiguous)
		if sum == resultPart1 {
			min := 0
			max := 0
			for _, v := range contiguous {
				if min == 0 || v < min {
					min = v
				}
				if v > max {
					max = v
				}
			}
			fmt.Println(min + max)
			break
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main.go [1|2]")
		os.Exit(1)
	}
	if os.Args[1] == "1" {
		if len(os.Args) != 3 {
			fmt.Println("Usage: main.go 1 <preambleSize>")
			os.Exit(1)
		}
		preambleSize, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("preambleSize must be an int")
			os.Exit(1)
		}
		part1(preambleSize)
	} else if os.Args[1] == "2" {
		if len(os.Args) != 3 {
			fmt.Println("Usage: main.go 1 <resultPart1>")
			os.Exit(1)
		}
		resultPart1, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("resultPart1 must be an int")
			os.Exit(1)
		}
		part2(resultPart1)
	}
}
