package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	utils "mromero1591/advent-of-code/utilities"

	"github.com/TwiN/go-color"
)

var (
	year = "2023"
	day  = "3"
)

func Part1(input string) int {
	output := 0

	matrix := convertToMatrix(input)

	validInputs := []string{}

	for i, row := range matrix {
		currentNumber := ""
		adjacent := false
		for j, col := range row {
			if unicode.IsDigit(rune(col[0])) {
				currentNumber += col
				if !adjacent {
					adjacent = isAdjacent(matrix, i, j)
				}

				continue
			} else {
				if currentNumber == "" {
					adjacent = false
					continue
				} else {
					if adjacent {
						validInputs = append(validInputs, currentNumber)
					}
					currentNumber = ""
					adjacent = false
				}
			}
		}
	}

	fmt.Println(validInputs)
	for _, validInput := range validInputs {
		validNumber, err := strconv.Atoi(validInput)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		output += validNumber
	}
	return output
}

func Part2(input string) int {
	output := 0
	return output
}

func main() {
	year, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println("Error converting year to int:", err)
		return
	}
	day, err := strconv.Atoi(day)
	if err != nil {
		fmt.Println("Error converting day to int:", err)
		return
	}
	startTime := time.Now()
	input := utils.GetInputStringByName(year, day, "puzzle.txt")
	puzzleParseTime := time.Now()
	result1 := Part1(input)
	part1Time := time.Now()
	result2 := Part2(input)
	part2Time := time.Now()

	printWidth := 25
	title := fmt.Sprintf("Advent of Code %d - Day %d", year, day)

	fmt.Println(color.Bold + color.Purple + title + color.Reset)
	fmt.Println(color.Bold + color.Blue + strings.Repeat("=", 55) + color.Reset)
	fmt.Println(color.Bold + color.Blue + "Input:" + color.Reset)
	fmt.Printf(
		color.Bold+"Lines: %-*d\tParse: %v\n"+color.Reset,
		printWidth,
		len(input),
		puzzleParseTime.Sub(startTime),
	)
	fmt.Println(color.Bold + color.Blue + "Part 1:" + color.Reset)
	fmt.Printf(
		color.Bold+"Result: %-*d\tTime: %v\n"+color.Reset,
		printWidth,
		result1,
		part1Time.Sub(puzzleParseTime),
	)
	fmt.Println(color.Bold + color.Blue + "Part 2:" + color.Reset)
	fmt.Printf(
		color.Bold+"Result: %-*d\tTime: %v\n"+color.Reset,
		printWidth,
		result2,
		part2Time.Sub(part1Time),
	)
	fmt.Println(color.Bold+color.Blue+"Total Time:"+color.Reset, part2Time.Sub(startTime))
	fmt.Print("\n")
}

func convertToMatrix(str string) [][]string {
	rows := strings.Split(str, "\n")
	matrix := make([][]string, len(rows)-1)

	for i, row := range rows {
		if row == "" {
			continue
		}
		columns := strings.Split(row, "")
		matrix[i] = columns
	}
	return matrix
}

func isAdjacent(matrix [][]string, row int, col int) bool {
	firstRow := 0
	firstCol := 0
	lastRow := len(matrix) - 1
	lastCol := len(matrix[0]) - 1

	// is adjacent up
	if row != firstRow {
		char := matrix[row-1][col]
		if char != "." && !unicode.IsDigit(rune(char[0])) {
			return true
		}
	}

	// is adjacent down
	if row != lastRow {
		char := matrix[row+1][col]
		if char != "." && !unicode.IsDigit(rune(char[0])) {
			return true
		}
	}

	// is adjacent left
	if col != firstCol {
		char := matrix[row][col-1]
		if char != "." && !unicode.IsDigit(rune(char[0])) {
			return true
		}
	}

	// is adjacent right
	if col != lastCol {
		char := matrix[row][col+1]
		if char != "." && !unicode.IsDigit(rune(char[0])) {
			return true
		}
	}

	// is adjacent upleft
	// is adjacent up
	if row != firstRow && col != firstCol {
		char := matrix[row-1][col-1]
		if char != "." && !unicode.IsDigit(rune(char[0])) {
			return true
		}
	}

	// is adjacent upright
	if row != firstRow && col != lastCol {
		char := matrix[row-1][col+1]
		if char != "." && !unicode.IsDigit(rune(char[0])) {
			return true
		}
	}

	// is adacent downleft
	if row != lastRow && col != firstCol {
		char := matrix[row+1][col-1]
		if char != "." && !unicode.IsDigit(rune(char[0])) {
			return true
		}
	}

	// is adacent downright
	if row != lastRow && col != lastCol {
		char := matrix[row+1][col+1]
		if char != "." && !unicode.IsDigit(rune(char[0])) {
			return true
		}
	}

	return false
}
