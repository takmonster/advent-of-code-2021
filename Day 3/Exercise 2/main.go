package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var bn []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		bn = append(bn, sc.Text())
	}

	g := determineGeneratorRatingValue(0, bn)
	gd, _ := strconv.ParseInt(g, 2, 64)
	s := determineScrubberRatingValue(0, bn)
	sd, _ := strconv.ParseInt(s, 2, 64)

	fmt.Println(gd, sd, gd*sd)
}

func determineGeneratorRatingValue(pos int, numbers []string) string {
	if len(numbers) == 1 {
		return numbers[0]
	}

	zeros, ones := organizeNumbersAtPosition(pos, numbers)

	var v string
	if len(zeros) > len(ones) {
		v = determineGeneratorRatingValue(pos+1, zeros)
	} else if len(zeros) == len(ones) {
		v = determineGeneratorRatingValue(pos+1, ones)
	} else {
		v = determineGeneratorRatingValue(pos+1, ones)
	}

	return v
}

func determineScrubberRatingValue(pos int, numbers []string) string {
	if len(numbers) == 1 {
		return numbers[0]
	}

	zeros, ones := organizeNumbersAtPosition(pos, numbers)

	var v string
	if len(zeros) < len(ones) {
		v = determineScrubberRatingValue(pos+1, zeros)
	} else if len(zeros) == len(ones) {
		v = determineScrubberRatingValue(pos+1, zeros)
	} else {
		v = determineScrubberRatingValue(pos+1, ones)
	}

	return v
}

func organizeNumbersAtPosition(pos int, numbers []string) (zeros, ones []string) {
	for _, number := range numbers {
		switch string(number[pos]) {
		case "0":
			zeros = append(zeros, number)
		case "1":
			ones = append(ones, number)
		}
	}
	return
}
