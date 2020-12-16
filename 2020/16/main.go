package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func processTicket(vals []string) []int {
	res := make([]int, len(vals))
	for i, v := range vals {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func isValid(field int, rules []rule) bool {
	valid := false
	for _, rule := range rules {
		if field >= rule.min && field <= rule.max {
			valid = true
		}
	}
	return valid
}

type rule struct {
	min, max int
}

func readFile(ticketAnalyse func(ticket []int, rules []rule)) {
	scanner := bufio.NewScanner(os.Stdin)
	currSection := "validation"
	rules := []rule{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// scan line "your ticket"
			scanner.Scan()
			// scan ticket line
			scanner.Scan()
			myTicket := processTicket(strings.Split(scanner.Text(), ","))
			ticketAnalyse(myTicket, rules)
			// scan empty line
			scanner.Scan()
			// scan line "nearby tickets"
			scanner.Scan()
			currSection = "tickets"
		} else if currSection == "validation" {
			regexRule := regexp.MustCompile(`^[a-z ]+: (\d+)-(\d+) or (\d+)-(\d+)$`)
			submatches := regexRule.FindStringSubmatch(line)
			min1, _ := strconv.Atoi(submatches[1])
			max1, _ := strconv.Atoi(submatches[2])
			min2, _ := strconv.Atoi(submatches[3])
			max2, _ := strconv.Atoi(submatches[4])
			rules = append(rules, rule{min1, max1})
			rules = append(rules, rule{min2, max2})
		} else {
			ticket := processTicket(strings.Split(scanner.Text(), ","))
			ticketAnalyse(ticket, rules)
		}
	}
}

func part1() {
	errorRate := 0
	readFile(func(ticket []int, rules []rule) {
		for _, field := range ticket {
			if !isValid(field, rules) {
				errorRate += field
			}
		}
	})
	fmt.Println(errorRate)
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
