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
	for key, depthValue := range depths {
		previousKey := key - 1
		if previousKey < 0 {
			continue
		}

		fmt.Println(depthValue, depths[previousKey])
		if depthValue > depths[previousKey] {
			count = count + 1
		}
	}

	fmt.Println(count)
}
