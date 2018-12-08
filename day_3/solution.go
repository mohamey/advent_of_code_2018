package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

// Constants used in program
const claim_num = "claim"
const x_cor = "x"
const y_cor = "y"
const height = "height"
const width = "width"

// Parse each claim into a map
// Claims are of the format:
// #1 @ 146,196: 19x14
func parse_line(line string) map[string]int {
    var res_map map[string]int
    res_map = make(map[string]int)

    line_parts := strings.Split(line, " ")
    num_string := line_parts[0][1:]
    coords_string := line_parts[2][:len(line_parts[2]) - 1]
    shape_string := line_parts[3]

    coords_parts := strings.Split(coords_string, ",")
    shape_parts := strings.Split(shape_string, "x")

    res_map[claim_num], _ = strconv.Atoi(num_string)
    res_map[x_cor], _ = strconv.Atoi(coords_parts[0])
    res_map[y_cor], _ = strconv.Atoi(coords_parts[1])
    res_map[width], _ = strconv.Atoi(shape_parts[0])
    res_map[height], _ = strconv.Atoi(shape_parts[1])

    return res_map
}

func main() {
    // Read data, split into lines
    data, _ := ioutil.ReadFile("input.txt")
    str_data := string(data)
    lines := strings.Split(str_data, "\n")

    // Map is a representation of the fabric
    var fabric_map map[string]int
    fabric_map = make(map[string]int)

    claims := make([]map[string]int, 0)
    for _, line := range lines {
        if line == "" {
            continue
        }

        claim := parse_line(string(line))
        claims = append(claims, claim)

        // For each of the claim's coordinates,
        // increment number of claims on each square inch
        // of the fabric
        for x := 0; x < claim[width]; x++ {
            tmp_x := claim[x_cor] + x
            str_x := strconv.Itoa(tmp_x)
            for y := 0; y < claim[height]; y++ {
                tmp_y := claim[y_cor] + y
                str_y := strconv.Itoa(tmp_y)

                key := str_x + "x" + str_y
                fabric_map[key]++

            }
        }
    }

    // Iterate the fabric map, determining the
    // number of claims on each known square inch
    count := 0
    for _, v := range fabric_map {
        if v > 1 {
            count ++
        }
    }
    fmt.Println("Overlapping square inches:", count)

    // Problem 2
    // TODO: Find a better way of doing this

    // Loop through each claim, check if the claims
    // coordinates are contested using previously constructed
    // fabric map
    for _, claim := range claims {
        valid_claim := true

        for x := 0; x < claim[width]; x++ {
            tmp_x := claim[x_cor] + x
            str_x := strconv.Itoa(tmp_x)
            for y := 0; y < claim[height]; y++ {
                tmp_y := claim[y_cor] + y
                str_y := strconv.Itoa(tmp_y)

                key := str_x + "x" + str_y

                if fabric_map[key] > 1 {
                    valid_claim = false
                    break
                }
            }

            if !valid_claim {
                break
            }
        }

        if valid_claim {
            fmt.Println("Valid claim found, claim #", claim[claim_num])
            break
        }
    }

}
