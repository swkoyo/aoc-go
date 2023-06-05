package main

import (
	"fmt"
    "log"
    "bufio"
    "os"
    "strconv"
    "sort"
)

func main() {
    file, err := os.Open("./input.txt");
    if err != nil {
        log.Fatalln("Error reading file", err)
    }

    defer file.Close();

    scanner := bufio.NewScanner(file);

    max := make([]int, 3);
    count := 0;

    for scanner.Scan() {
        if scanner.Text() == "" {
            if count > max[0] {
                max[0] = count;
                sort.Ints(max);
            }
            count = 0;
            continue;
        }
        val, err := strconv.Atoi(scanner.Text());
        if err != nil {
            log.Fatalln("Error converting string to int", err)
        }
        count += val;
    }

    if err := scanner.Err(); err != nil {
        log.Fatalln("Error scanning file", err)
    }

    result := 0;

    for _, v := range max {
        result += v;
    }

    fmt.Printf("Result: %v\n", result);
}

