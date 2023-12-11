package main

import (
	utils "ashmortar/advent-of-code/utilities"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/TwiN/go-color"
)

var year = "2023"
var day = "1"

func Part1(input string) int {
	output := 0

	inputArray := strings.Split(input, "\n")
	for _, c := range inputArray {
		resp := RemoveStringChar(c)
		res := DetermineCode(resp)
		output += res
	}

	return output
}

func Part2(input string) int {
	output := 0
	inputArray := strings.Split(input, "\n")
	for _, c := range inputArray {
		preparedString := prepareString(c)
		result := RemoveStringChar(preparedString)
		code := DetermineCode(result)
		output += code
	}
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
	input := utils.GetInputString(year, day)
	puzzleParseTime := time.Now()
	result1 := Part1(input)
	part1Time := time.Now()
	result2 := Part2(input)
	part2Time := time.Now()

	printWidth := 25
	title := fmt.Sprintf("\nAdvent of Code %d - Day %d", year, day)

	fmt.Println(color.Bold + color.Purple + title + color.Reset)
	fmt.Println(color.Bold + color.Blue + strings.Repeat("=", 55) + color.Reset)
	fmt.Println(color.Bold + color.Blue + "Input:" + color.Reset)
	fmt.Printf(color.Bold+"Lines: %-*d\tParse: %v\n"+color.Reset, printWidth, len(input), puzzleParseTime.Sub(startTime))
	fmt.Println(color.Bold + color.Blue + "Part 1:" + color.Reset)
	fmt.Printf(color.Bold+"Result: %-*d\tTime: %v\n"+color.Reset, printWidth, result1, part1Time.Sub(puzzleParseTime))
	fmt.Println(color.Bold + color.Blue + "Part 2:" + color.Reset)
	fmt.Printf(color.Bold+"Result: %-*d\tTime: %v\n"+color.Reset, printWidth, result2, part2Time.Sub(part1Time))
	fmt.Println(color.Bold+color.Blue+"Total Time:"+color.Reset, part2Time.Sub(startTime))
	fmt.Print("\n")
}

func DetermineCode(s string) int {
	//if the is empty return 0
	if s == "" {
		return 0
	}
	//if the string is of size 1 return a sum of char + char
	if len(s) == 1 {
		combinedValue := s + s
		return convertCodeToNumber(combinedValue)
	}
	//return sum of char + lastChar
	firstCharacter := string(s[0])
	lastCharacter := string(s[len(s)-1])
	combinedValue := firstCharacter + lastCharacter

	return convertCodeToNumber(combinedValue)
}

func convertCodeToNumber(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("failed to convert %s to a number", s)
		return 0
	}
	return num
}

type Pair struct {
	val string
	loc int
}

func prepareString(str string) string {
	values := []Pair{}
	numbersArray := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}

	numberMap := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"0":     "0",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	for _, v := range numbersArray {
		locs := findAllSubstringIndices(str, v)
		for _, loc := range locs {
			value := numberMap[v]
			values = append(values, Pair{val: value, loc: loc})
		}
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i].loc < values[j].loc
	})

	codeString := ""
	for _, pair := range values {
		codeString += pair.val
	}

	return codeString
}

func findAllSubstringIndices(s, substr string) []int {
	var indices []int
	for i := 0; i < len(s); i++ {
		if strings.HasPrefix(s[i:], substr) {
			indices = append(indices, i)
		}
	}
	return indices
}

func RemoveStringChar(input string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, input)
}
