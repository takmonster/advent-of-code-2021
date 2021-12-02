package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var depths []int
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		x, _ := strconv.Atoi(sc.Text())
		depths = append(depths, x)
	}

	var count int
	for key, current := range depths {
		previousKey := key - 1
		previousPreviousKey := key - 2
		previousPreviousPreviousKey := key - 3

		if previousKey < 0 || previousPreviousKey < 0 || previousPreviousPreviousKey < 0 {
			continue
		}

		previousValue := depths[previousKey]
		previousPreviousValue := depths[previousPreviousKey]
		previousPreviousPreviousValue := depths[previousPreviousPreviousKey]

		previousSum := previousPreviousPreviousValue + previousPreviousValue + previousValue
		sum := previousPreviousValue + previousValue + current

		fmt.Println(previousPreviousPreviousValue, previousPreviousValue, previousValue, current, previousSum, sum)

		if sum > previousSum {
			count = count + 1
		}
	}

	fmt.Println(count)
}
