package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func (s Stack) Push(v string) Stack {
	return append(s, v)
}

func (s Stack) Pop() (string, Stack) {
	l := len(s)
	res := s[l-1]
	return res, s[:l-1]
}

func getChar(str string, index int) rune {
    return []rune(str)[index]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Error reading file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	crateLines := []string{}
	crates := map[int]Stack{}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error scanning file", err)
	}

	for scanner.Scan() {
		line := []rune(scanner.Text())
		if strings.Contains(scanner.Text(), "[") {
			crateLines = append(crateLines, scanner.Text())
		} else if scanner.Text() ==  " 1   2   3   4   5   6   7   8   9 " {
			for i := 1; i < 10; i++ {
				crates[i] = Stack{}
			}
			for i := len(crateLines) - 1; i >= 0; i-- {
				crateLine := crateLines[i]
				crateNum := 1
				for j := 1; j < len(crateLine); j += 4 {
					if string(crateLine[j]) != " " {
						crates[crateNum] = crates[crateNum].Push(string(crateLine[j]))
					}
					crateNum++
				}
			}
		} else if strings.Contains(scanner.Text(), "move") {
            dir := strings.Split(string(line), " ")
            // 1 3 5
            count, err := strconv.Atoi(dir[1])
            if err != nil {
                log.Fatalln("Error converting count", err)
            }
            from, err := strconv.Atoi(dir[3])
            if err != nil {
                log.Fatalln("Error converting from", err)
            }
            to, err := strconv.Atoi(dir[5])
            if err != nil {
                log.Fatalln("Error converting to", err)
            }
            temp := []string{}
            for i := 0; i < count; i++ {
                val, s := crates[from].Pop()
                crates[from] = s
                temp = append(temp, val)
            }
            for i := len(temp) - 1; i >= 0; i-- {
                crates[to] = crates[to].Push(temp[i])
            }
		}
	}

    result := ""

    for i := 1; i < 10; i++ {
        val, _ := crates[i].Pop()
        result += val
    }

    fmt.Println(result)
}
