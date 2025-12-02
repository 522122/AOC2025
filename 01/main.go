package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MIN = 0
const MAX = 99

// part 1
// func turn(direction byte, pos int, value int) int {
// 	if direction == 'L' {
// 		value = value % (MAX + 1)
// 		pos -= value
// 		if pos < MIN {
// 			pos += (MAX + 1)
// 		}
// 	} else if direction == 'R' {
// 		value = value % (MAX + 1)
// 		pos += value
// 		if pos > MAX {
// 			pos -= (MAX + 1)
// 		}
// 	}
// 	return pos
// }

func turn(direction byte, pos int, value int) (int, int) {
	clicks := 0

	if direction == 'L' {
		for range value {
			pos -= 1
			if pos == 0 {
				clicks += 1
			}
			if pos < MIN {
				pos = MAX
			}
		}
	} else if direction == 'R' {
		for range value {
			pos += 1
			if pos > MAX {
				pos = MIN
				clicks += 1
			}
		}
	}
	return pos, clicks
}

func main() {
	file, _ := os.Open("01.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	pos := 50
	password := 0
	clicks := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		value, _ := strconv.Atoi(line[1:])

		// part 1
		// pos = turn(direction, pos, value)
		// if pos == 0 {
		// 	password += 1
		// }

		pos, clicks = turn(direction, pos, value)
		password += clicks
	}
	fmt.Println("Password:", password)
}
