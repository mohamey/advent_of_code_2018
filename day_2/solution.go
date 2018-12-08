package main

import (
    "fmt"
    "io/ioutil"
    "strings"
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

    var twos_count, threes_count int

    for _, line := range lines {
        var char_map map[string]int
        char_map = make(map[string]int)
        var pair_found, triplet_found bool

        for _, char := range line {
            char_str := string(char)
            char_map[char_str]++
        }

        for _, v := range char_map {
            if v == 2 && !pair_found{
                twos_count++
                pair_found = true
            } else if v == 3 && !triplet_found {
                threes_count++
                triplet_found = true
            }
        }

    }

    fmt.Println("Pairs found: ", twos_count)
    fmt.Println("Triplets found: ", threes_count)
    fmt.Println("Checksum: ", twos_count * threes_count)

}
