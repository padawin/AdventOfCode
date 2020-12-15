package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part(endTurn int) {
	scanner := bufio.NewScanner(os.Stdin)
	currTurn := 0
	lastSpokenNumber := ""
	spokenToTurn := map[string]int{}
	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")
	for _, line := range input {
		currTurn++
		spokenToTurn[line] = currTurn
		lastSpokenNumber = line
	}
	delete(spokenToTurn, lastSpokenNumber)
	for currTurn != endTurn {
		currTurn++
		if turn, found := spokenToTurn[lastSpokenNumber]; !found {
			spokenToTurn[lastSpokenNumber] = currTurn - 1
			lastSpokenNumber = "0"
		} else {
			spokenToTurn[lastSpokenNumber] = currTurn - 1
			lastSpokenNumber = fmt.Sprintf("%d", currTurn-1-turn)
		}
	}
	fmt.Println(lastSpokenNumber)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go [1|2]")
		os.Exit(1)
	}
	if os.Args[1] == "1" {
		part(2020)
	} else {
		part(30000000)
	}
}
