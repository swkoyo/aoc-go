package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	lines := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		for _, c := range line {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			row = append(row, n)
		}
		lines = append(lines, row)
	}

	max := 0
	set := make(map[[2]int]bool)

	for i := 0; i < len(lines); i++ {
        max = 0
		for j := 0; j < len(lines[i]); j++ {
			_, ok := set[[2]int{i, j}]
			if j == 0 {
				if !ok {
					set[[2]int{i, j}] = true
				}
			} else {
				if lines[i][j-1] > max {
					max = lines[i][j-1]
				}
				if lines[i][j] > max && !ok {
					set[[2]int{i, j}] = true
				}
			}
		}
	}

	for i := len(lines) - 1; i >= 0; i-- {
        max = 0
		for j := len(lines[i]) - 1; j >= 0; j-- {
			_, ok := set[[2]int{i, j}]
			if j == len(lines[i])-1 {
				if !ok {
					set[[2]int{i, j}] = true
				}
			} else {
				if lines[i][j+1] > max {
					max = lines[i][j+1]
				}
				if lines[i][j] > max && !ok {
					set[[2]int{i, j}] = true
				}
			}
		}
	}

	for i := 0; i < len(lines[0]); i++ {
        max = 0
		for j := 0; j < len(lines); j++ {
			_, ok := set[[2]int{j, i}]
			if j == 0 {
				if !ok {
					set[[2]int{j, i}] = true
				}
			} else {
				if lines[j-1][i] > max {
					max = lines[j-1][i]
				}
				if lines[j][i] > max && !ok {
					set[[2]int{j, i}] = true
				}
			}
		}
	}

	for i := len(lines[0]) - 1; i >= 0; i-- {
        max = 0
		for j := len(lines) - 1; j >= 0; j-- {
			_, ok := set[[2]int{j, i}]
			if j == len(lines)-1 {
				if !ok {
					set[[2]int{j, i}] = true
				}
			} else {
				if lines[j+1][i] > max {
					max = lines[j+1][i]
				}
				if lines[j][i] > max && !ok {
					set[[2]int{j, i}] = true
				}
			}
		}
	}

	fmt.Println(len(set))
}
