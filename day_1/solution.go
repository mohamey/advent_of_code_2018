package main

import (
    "fmt"
    "strings"
    "io/ioutil"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)

    str_data := string(data)
    lines := strings.Split(str_data, "\n")

    frequency := 0
    var frequencies map[int]bool
    frequencies = make(map[int]bool)
    frequencies[frequency] = true
    repeat_found := false

    for repeat_found == false {
        fmt.Println("looping...")
        for i:=0; i < len(lines); i++ {
            if lines[i] != "" {
                num, err := strconv.Atoi(lines[i])
                check(err)
                frequency = frequency + num

                // Check if frequency has already occurred
                _, key_exists := frequencies[frequency]
                if key_exists {
                    fmt.Println("Repeat Frequency: ", frequency)
                    repeat_found = true
                    break
                } else {
                    frequencies[frequency] = true
                }
            }
        }
    }
}

