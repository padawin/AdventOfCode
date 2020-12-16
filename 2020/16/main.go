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

func isFieldValidForRule(field int, rule rule) bool {
	return field >= rule.min1 && field <= rule.max1 || field >= rule.min2 && field <= rule.max2
}

func isValid(field int, rules []rule) bool {
	valid := false
	for _, rule := range rules {
		if isFieldValidForRule(field, rule) {
			valid = true
		}
	}
	return valid
}

type rule struct {
	name       string
	min1, max1 int
	min2, max2 int
}

func readFile(ticketAnalyse func(ticket []int, rules []rule)) ([]int, []rule) {
	scanner := bufio.NewScanner(os.Stdin)
	currSection := "validation"
	rules := []rule{}
	myTicket := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// scan line "your ticket"
			scanner.Scan()
			// scan ticket line
			scanner.Scan()
			myTicket = processTicket(strings.Split(scanner.Text(), ","))
			ticketAnalyse(myTicket, rules)
			// scan empty line
			scanner.Scan()
			// scan line "nearby tickets"
			scanner.Scan()
			currSection = "tickets"
		} else if currSection == "validation" {
			regexRule := regexp.MustCompile(`^([a-z ]+): (\d+)-(\d+) or (\d+)-(\d+)$`)
			submatches := regexRule.FindStringSubmatch(line)
			min1, _ := strconv.Atoi(submatches[2])
			max1, _ := strconv.Atoi(submatches[3])
			min2, _ := strconv.Atoi(submatches[4])
			max2, _ := strconv.Atoi(submatches[5])
			rules = append(rules, rule{submatches[1], min1, max1, min2, max2})
		} else {
			ticket := processTicket(strings.Split(scanner.Text(), ","))
			ticketAnalyse(ticket, rules)
		}
	}
	return myTicket, rules
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

func part2() {
	foundRules := map[int]int{}
	countValidRulesForFields := []int{}
	rulesForFields := []map[int]bool{}
	myTicket, rls := readFile(func(ticket []int, rules []rule) {
		for _, field := range ticket {
			if !isValid(field, rules) {
				return
			}
		}
		for i, field := range ticket {
			if len(rulesForFields) == i {
				rulesForFields = append(rulesForFields, map[int]bool{})
				countValidRulesForFields = append(countValidRulesForFields, 0)
			}
			for ridx, rule := range rules {
				if isFieldValidForRule(field, rule) {
					if _, found := rulesForFields[i][ridx]; !found {
						rulesForFields[i][ridx] = true
						countValidRulesForFields[i]++
					}
				} else if rulesForFields[i][ridx] {
					rulesForFields[i][ridx] = false
					countValidRulesForFields[i]--
				}
			}
		}
	})
	hasDuplicates := true
	for hasDuplicates {
		hasDuplicates = false
		for field, rules := range rulesForFields {
			if countValidRulesForFields[field] == 1 {
				for rule, isValid := range rules {
					if isValid {
						foundRules[rule] = field
					}
				}
			} else {
				hasDuplicates = true
				for rule, isValid := range rules {
					if !isValid {
						continue
					}
					if validField, found := foundRules[rule]; found {
						if validField != field {
							rules[rule] = false
							countValidRulesForFields[field]--
						}
					}
				}
			}
		}
	}
	lDep := len("departure")
	res := 1
	for ridx, rule := range rls {
		if len(rule.name) > lDep && rule.name[:lDep] == "departure" {
			res *= myTicket[foundRules[ridx]]
		}
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
