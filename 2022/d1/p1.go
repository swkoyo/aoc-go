package main

import (
	"fmt"
    "log"
    "bufio"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("./input.txt");
    if err != nil {
        log.Fatalln("Error reading file", err)
    }

    defer file.Close();

    scanner := bufio.NewScanner(file);

    max := 0;
    count := 0;

    for scanner.Scan() {
        if scanner.Text() == "" {
            if count > max {
                max = count;
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

    fmt.Printf("Max: %v\n", max);
}
