package main

import (
	"bufio"
	"fmt"
	"log"
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

	// assume inputs have uniform length
	numlen := len(bn[0])
	inputlen := len(bn)

	var gbin, ebin string
	for i := 0; i < numlen; i++ {
		var zct, oct int
		for j := 0; j < inputlen; j++ {
			switch string(bn[j][i]) {
			case "0":
				zct += 1
			case "1":
				oct += 1
			default:
				log.Fatal("unexpected input character")
			}
		}

		if zct > oct {
			gbin += "0"
			ebin += "1"
		} else {
			gbin += "1"
			ebin += "0"
		}
	}

	x, _ := strconv.ParseInt(gbin, 2, 64)
	y, _ := strconv.ParseInt(ebin, 2, 64)
	fmt.Println(x, y, x*y)
}
