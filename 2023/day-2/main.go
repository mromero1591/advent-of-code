package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"

	utils "ashmortar/advent-of-code/utilities"
)

var (
	year = "2023"
	day  = "2"
)

type Game struct {
	Title    string
	Spot     int
	Outcomes [][]GameOutcome
}

type GameOutcome struct {
	Color string
	Value int
}

func Part1(input string) int {
	output := 0
	inputArray := strings.Split(input, "\n")
	games := []Game{}

	for _, gameStr := range inputArray {
		game, err := parseGame(gameStr)
		if err != nil {
			log.Fatal("Failed to parse game", err)
		}
		games = append(games, game)
	}

	// only 12 red cubes, 13 green cubes, and 14 blue cubes?
	gameCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	results := map[int]bool{}
	for _, game := range games {
		for _, outcome := range game.Outcomes {
			validGame := determineValidGame(gameCubes, outcome)
			if !validGame {
				results[game.Spot] = validGame
				break
			}
			results[game.Spot] = true
		}

	}

	for spot, valid := range results {
		if valid {
			output += spot
		}
	}
	return output
}

func Part2(input string) int {
	output := 0
	inputArray := strings.Split(input, "\n")
	games := []Game{}

	for _, gameStr := range inputArray {
		game, err := parseGame(gameStr)
		if err != nil {
			log.Fatal("Failed to parse game", err)
		}
		games = append(games, game)
	}
	for _, game := range games {
		set := highestValues(game.Outcomes)
		if len(set) == 3 {
			power := 1
			for _, val := range set {
				power *= val
			}
			output += power
		}

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

func parseGame(gameString string) (Game, error) {
	if gameString == "" {
		return Game{}, nil
	}
	gameParts := strings.Split(gameString, ":")
	if len(gameParts) != 2 {
		return Game{}, fmt.Errorf("invalid game format: \t%s", gameString)
	}

	gameTitle := gameParts[0]
	gameSpot := 0

	_, err := fmt.Sscanf(gameTitle, "Game %d", &gameSpot)
	if err != nil {
		fmt.Println("invalid game format for game spot:", err)
	}

	outcomeStrings := strings.Split(gameParts[1], ";")

	var gameOutcomes [][]GameOutcome

	for _, outcomeGroup := range outcomeStrings {
		outcomeParts := strings.Split(strings.TrimSpace(outcomeGroup), ",")

		var outcomes []GameOutcome

		for _, outcome := range outcomeParts {
			valueColor := strings.Fields(strings.TrimSpace(outcome))
			if len(valueColor) != 2 {
				return Game{}, fmt.Errorf("invalid outcome format: \t%s", gameString)
			}

			value, err := strconv.Atoi(valueColor[0])
			if err != nil {
				return Game{}, fmt.Errorf("invalid value format: %w", err)
			}

			outcomes = append(outcomes, GameOutcome{
				Color: valueColor[1],
				Value: value,
			})
		}

		gameOutcomes = append(gameOutcomes, outcomes)
	}

	return Game{
		Title:    gameTitle,
		Spot:     gameSpot,
		Outcomes: gameOutcomes,
	}, nil
}

func determineValidGame(gameCubes map[string]int, gameResults []GameOutcome) bool {
	for _, results := range gameResults {
		cubes, inlcuded := gameCubes[results.Color]
		if !inlcuded || cubes < results.Value {
			return false
		}

	}
	return true
}

func highestValues(data [][]GameOutcome) map[string]int {
	colorValues := make(map[string]int)

	for _, set := range data {
		for _, colorVal := range set {
			if currentMax, exists := colorValues[colorVal.Color]; exists {
				if colorVal.Value > currentMax {
					colorValues[colorVal.Color] = colorVal.Value
				}
			} else {
				colorValues[colorVal.Color] = colorVal.Value
			}
		}
	}

	return colorValues
}
