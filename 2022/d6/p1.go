package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getDistinctCount(set *map[rune]int) int {
    count := 0
    for _, v := range *set {
        if v == 1 {
            count++
        }
    }
    return count
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Error reading file", err)
	}
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatalln("Error scanning file", err)
	}
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
	}
	set := map[rune]int{}
	for i := rune('a'); i <= rune('z'); i++ {
		set[i] = 0
	}
	for i := 0; i < 3; i++ {
		set[rune(line[i])]++
	}
	l := 0
	res := 0

	for r := 3; r < len(line); r++ {
		set[rune(line[r])]++
        if getDistinctCount(&set) == 4 {
			res = r + 1
			break
		}
		set[rune(line[l])]--
		l++
	}

	fmt.Println(res)
}
