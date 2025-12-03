package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findMax(line string) (rune, int) {
	max := rune(line[0])
	maxPos := 0
	for i, char := range line {
		if char > max {
			max = char
			maxPos = i
		}
	}

	return max, maxPos
}

func runeArrToInt(arr []rune) int {
	result := ""
	for _, r := range arr {
		result += string(r)
	}
	value, _ := strconv.Atoi(result)
	return value
}

func main() {
	file, _ := os.Open("01.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		// part 1
		// maxFirst, maxFirstPosition := findMax(line[0 : len(line)-1])
		// maxSecond, _ := findMax(line[maxFirstPosition+1:])

		runes := []rune{}

		lastIndex := 0
		missing := 12
		for missing > 0 {
			slice := line[lastIndex : len(line)-(missing-1)]
			max, index := findMax(slice)
			runes = append(runes, max)
			missing--
			lastIndex = lastIndex + index + 1
		}

		// part 1
		// numberString := string(maxFirst) + string(maxSecond)
		// number, _ := strconv.Atoi(numberString)
		// sum += number

		sum += runeArrToInt(runes)
	}
	fmt.Println("sum:", sum)
}
