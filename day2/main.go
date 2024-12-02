package main

import (
	"fmt"
	"os"
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
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(data), "\n")
	var result1 int
	var result2 int

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		numbers_slice := make([]int, len(parts))

		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println(err)
				break
			}
			numbers_slice[i] = num
		}
		var safe_1 bool = true
		var incresing_counter_1 int = 0
		var decreasing_counter_1 int = 0

		for i := 0; i < len(numbers_slice)-1; i++ {
			diff := Abs(numbers_slice[i] - numbers_slice[i+1])
			if diff < 1 || diff > 3 {
				safe_1 = false
				break
			}

			if numbers_slice[i] < numbers_slice[i+1] {
				incresing_counter_1++
			} else if numbers_slice[i] > numbers_slice[i+1] {
				decreasing_counter_1++
			}
		}

		if safe_1 && (incresing_counter_1 == len(parts)-1 || decreasing_counter_1 == len(parts)-1) {
			result1++
		} else {
			safe_1 = false
		}

		for !safe_1 {
			for i := 0; i < len(numbers_slice); i++ {
				var empty_slice []int
				begin := append(empty_slice, numbers_slice[:i]...)
				new_slice := append(begin, numbers_slice[i+1:]...)
				var safe_2 bool = true
				var incresing_counter_2 int = 0
				var decreasing_counter_2 int = 0

				for j := 0; j < len(new_slice)-1; j++ {
					diff := Abs(new_slice[j] - new_slice[j+1])

					if diff < 1 || diff > 3 {
						safe_2 = false
						break
					}

					if new_slice[j] < new_slice[j+1] {
						incresing_counter_2 += 1
					} else if new_slice[j] > new_slice[j+1] {
						decreasing_counter_2 += 1
					}
				}

				if safe_2 && (incresing_counter_2 == len(new_slice)-1 || decreasing_counter_2 == len(new_slice)-1) {
					result2++
					break
				}
			}
			break
		}
	}

	fmt.Println("Result1 :", result1)
	fmt.Println("Result2 :", result2)
	fmt.Println("Total Result:", result1+result2)
}
