package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	mask := ""
	memory := map[string]int{}
	regexLine := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:3] == "mas" {
			mask = line[7:]
		} else {
			submatches := regexLine.FindStringSubmatch(line)
			memVal, _ := strconv.Atoi(submatches[2])
			val := 0
			for i := 0; i < 36; i++ {
				if mask[35-i] == '1' {
					val = val | 1<<i
				} else if mask[35-i] == 'X' {
					val = val | (memVal & (1 << i))
				}
			}
			memory[submatches[1]] = val
		}
	}
	res := 0
	for _, val := range memory {
		res += val
	}
	fmt.Println(res)
}

func getMemoryAddresses(maskTemplate string, floatingBits map[int]int, baseAddress int) []int {
	addresses := []int{}
	for floatValue := 0; floatValue < 1<<len(floatingBits); floatValue++ {
		mask := maskTemplate
		val := 0
		for i := 0; i < 36; i++ {
			if mask[35-i] == 'X' {
				bitValue := (floatValue & (1 << floatingBits[i])) >> floatingBits[i]
				if bitValue == 1 {
					val |= 1 << i
				}
			} else if mask[35-i] == '1' {
				val |= 1 << i
			} else if mask[35-i] == '0' {
				val |= baseAddress & (1 << i)
			}
		}
		addresses = append(addresses, val)
	}
	return addresses
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	mask := ""
	floatingBits := map[int]int{}
	memory := map[int]int{}
	regexLine := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:3] == "mas" {
			mask = line[7:]
			floatingBits = map[int]int{}
			for i := 0; i < 36; i++ {
				if mask[35-i] == 'X' {
					floatingBits[i] = len(floatingBits)
				}
			}
		} else {
			submatches := regexLine.FindStringSubmatch(line)
			memAddr, _ := strconv.Atoi(submatches[1])
			memVal, _ := strconv.Atoi(submatches[2])
			for _, memoryAddress := range getMemoryAddresses(mask, floatingBits, memAddr) {
				memory[memoryAddress] = memVal
			}
		}
	}
	res := 0
	for _, val := range memory {
		res += val
	}
	fmt.Println(res)
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
