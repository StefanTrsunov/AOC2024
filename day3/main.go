package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(data), "\n")

	mul_substring := regexp.MustCompile(`mul\(\d+,\d+\)`)

	var result_part_1 int
	for _, line := range lines {
		matches := mul_substring.FindAllString(line, -1)
		for _, match := range matches {
			match_number := regexp.MustCompile(`\d+`)
			match_numbers := match_number.FindAllString(match, -1)

			match_numbers_int := []int{}
			for _, num := range match_numbers {
				num_int, err := strconv.Atoi(num)
				if err != nil {
					fmt.Println(err)
					continue
				}
				match_numbers_int = append(match_numbers_int, num_int)
			}

			multiplied_numbers := match_numbers_int[0] * match_numbers_int[1]
			result_part_1 += multiplied_numbers
		}
	}
	fmt.Println("Result 1:", result_part_1)

	var result_part_2 int
	doRegex := regexp.MustCompile(`do\(\)`)
	donotRegex := regexp.MustCompile(`don't\(\)`)

	for _, line := range lines {
		for {
			doIndex := doRegex.FindStringIndex(line)
			donotIndex := donotRegex.FindStringIndex(line)

			if doIndex == nil && donotIndex == nil {
				break
			}

			var substring string
			if doIndex != nil && (donotIndex == nil || doIndex[1] < donotIndex[0]) {
				substring = line[doIndex[1]:]
			} else if donotIndex != nil {
				if doIndex != nil && doIndex[1] < donotIndex[0] {
					substring = line[doIndex[1]:donotIndex[0]]
				} else {
					substring = line[:donotIndex[0]]
				}
			}

			mulMatches := mul_substring.FindAllString(substring, -1)
			for _, match := range mulMatches {
				match_number := regexp.MustCompile(`\d+`)
				match_numbers := match_number.FindAllString(match, -1)

				match_numbers_int := []int{}
				for _, num := range match_numbers {
					num_int, err := strconv.Atoi(num)
					if err != nil {
						fmt.Println(err)
						continue
					}
					match_numbers_int = append(match_numbers_int, num_int)
				}

				multiplied_numbers := match_numbers_int[0] * match_numbers_int[1]
				result_part_2 += multiplied_numbers
			}
			if donotIndex == nil {
				line = line[doIndex[1]:]
			} else if doIndex == nil {
				line = line[donotIndex[1]:]
			} else if doIndex[1] < donotIndex[0] {
				line = line[doIndex[1]:]
			} else {
				line = line[donotIndex[1]:]
			}

		}

	}
	fmt.Println("Result 2:", result_part_2)
}
