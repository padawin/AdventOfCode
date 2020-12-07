package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func part1() {
	regexContainerBag := regexp.MustCompile(`^([a-z]+ [a-z]+) bags contain`)
	regexContainedBags := regexp.MustCompile(`\d+ ([a-z]+ [a-z]+) bag`)
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go [1|2]")
		os.Exit(1)
	}
	if os.Args[1] == "1" {
		part1()
	}
}
