package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type rule struct {
	id       int
	raw      string
	computed string
	val1     int
	val2     int
	val3     int
	val4     int
}

func (r *rule) compute(rules map[int]*rule) {
	if r.computed != "" {
		return
	}
	val := ""
	val = rules[r.val1].computed
	if r.val2 != -1 {
		val = fmt.Sprintf("%s%s", val, rules[r.val2].computed)
	}
	if r.val3 != -1 {
		valOr := rules[r.val3].computed
		if r.val4 != -1 {
			valOr = fmt.Sprintf("%s%s", valOr, rules[r.val4].computed)
		}
		val = fmt.Sprintf("(?:%s|%s)", val, valOr)
	}
	r.computed = val
}

func computeRules(computedRules map[int]struct{}, rules map[int]*rule) {
	for {
		countComputed := 0
		for _, r := range rules {
			if _, found := computedRules[r.id]; found {
				countComputed++
				continue
			}
			_, foundVal1 := computedRules[r.val1]
			_, foundVal2 := computedRules[r.val2]
			_, foundVal3 := computedRules[r.val3]
			_, foundVal4 := computedRules[r.val4]
			if (r.val1 == -1 || foundVal1) && (r.val2 == -1 || foundVal2) && (r.val3 == -1 || foundVal3) && (r.val4 == -1 || foundVal4) {
				r.compute(rules)
				computedRules[r.id] = struct{}{}
				countComputed++
			}
		}
		if countComputed == len(rules) {
			return
		}
	}
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	res := 0
	rules := map[int]*rule{}
	current := "rules"
	var regexRuleZero *regexp.Regexp
	computedRules := map[int]struct{}{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			computeRules(computedRules, rules)
			r := fmt.Sprintf("^%s$", rules[0].computed)
			regexRuleZero = regexp.MustCompile(r)
			current = "messages"
		} else if current == "rules" {
			regexRule := regexp.MustCompile(`^(\d+): (?:"([a-z])"|(\d+)(?: (\d+))?(?: \| (\d+)(?: (\d+))?)?)`)
			submatches := regexRule.FindStringSubmatch(line)
			r := rule{}
			r.id, _ = strconv.Atoi(submatches[1])
			if submatches[2] != "" {
				r.computed = submatches[2]
				computedRules[r.id] = struct{}{}
			} else {
				var err error
				r.val1, _ = strconv.Atoi(submatches[3])
				r.val2, err = strconv.Atoi(submatches[4])
				if err != nil {
					r.val2 = -1
				}
				r.val3, err = strconv.Atoi(submatches[5])
				if err != nil {
					r.val3 = -1
				}
				r.val4, err = strconv.Atoi(submatches[6])
				if err != nil {
					r.val4 = -1
				}
			}
			rules[r.id] = &r
		} else if current == "messages" {
			if regexRuleZero.Match([]byte(line)) {
				res++
			}
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
	}
}
