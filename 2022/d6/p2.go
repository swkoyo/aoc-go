package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
	set := map[string]int{}
	for i := 0; i < 13; i++ {
		if _, ok := set[string(line[i])]; ok {
			set[string(line[i])]++
		} else {
			set[string(line[i])] = 1
		}
	}
	l := 0
	res := 0

	for r := 13; r < len(line); r++ {
		if _, ok := set[string(line[r])]; ok {
			set[string(line[r])]++
		} else {
			set[string(line[r])] = 1
		}
		if len(set) == 14 {
			res = r + 1
			break
		}
		set[string(line[l])]--
		if set[string(line[l])] == 0 {
			delete(set, string(line[l]))
		}
		l++
	}

	fmt.Println(res)
}
