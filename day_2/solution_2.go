package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    // Read file as bytes
    data, err := ioutil.ReadFile("input.txt")
    check(err)

    // Format data, split text on lines
    str_data := string(data)
    lines := strings.Split(str_data, "\n")

    // Calculate how many characters need to appear in both strings
    required_overlap := len(lines[0]) - 1

    // Iterate over all lines
    for index, line := range lines {
        // Read the current line's characters into a map
        var char_map map[int]string
        char_map = make(map[int]string)

        for i, char := range line {
            char_map[i] = string(char)
        }

        // Loop through all the lines that come after current line
        // and find characters that overlap
        for j := index + 1; j < len(lines); j++ {
            var same_chars []string

            for tmp_char_index, tmp_char := range lines[j] {
                tmp_char_str := string(tmp_char)

                char, key_exists := char_map[tmp_char_index]
                if key_exists && (string(char) == tmp_char_str) {
                    same_chars = append(same_chars, tmp_char_str)
                }
            }

            if len(same_chars) == required_overlap {
                fmt.Println("Similar Packages found, common chars are:", strings.Join(same_chars, ""))
                break
            }
        }
    }

}
