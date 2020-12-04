package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

func isDocumentValid(document map[string]string, validateData bool) bool {
	valid := true
	for _, field := range neededFields {
		if _, found := document[field]; !found {
			valid = false
			break
		}
	}

	if valid && validateData {
		// byr (Birth Year) - four digits; at least 1920 and at most 2002.
		byr, err := strconv.Atoi(document["byr"])
		if err != nil || byr < 1920 || byr > 2002 {
			valid = false
		}
		// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		iyr, err := strconv.Atoi(document["iyr"])
		if err != nil || iyr < 2010 || iyr > 2020 {
			valid = false
		}
		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		eyr, err := strconv.Atoi(document["eyr"])
		if err != nil || eyr < 2020 || eyr > 2030 {
			valid = false
		}
		/*
			hgt (Height) - a number followed by either cm or in:
				If cm, the number must be at least 150 and at most 193.
				If in, the number must be at least 59 and at most 76.
		*/
		regexHeight := regexp.MustCompile(`^(\d{2,3})(cm|in)$`)
		submatches := regexHeight.FindStringSubmatch(document["hgt"])
		if submatches == nil {
			valid = false
		} else {
			value, _ := strconv.Atoi(submatches[1])
			unit := submatches[2]
			if unit == "cm" && (value < 150 || value > 193) || unit == "in" && (value < 59 || value > 76) {
				valid = false
			}
		}
		// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		regexHairColor := regexp.MustCompile(`^#[a-f0-9]{6}$`)
		if regexHairColor.FindStringSubmatch(document["hcl"]) == nil {
			valid = false
		}
		// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		regexEyeColor := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
		if regexEyeColor.FindStringSubmatch(document["ecl"]) == nil {
			valid = false
		}
		// pid (Passport ID) - a nine-digit number, including leading zeroes.
		regexPassportID := regexp.MustCompile(`^[0-9]{9}$`)
		if regexPassportID.FindStringSubmatch(document["pid"]) == nil {
			valid = false
		}
	}
	return valid
}

func part(validateData bool) {
	currentDocument := map[string]string{}
	countValidDocs := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if isDocumentValid(currentDocument, validateData) {
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
	if isDocumentValid(currentDocument, validateData) {
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
		part(false)
	} else {
		part(true)
	}
}
