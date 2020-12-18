package main

import (
	"bufio"
	"fmt"
	"os"
)

func countNeighbourInDirection(grid map[int]map[int]map[int]byte, x, y, z, dirX, dirY, dirZ int) int {
	x += dirX
	y += dirY
	z += dirZ
	if val, found := grid[z][y][x]; !found || val == '.' {
		return 0
	} else {
		return 1
	}
}

func part1() {
	grid := map[int]map[int]map[int]byte{}
	grid[0] = map[int]map[int]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	minZ := 0
	maxZ := 0
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid[0][j] = map[int]byte{}
		for i := 0; i < len(line); i++ {
			grid[0][j][i] = line[i]
		}
		maxX = len(line) - 1
		j++
	}
	maxY = j - 1
	countOccupied := 0

	nextGrid := map[int]map[int]map[int]byte{}
	for n := 0; n < 6; n++ {
		countOccupied = 0
		nextGrid = map[int]map[int]map[int]byte{}
		for k := minZ - 1; k <= maxZ+1; k++ {
			nextGrid[k] = map[int]map[int]byte{}
			for j := minY - 1; j <= maxY+1; j++ {
				nextGrid[k][j] = map[int]byte{}
				for i := minX - 1; i <= maxX+1; i++ {
					countNeighbours := 0
					for z := -1; z <= 1; z++ {
						for y := -1; y <= 1; y++ {
							for x := -1; x <= 1; x++ {
								if !(x == 0 && y == 0 && z == 0) {
									countNeighbours += countNeighbourInDirection(grid, i, j, k, x, y, z)
								}
							}
						}
					}
					var r byte
					var found bool

					if r, found = grid[k][j][i]; !found {
						r = '.'
					}
					if r == '#' && countNeighbours != 2 && countNeighbours != 3 {
						r = '.'
					} else if r == '.' && countNeighbours == 3 {
						r = '#'
					}
					nextGrid[k][j][i] = r
					if nextGrid[k][j][i] == '#' {
						countOccupied++
					}
				}
			}
		}
		minX--
		minY--
		minZ--
		maxX++
		maxY++
		maxZ++
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
		part1()
	}
}
