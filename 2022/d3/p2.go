package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func getCharCode(letter rune) int {
	if unicode.IsUpper(letter) {
		return int(letter) - 38
	}
	return int(letter) - 96
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Error reading file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([][]rune, 0)

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error scanning file", err)
	}

	total := 0

	for scanner.Scan() {
		line := []rune(scanner.Text())
		lines = append(lines, line)
	}

	for i := 0; i < len(lines); i += 3 {
		set1 := make(map[rune]bool)
		set2 := make(map[rune]bool)
		line1 := lines[i]
		line2 := lines[i+1]
		line3 := lines[i+2]

		if len(line2) < len(line1) && len(line2) < len(line3) {
			line1, line2 = line2, line1
		} else if len(line3) < len(line1) && len(line3) < len(line2) {
			line1, line3 = line3, line1
		}

		for c := range line2 {
			set1[line2[c]] = true
		}

		for c := range line3 {
			set2[line3[c]] = true
		}

		for c := range line1 {
			if set1[line1[c]] && set2[line1[c]] {
				total += getCharCode(line1[c])
				break
			}
		}
	}

	fmt.Println(total)
}
