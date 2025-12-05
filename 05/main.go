package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	From int
	To   int
}

// func isFresh(ingredient int, ranges []Range) bool {
// 	for _, r := range ranges {
// 		if ingredient >= r.From && ingredient <= r.To {
// 			return true
// 		}
// 	}
// 	return false
// }

func countUniqueItemsInRanges(ranges []Range) int {
	if len(ranges) == 0 {
		return 0
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].From < ranges[j].From
	})

	merged := []Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		current := ranges[i]
		last := &merged[len(merged)-1]

		if current.From <= last.To+1 {
			if current.To > last.To {
				last.To = current.To
			}
		} else {
			merged = append(merged, current)
		}
	}

	count := 0
	for _, r := range merged {
		count += r.To - r.From + 1
	}

	return count
}

// 371992492421325 too high
// 360629203100449 too high
// 331022713059933 too low

func main() {
	file, _ := os.Open("01.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var rangesList []Range

	// part 1
	// ranges := true
	// sumOfFresh := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {

			break
			//part 1
			// ranges = false
			// continue
		}

		// if ranges {
		r := strings.Split(line, "-")

		rFrom, _ := strconv.Atoi(r[0])
		rTo, _ := strconv.Atoi(r[1])

		rangesList = append(rangesList, Range{From: rFrom, To: rTo})
		// }

		// part 1
		// ingredient, _ := strconv.Atoi(line)

		// if isFresh(ingredient, rangesList) {
		// 	fmt.Println("ingredient", ingredient, "is fresh")
		// 	sumOfFresh += 1
		// }
	}

	fmt.Println("Sum:", countUniqueItemsInRanges(rangesList))

}
