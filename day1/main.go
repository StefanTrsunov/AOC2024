package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	var leftNumbers []int
	var rightNumbers []int

	for _, line := range lines {
		line = strings.TrimSpace(line)

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting numbers:", line)
			continue
		}

		leftNumbers = append(leftNumbers, num1)
		rightNumbers = append(rightNumbers, num2)
	}
	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	var result_part_1 int

	for i := 0; i < len(leftNumbers); i++ {
		result_part_1 += Abs(rightNumbers[i] - leftNumbers[i])
	}

	var counter int
	var result_part_2 int
	for i := 0; i < len(leftNumbers); i++ {
		counter = 0
		for j := 0; j < len(rightNumbers); j++ {
			if leftNumbers[i] == rightNumbers[j] {

				counter += 1
			}
		}
		result_part_2 += counter * leftNumbers[i]
	}

	fmt.Println("Result 1:", result_part_1)
	fmt.Println("Result 2:", result_part_2)
}
