package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := []rune(scanner.Text())
		l := 0
		r := len(line) - 1

		for l < len(line) {
			if unicode.IsDigit(line[l]) {
				break
			}
			l++
		}

		left := string(line[l])
		for r >= 0 {

			if unicode.IsDigit(line[r]) {
				break
			}
			r--
		}

		right := string(line[r])

		num, err := strconv.Atoi(left + right)

		if err != nil {
			log.Fatal(err)
		}

		total += num
	}

	fmt.Println(total)
	file.Close()
}

func partTwo() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := []rune(scanner.Text())

		l := 0
		left := '0'

		for l < len(line) {
			if unicode.IsDigit(line[l]) {
				left = line[l]
				break
			}
			if ok := checkNumString(l, &line, &left); ok {
				break
			}
			l++
		}

		r := len(line) - 1
		right := '0'

		for r >= 0 {
			if unicode.IsDigit(line[r]) {
				right = line[r]
				break
			}
			if ok := checkNumStringReverse(r, &line, &right); ok {
				break
			}
			r--
		}

		num, err := strconv.Atoi(string(left) + string(right))
		if err != nil {
			log.Fatal(err)
		}

		total += num
	}

	println(total)
	file.Close()
}

func checkNumString(idx int, line *[]rune, char *rune) bool {
	length := len(*line)

	if idx <= length-3 {
		switch string(*line)[idx : idx+3] {
		case "one":
			*char = '1'
			return true
		case "two":
			*char = '2'
			return true
		case "six":
			*char = '6'
			return true
		}
	}

	if idx <= length-4 {
		switch string(*line)[idx : idx+4] {
		case "four":
			*char = '4'
			return true
		case "five":
			*char = '5'
			return true
		case "nine":
			*char = '9'
			return true
		}
	}

	if idx <= length-5 {
		switch string(*line)[idx : idx+5] {
		case "three":
			*char = '3'
			return true
		case "seven":
			*char = '7'
			return true
		case "eight":
			*char = '8'
			return true
		}
	}

	return false
}

func checkNumStringReverse(idx int, line *[]rune, char *rune) bool {
	if idx > 3 {
		switch string(*line)[idx-4 : idx+1] {
		case "three":
			*char = '3'
			return true
		case "seven":
			*char = '7'
			return true
		case "eight":
			*char = '8'
			return true
		}
	}

	if idx > 2 {
		switch string(*line)[idx-3 : idx+1] {
		case "four":
			*char = '4'
			return true
		case "five":
			*char = '5'
			return true
		case "nine":
			*char = '9'
			return true
		}
	}

	if idx > 1 {
		switch string(*line)[idx-2 : idx+1] {
		case "one":
			*char = '1'
			return true
		case "two":
			*char = '2'
			return true
		case "six":
			*char = '6'
			return true
		}
	}

	return false
}
