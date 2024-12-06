package main

import (
	"fmt"
	"os"
	"sort"
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
	// fmt.Println("Lines:", lines)

	rules := []string{}
	pages := [][]string{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, "|") {
			rules = append(rules, line)
		} else {
			pages = append(pages, strings.Split(line, ","))
		}
	}

	var result_part1 int
	var result_part2 int
	for _, page := range pages {
		if followsRules(page, rules) {
			// fmt.Println("True:", page)
			middleIndex := len(page) / 2
			middleNumber, err := strconv.Atoi(page[middleIndex])
			if err != nil {
				fmt.Println(err)
				continue
			}
			result_part1 += middleNumber

		} else {
			// fmt.Println("False:", page)
			sortPageByRules(page, rules)
			middleIndex := len(page) / 2
			middleNumber, err := strconv.Atoi(page[middleIndex])
			if err != nil {
				fmt.Println(err)
				continue
			}
			result_part2 += middleNumber
		}
	}
	fmt.Println("Result 1:", result_part1)
	fmt.Println("Result 2:", result_part2)
}

func followsRules(page []string, rules []string) bool {
	for _, rule := range rules {
		parts := strings.SplitN(rule, "|", 2)
		if len(parts) != 2 {
			continue
		}

		before := parts[0]
		after := parts[1]

		beforeIndex := -1
		afterIndex := -1

		for i, p := range page {
			if strings.TrimSpace(p) == before {
				beforeIndex = i
			}
			if strings.TrimSpace(p) == after {
				afterIndex = i
			}

			if beforeIndex != -1 && afterIndex != -1 {
				if beforeIndex > afterIndex {
					return false
				}
			}
		}
	}

	return true
}

func sortPageByRules(page []string, rules []string) {
	sort.Slice(page, func(i, j int) bool {
		for _, rule := range rules {
			parts := strings.SplitN(rule, "|", 2)
			if len(parts) != 2 {
				continue
			}

			before := parts[0]
			after := parts[1]

			if page[i] == before && page[j] == after {
				return true
			}
			if page[i] == after && page[j] == before {
				return false
			}
		}
		return page[i] < page[j]
	})
}
