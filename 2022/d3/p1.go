package main

import (
	"fmt"
    "log"
    "bufio"
    "os"
    "unicode"
)

func getCharCode(letter rune) int {
    if (unicode.IsUpper(letter)) {
       return int(letter) - 38; 
    }
    return int(letter) - 96;
}

func main() {
    file, err := os.Open("./input.txt");
    if err != nil {
        log.Fatalln("Error reading file", err)
    }

    defer file.Close();

    scanner := bufio.NewScanner(file);

    if err := scanner.Err(); err != nil {
        log.Fatalln("Error scanning file", err)
    }

    total := 0;

    set := make(map[rune]bool);

    for scanner.Scan() {
        line := []rune(scanner.Text());
        mid := len(scanner.Text()) / 2;
        for k := range set {
            delete(set, k)
        }
        for i := 0; i < mid; i++ {
            set[line[i]] = true;
        }
        
        for i := mid; i < len(scanner.Text()); i++ {
            if (set[line[i]]) {
                total += getCharCode(line[i]);
                break;
            }
        }
    }

    fmt.Println(total);
}
