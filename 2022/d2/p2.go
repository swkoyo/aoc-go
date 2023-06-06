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

func getRequiredHand(playerHand string, result string) string {
    if (playerHand == "A") {
        if (result == "X") {
            return "Z";
        } else if (result == "Y") {
            return "X";
        } else {
            return "Y";
        }
    } else if (playerHand == "B") {
        if (result == "X") {
            return "X";
        } else if (result == "Y") {
            return "Y";
        } else {
            return "Z";
        }
    } else {
        if (result == "X") {
            return "Y";
        } else if (result == "Y") {
            return "Z";
        } else {
            return "X";
        }
    }
}

func getResultScore(result string) int {
    if (result == "X") {
        return 0;
    } else if (result == "Y") {
        return 3;
    } else {
        return 6;
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
            myHand := getRequiredHand(hand[0], hand[1]);
            score += getHandScore(myHand) + getResultScore(hand[1]);
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatalln("Error scanning file", err)
    }

    fmt.Println(score);
}
