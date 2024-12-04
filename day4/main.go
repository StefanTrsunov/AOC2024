package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	board := [][]string{}

	part1_word := "XMAS"
	result_part1 := 0
	result_part2 := 0

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		row := strings.Split(line, "")

		board = append(board, row)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			// 1. top-bottom
			if j <= len(board)-4 {
				if fmt.Sprintf("%s%s%s%s", board[j][i], board[j+1][i], board[j+2][i], board[j+3][i]) == part1_word {
					result_part1++
				}
			}
			// 2. bottom-top
			if j >= 3 {
				if fmt.Sprintf("%s%s%s%s", board[j][i], board[j-1][i], board[j-2][i], board[j-3][i]) == part1_word {
					result_part1++
				}
			}
			// 3. left-right
			if i <= len(board[j])-4 {
				if fmt.Sprintf("%s%s%s%s", board[j][i], board[j][i+1], board[j][i+2], board[j][i+3]) == part1_word {
					result_part1++
				}
			}
			// 4. right-left
			if i >= 3 {
				if fmt.Sprintf("%s%s%s%s", board[j][i], board[j][i-1], board[j][i-2], board[j][i-3]) == part1_word {
					result_part1++
				}
			}
			// 5. bottomleft-topright
			if j >= 3 && i <= len(board[j])-4 {
				if fmt.Sprintf("%s%s%s%s", board[j][i], board[j-1][i+1], board[j-2][i+2], board[j-3][i+3]) == part1_word {
					result_part1++
				}
			}
			// 6. topright-bottom-left
			if j <= len(board)-4 && i >= 3 {
				if fmt.Sprintf("%s%s%s%s", board[j][i], board[j+1][i-1], board[j+2][i-2], board[j+3][i-3]) == part1_word {
					result_part1++
				}
			}
			// 7. Bottomright-topleft
			if j >= 3 && i >= 3 {
				if fmt.Sprintf("%s%s%s%s", board[j][i], board[j-1][i-1], board[j-2][i-2], board[j-3][i-3]) == part1_word {
					result_part1++
				}
			}
			// 8. topleft-bottomright
			if j <= len(board)-4 && i <= len(board[j])-4 {
				if fmt.Sprintf("%s%s%s%s", board[j][i], board[j+1][i+1], board[j+2][i+2], board[j+3][i+3]) == part1_word {
					result_part1++
				}
			}
			// M.S
			// .A.
			// M.S
			if j+2 < len(board) && i+2 < len(board[j]) {
				if board[j][i] == "M" && board[j+1][i+1] == "A" && board[j+2][i+2] == "S" && board[j+2][i] == "M" && board[j][i+2] == "S" {
					result_part2++
				}
			}

			// M.M
			// .A.
			// S.S
			if j+2 < len(board) && i+2 < len(board[j]) {
				if board[j][i] == "M" && board[j][i+2] == "M" && board[j+1][i+1] == "A" && board[j+2][i] == "S" && board[j+2][i+2] == "S" {
					result_part2++
				}
			}
			// S.M
			// .A.
			// S.M
			if j+2 < len(board) && i+2 < len(board[j]) {
				if board[j][i] == "S" && board[j+1][i+1] == "A" && board[j+2][i+2] == "M" && board[j+2][i] == "S" && board[j][i+2] == "M" {
					result_part2++
				}
			}
			// S.S
			// .A.
			// M.M
			if j+2 < len(board) && i+2 < len(board[j]) {
				if board[j][i] == "S" && board[j][i+2] == "S" && board[j+1][i+1] == "A" && board[j+2][i] == "M" && board[j+2][i+2] == "M" {
					result_part2++
				}
			}
		}
	}

	fmt.Println("result_part1: ", result_part1)
	fmt.Println("result_part2: ", result_part2)
}
