package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tree struct {
    left int
    right int
    up int
    down int
}

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
    trees := make(map[[2]int]*Tree)

    for i := 1; i < len(lines) - 1; i++ {
        for j := 1; j < len(lines[i]) - 1; j++ {
            left := 0
            right := 0

            height := lines[i][j]

            k := j - 1

            for k >= 0 {
                left++
                if lines[i][k] >= height {
                    break
                }
                k--
            }

            k = j + 1

            for k < len(lines[i]) {
                right++
                if lines[i][k] >= height {
                    break
                }
                k++
            }

            trees[[2]int{i, j}] = &Tree{left, right, 0, 0}
        }
    }

    for i := 1; i < len(lines[0]) - 1; i++ {
        for j := 1; j < len(lines) - 1; j++ {
            up := 0
            down := 0

            height := lines[j][i]

            k := j - 1

            for k >= 0 {
                up++
                if lines[k][i] >= height {
                    break
                }
                k--
            }

            k = j + 1

            for k < len(lines) {
                down++
                if lines[k][i] >= height {
                    break
                }
                k++
            }

            t := trees[[2]int{j, i}]
            score := t.left * t.right * up * down
            if score > max {
                max = score
            }
        }
    }

    fmt.Println(max)
}
