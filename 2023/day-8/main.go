package main

import (
	"fmt"
	utils "mromero1591/advent-of-code/utilities"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

var (
	year = "2023"
	day  = "8"
)

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
	title := fmt.Sprintf("Advent of Code %d - Day %d", year, day)

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

func Part1(input string) int {
	output := 0
	maze, directions := formatInput(input)
	output = moveThroughMaze(maze, directions)
	return output
}

func Part2(input string) int {
	output := 0
	maze, directions := formatInput(input)
	output = moveThroughMaze2(maze, directions)
	return output
}

func formatInput(input string) (map[string]Node, []string) {
	inputLines := strings.Split(input, "\n")
	directions := strings.Split(inputLines[0], "")
	inputLines = inputLines[1:]

	maze := map[string]Node{}

	for _, line := range inputLines {
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		node := Node{}

		fmt.Sscanf(line, "%s = (%s %s)", &node.Name, &node.L, &node.R)
		node.L = strings.ReplaceAll(node.L, ",", "")
		node.R = strings.ReplaceAll(node.R, ")", "")

		maze[node.Name] = node
	}
	return maze, directions
}

type Node struct {
	Name string
	L    string
	R    string
}

func moveThroughMaze(maze map[string]Node, directions []string) int {
	endLocation := "ZZZ"
	currentLocation := maze["AAA"]
	directionPointer := 0
	stepCount := 0

	for currentLocation.Name != endLocation {
		// L or R
		moveTo := directions[directionPointer]
		// node props.
		v := reflect.ValueOf(currentLocation)

		//next key to move to
		moveToKey := v.FieldByName(moveTo)

		newLocation := maze[moveToKey.String()]
		currentLocation = newLocation
		if directionPointer == len(directions)-1 {
			directionPointer = 0
		} else {
			directionPointer++
		}
		stepCount++
	}
	return stepCount
}

func moveThroughMaze2(maze map[string]Node, directions []string) int {
	currentLocations := []Node{}
	//find starting locations
	for _, node := range maze {
		if node.Name[len(node.Name)-1] == 65 {
			currentLocations = append(currentLocations, node)
		}
	}

	allEndCounts := []int{}

	for _, node := range currentLocations {
		stepsCount := moveThroughMazeByLastLetter(maze, directions, node)

		allEndCounts = append(allEndCounts, stepsCount)
	}

	lcm := allEndCounts[0]
	for i := 1; i < len(allEndCounts); i++ {
		lcm = findLcm(lcm, allEndCounts[i])
	}
	return lcm
}

func moveThroughMazeByLastLetter(maze map[string]Node, directions []string, start Node) int {
	currentLocation := start
	directionPointer := 0
	stepCount := 0

	for !atEndLocation(currentLocation.Name) {
		// L or R
		moveTo := directions[directionPointer]
		// node props.
		v := reflect.ValueOf(currentLocation)

		//next key to move to
		moveToKey := v.FieldByName(moveTo)

		newLocation := maze[moveToKey.String()]
		currentLocation = newLocation
		if directionPointer == len(directions)-1 {
			directionPointer = 0
		} else {
			directionPointer++
		}
		stepCount++
	}
	return stepCount
}

func atEndLocation(name string) bool {
	return name[len(name)-1] == 90
}

func findGcd(a, b int) int {
	if b == 0 {
		return a
	}
	return findGcd(b, a%b)
}

func findLcm(a, b int) int {
	return a * b / findGcd(a, b)
}
