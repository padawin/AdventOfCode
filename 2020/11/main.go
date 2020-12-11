package main

import (
	"bufio"
	"fmt"
	"os"
)

func countPersonInDirection(grid [][]rune, x, y, dirX, dirY int, shootRay bool) int {
	for {
		x += dirX
		y += dirY
		if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
			return 0
		} else if grid[y][x] == '#' {
			return 1
		} else if grid[y][x] == 'L' {
			return 0
		}
		if !shootRay {
			return 0
		}
	}
}

func part(shootRay bool, minAcceptedNeighbours int) {
	grid := [][]rune{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	countOccupied := 0
	for {
		countOccupied = 0
		hasChanges := false
		nextGrid := make([][]rune, len(grid))
		for j := range grid {
			for i := range grid[j] {
				countNeighbours := countPersonInDirection(grid, i, j, -1, -1, shootRay)
				countNeighbours += countPersonInDirection(grid, i, j, 0, -1, shootRay)
				countNeighbours += countPersonInDirection(grid, i, j, 1, -1, shootRay)
				countNeighbours += countPersonInDirection(grid, i, j, -1, 0, shootRay)
				countNeighbours += countPersonInDirection(grid, i, j, 1, 0, shootRay)
				countNeighbours += countPersonInDirection(grid, i, j, -1, 1, shootRay)
				countNeighbours += countPersonInDirection(grid, i, j, 0, 1, shootRay)
				countNeighbours += countPersonInDirection(grid, i, j, 1, 1, shootRay)
				var r rune
				if grid[j][i] == 'L' && countNeighbours == 0 {
					r = '#'
					hasChanges = true
				} else if grid[j][i] == '#' && countNeighbours >= minAcceptedNeighbours {
					r = 'L'
					hasChanges = true
				} else {
					r = grid[j][i]
				}
				if r == '#' {
					countOccupied++
				}
				nextGrid[j] = append(nextGrid[j], r)
			}
		}
		if !hasChanges {
			break
		}
		grid = nextGrid
	}
	fmt.Println(countOccupied)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go [1|2]")
		os.Exit(1)
	}
	if os.Args[1] == "1" {
		part(false, 4)
	} else {
		part(true, 5)
	}
}
