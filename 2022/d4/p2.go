package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Error reading file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error scanning file", err)
	}

	total := 0

	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")
		assignment1 := strings.Split(assignments[0], "-")
		assignment2 := strings.Split(assignments[1], "-")

		assignment1Start, err := strconv.Atoi(assignment1[0])
		if err != nil {
			log.Fatalln("Error converting to int", err)
		}
		assignment1End, err := strconv.Atoi(assignment1[1])
		if err != nil {
			log.Fatalln("Error converting to int", err)
		}
		assignment2Start, err := strconv.Atoi(assignment2[0])
		if err != nil {
			log.Fatalln("Error converting to int", err)
		}
		assignment2End, err := strconv.Atoi(assignment2[1])
		if err != nil {
			log.Fatalln("Error converting to int", err)
		}

		if (assignment1Start >= assignment2Start && assignment1Start <= assignment2End) || (assignment1End >= assignment2Start && assignment1End <= assignment2End) || (assignment2Start >= assignment1Start && assignment2Start <= assignment1End) || (assignment2End >= assignment1Start && assignment2End <= assignment1End) {
			total++
		}
	}

	fmt.Println(total)
}
