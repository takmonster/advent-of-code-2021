package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var moves []string
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		moves = append(moves, sc.Text())
	}

	var px, py int
	for _, move := range moves {
		mvsl := strings.Fields(move)
		action := mvsl[0]
		amount, _ := strconv.Atoi(mvsl[1])

		switch action {
		case "forward":
			py = py + amount
		case "down": // down is plus!
			px = px + amount
		case "up": // up is minus!
			px = px - amount
		}
	}

	fmt.Println(px, py, px*py)
}
