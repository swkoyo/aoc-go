package main

import (
	"fmt"
    "log"
    "bufio"
    "os"
    "strings"
)

func getHandScore(hand string) int {
    if (hand == "X") {
        return 1;
    } else if (hand == "Y") {
        return 2;
    } else {
        return 3;
    }
}

func main() {
    file, err := os.Open("./input.txt");
    if err != nil {
        log.Fatalln("Error reading file", err)
    }

    defer file.Close();

    scanner := bufio.NewScanner(file);

    score := 0;

    for scanner.Scan() {
        if (scanner.Text() != "") {
            line := scanner.Text();
            hand := strings.Split(line, " ");
            score += getHandScore(hand[1]);
            if ((hand[0] == "A" && hand[1] == "Z") || (hand[0] == "B" && hand[1] == "X") || (hand[0] == "C" && hand[1] == "Y")) {
                score += 0;
            } else if ((hand[0] == "A" && hand[1] == "Y") || (hand[0] == "B" && hand[1] == "Z") || (hand[0] == "C" && hand[1] == "X")) {
                score += 6;
            } else if ((hand[0] == "A" && hand[1] == "X") || (hand[0] == "B" && hand[1] == "Y") || (hand[0] == "C" && hand[1] == "Z")) {
                score += 3;
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatalln("Error scanning file", err)
    }

    fmt.Println(score);
}
