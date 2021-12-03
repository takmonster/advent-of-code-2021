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

	var pa, px, py int
	for _, move := range moves {
		mvsl := strings.Fields(move)
		action := mvsl[0]
		amount, _ := strconv.Atoi(mvsl[1])

		switch action {
		case "forward":
			py = py + amount        // increase horizontal position by amount
			px = px + (pa * amount) // increase depth by (aim multiplied by amount)
		case "down": // down is plus!
			pa = pa + amount // increases aim by amount
		case "up": // up is minus!
			pa = pa - amount // decreases aim by amount
		}
	}

	fmt.Println(pa, px, py, px*py)
}
