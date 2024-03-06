package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// partOne(file)
	partTwo(file)
	file.Close()
}

func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		isValid := true

		gameSplit := strings.Split(line, ": ")

		gameId, err := strconv.Atoi(strings.Split(gameSplit[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		sets := strings.Split(gameSplit[1], "; ")

		for _, set := range sets {
			// INFO: Balls
			balls := strings.Split(set, ", ")

			for _, ball := range balls {
				ballVal := strings.Split(ball, " ")

				val, err := strconv.Atoi(ballVal[0])
				if err != nil {
					log.Fatal(err)
				}

				if (ballVal[1] == "blue" && val > MAX_BLUE) || (ballVal[1] == "red" && val > MAX_RED) || (ballVal[1] == "green" && val > MAX_GREEN) {
					isValid = false
					break
				}
			}

			if !isValid {
				break
			}
		}

		if isValid {
			total += gameId
		}
	}

	fmt.Printf("P1: %d\n", total)
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		gameSplit := strings.Split(line, ": ")

		sets := strings.Split(gameSplit[1], "; ")

		hash := map[string]int{
			"green": 0,
			"red":   0,
			"blue":  0,
		}

		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				cubeVal := strings.Split(cube, " ")
				val, err := strconv.Atoi(cubeVal[0])
				if err != nil {
					log.Fatal(err)
				}
				if hash[cubeVal[1]] < val {
					hash[cubeVal[1]] = val
				}
			}
		}

		powerSet := 1

		for _, val := range hash {
            powerSet *= val
		}

		total += powerSet
	}

	fmt.Printf("P2: %d\n", total)
}
