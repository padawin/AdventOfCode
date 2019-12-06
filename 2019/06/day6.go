package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type object struct {
	name         string
	orbitsAround int
}

func getObjects() map[string]string {
	objects := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := reader.ReadLine()
		sLine := string(line)
		if sLine == "" {
			break
		}
		names := strings.Split(sLine, ")")
		centerName, objName := names[0], names[1]
		_, foundObj := objects[objName]
		_, foundCenter := objects[centerName]
		if !foundObj {
			objects[objName] = ""
		}
		if !foundCenter {
			objects[centerName] = ""
		}
		objects[objName] = centerName
	}
	return objects
}

func countOrbits(orbits map[string]string) int {
	res := 0
	for _, centerName := range orbits {
		for centerName != "" {
			centerName = orbits[centerName]
			res += 1
		}
	}
	return res
}

func part1() {
	objects := getObjects()
	res := countOrbits(objects)
	fmt.Println(res)
}

func part2() {
}

func main() {
	if len(os.Args) != 2 {
		return
	} else if os.Args[1] == "1" {
		part1()
	} else if os.Args[1] == "2" {
		part2()
	}
}
