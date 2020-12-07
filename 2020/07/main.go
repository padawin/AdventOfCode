package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func processPossibleContainers(allContainers map[string]map[string]struct{}, possibleContainers map[string]struct{}, out map[string]struct{}) map[string]struct{} {
	for bag, _ := range possibleContainers {
		if _, found := out[bag]; !found {
			out[bag] = struct{}{}
			out = processPossibleContainers(allContainers, allContainers[bag], out)
		}
	}
	return out
}

var regexContainerBag = regexp.MustCompile(`^([a-z]+ [a-z]+) bags contain`)

func part1() {
	var regexContainedBags = regexp.MustCompile(`\d+ ([a-z]+ [a-z]+) bag`)
	containedPossibilities := map[string]map[string]struct{}{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		submatchContainer := regexContainerBag.FindStringSubmatch(line)
		container := submatchContainer[1]
		submatchesContained := regexContainedBags.FindAllStringSubmatch(line, -1)
		for _, submatch := range submatchesContained {
			if _, found := containedPossibilities[submatch[1]]; !found {
				containedPossibilities[submatch[1]] = map[string]struct{}{}
			}
			containedPossibilities[submatch[1]][container] = struct{}{}
		}
	}
	res := processPossibleContainers(containedPossibilities, containedPossibilities["shiny gold"], map[string]struct{}{})
	fmt.Println(len(res))
}

func processRequiredContent(prefix string, allContainers map[string]map[string]int, start string, multiplier int) int {
	count := 0
	//fmt.Printf("%s%s has to contain:\n", prefix, start)
	for bag, nb := range allContainers[start] {
		//fmt.Printf("%s%d %s (total: %d)\n", prefix, nb, bag, multiplier*nb)
		count += multiplier * nb
		count += processRequiredContent(fmt.Sprintf("\t%s", prefix), allContainers, bag, nb*multiplier)
	}
	return count
}

func part2() {
	var regexContainedBags = regexp.MustCompile(`(\d+) ([a-z]+ [a-z]+) bag`)
	containedPossibilities := map[string]map[string]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		submatchContainer := regexContainerBag.FindStringSubmatch(line)
		container := submatchContainer[1]
		submatchesContained := regexContainedBags.FindAllStringSubmatch(line, -1)
		containedPossibilities[container] = map[string]int{}
		for _, submatch := range submatchesContained {
			count, _ := strconv.Atoi(submatch[1])
			containedPossibilities[container][submatch[2]] = count
		}
	}
	res := processRequiredContent("", containedPossibilities, "shiny gold", 1)
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
