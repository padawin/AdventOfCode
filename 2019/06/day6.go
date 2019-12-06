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

func findPathToRoot(objects map[string]string, obj string) []string {
	var res []string
	for obj != "" {
		res = append([]string{obj}, res...)
		obj = objects[obj]
	}
	return res
}

func findShortestPath(objects map[string]string, obj1 string, obj2 string) int {
	obj1ToRoot := findPathToRoot(objects, obj1)
	obj2ToRoot := findPathToRoot(objects, obj2)
	for obj1ToRoot[0] == obj2ToRoot[0] {
		obj1ToRoot = obj1ToRoot[1:]
		obj2ToRoot = obj2ToRoot[1:]
	}
	return len(obj1ToRoot) - 1 + len(obj2ToRoot) - 1
}

func part1() {
	objects := getObjects()
	res := countOrbits(objects)
	fmt.Println(res)
}

func part2() {
	objects := getObjects()
	res := findShortestPath(objects, "YOU", "SAN")
	fmt.Println(res)
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
