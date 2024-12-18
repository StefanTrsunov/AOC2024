package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Println("Lines:", lines)

	var result_part1_slice []int
	var result_part1 int

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		number, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println(err)
			continue
		}

		number_slice := strings.Fields(parts[1])
		var numberSlice []int
		for _, numStr := range number_slice {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println(err)
				continue
			}
			numberSlice = append(numberSlice, num)
		}

		if checkCombinations(number, numberSlice) {
			result_part1_slice = append(result_part1_slice, number)
		}

	}

	for _, number := range result_part1_slice {
		result_part1 += number
	}
	fmt.Println("Result 1:", result_part1)
}

func checkCombinations(number int, numberSlice []int) bool {
	evaluate := func(ops []string) int {
		result := numberSlice[0]
		for i := 1; i < len(numberSlice); i++ {
			if ops[i-1] == "+" {
				result += numberSlice[i]
			} else if ops[i-1] == "*" {
				result *= numberSlice[i]
			}
		}
		return result
	}

	n := len(numberSlice) - 1
	ops := make([]string, n)

	for i := 0; i < (1 << n); i++ {
		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				ops[j] = "+"
			} else {
				ops[j] = "*"
			}
		}
		if evaluate(ops) == number {
			return true
		}
	}
	return false
}
