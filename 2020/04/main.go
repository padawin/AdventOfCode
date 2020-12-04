package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var neededFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func isDocumentValid(document map[string]string) bool {
	valid := true
	for _, field := range neededFields {
		if _, found := document[field]; !found {
			valid = false
			break
		}
	}
	return valid
}

func part1() {
	currentDocument := map[string]string{}
	countValidDocs := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if isDocumentValid(currentDocument) {
				countValidDocs++
			}
			currentDocument = map[string]string{}
			continue
		}
		fields := strings.Split(line, " ")
		for _, field := range fields {
			data := strings.Split(field, ":")
			currentDocument[data[0]] = data[1]
		}
	}
	if isDocumentValid(currentDocument) {
		countValidDocs++
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println(countValidDocs)
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
