package main

import (
	"bufio"
	"os"
)

type Vector struct {
	X int
	Y int
}

var directionsList = []Vector{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
	{-1, -1},
	{1, -1},
	{-1, 1},
	{1, 1},
}

func countAdjacentRolls(grid []string, v Vector) int {
	count := 0
	for _, dir := range directionsList {
		newX := v.X + dir.X
		newY := v.Y + dir.Y
		if newY >= 0 && newY < len(grid) && newX >= 0 && newX < len(grid[0]) {
			if grid[newY][newX] == '@' {
				count++
			}
		}
	}
	return count
}

func getAccessibleByForklift(grid []string) []Vector {
	withLessThan4Adjacent := []Vector{}
	width := len(grid[0])
	height := len(grid)
	for y := 0; y < height; y++ {
		for x := range width {
			if grid[y][x] != '@' {
				continue
			}
			v := Vector{X: x, Y: y}
			count := countAdjacentRolls(grid, v)

			if count < 4 {
				withLessThan4Adjacent = append(withLessThan4Adjacent, v)
			}
		}
	}
	return withLessThan4Adjacent
}

func remove(grid []string, positions []Vector) []string {
	for _, pos := range positions {
		row := []rune(grid[pos.Y])
		row[pos.X] = 'x'
		grid[pos.Y] = string(row)
	}
	return grid
}

func main() {
	file, _ := os.Open("01.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	width := 0
	height := 0

	grid := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
		if width == 0 {
			width = len(line)
		}
		height++
	}

	totalRemoved := 0
	for {
		withLessThan4Adjacent := getAccessibleByForklift(grid)
		if len(withLessThan4Adjacent) == 0 {
			break
		}
		totalRemoved += len(withLessThan4Adjacent)
		grid = remove(grid, withLessThan4Adjacent)
	}

	println("Total removed:", totalRemoved)

	// part 1
	// withLessThan4Adjacent := getAccessibleByForklift(grid)
	// fmt.Println(withLessThan4Adjacent)
}
