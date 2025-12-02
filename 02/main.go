package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("02.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		ranges := strings.SplitSeq(strings.Trim(line, ","), ",")

		for r := range ranges {
			fromTo := strings.Split(r, "-")
			fromStr := fromTo[0]
			toStr := fromTo[1]
			fromInt, _ := strconv.Atoi(strings.TrimSpace(fromStr))
			toInt, _ := strconv.Atoi(strings.TrimSpace(toStr))

			for i := fromInt; i <= toInt; i++ {

				currStr := strconv.Itoa(i)

				maxSeqLength := len(currStr) / 2

				for j := 1; j <= maxSeqLength; j++ {
					seq := currStr[0:j]
					testStr := strings.Repeat(seq, len(currStr)/len(seq))
					if testStr == currStr {
						sum += i
						break
					}
				}

				// part 1
				// currStr := strconv.Itoa(i)
				// if len(currStr)%2 == 0 {
				// 	firstHalf := currStr[:len(currStr)/2]
				// 	otherHalf := currStr[len(currStr)/2:]
				// 	if firstHalf == otherHalf {
				// 		fmt.Println("Found:", currStr)
				// 		sum += i
				// 	}
				// }
			}
		}
	}

	fmt.Println("sum", sum)
}
