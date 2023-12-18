package main

import (
	"fmt"
	utils "mromero1591/advent-of-code/utilities"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

var (
	year = "2023"
	day  = "6"
)

type Race struct {
	RaceTime int
	Record   int
}

func Part1(input string) int {
	output := 1
	races := formatInput(input, false)
	output = determineOutcomes(races)

	return output
}

func Part2(input string) int {
	output := 1
	races := formatInput(input, true)
	output = determineOutcomes(races)
	return output
}

func formatInput(str string, reduce bool) []Race {
	inputLines := strings.Split(str, "\n")
	races := []Race{}
	raceTimes := strings.Fields(inputLines[0])[1:]
	raceRecords := strings.Fields(inputLines[1])[1:]

	if reduce {
		raceTime := strings.Join(raceTimes, "")
		raceTimes = []string{raceTime}

		record := strings.Join(raceRecords, "")
		raceRecords = []string{record}
	}

	for i := 0; i < len(raceTimes); i++ {
		raceNum, _ := strconv.Atoi(raceTimes[i])
		recordNum, _ := strconv.Atoi(raceRecords[i])
		races = append(races, Race{RaceTime: raceNum, Record: recordNum})
	}

	return races
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
	fmt.Printf(color.Bold+"Lines: %-*d\tParse: %v\n"+color.Reset, printWidth, len(input), puzzleParseTime.Sub(startTime))
	fmt.Println(color.Bold + color.Blue + "Part 1:" + color.Reset)
	fmt.Printf(color.Bold+"Result: %-*d\tTime: %v\n"+color.Reset, printWidth, result1, part1Time.Sub(puzzleParseTime))
	fmt.Println(color.Bold + color.Blue + "Part 2:" + color.Reset)
	fmt.Printf(color.Bold+"Result: %-*d\tTime: %v\n"+color.Reset, printWidth, result2, part2Time.Sub(part1Time))
	fmt.Println(color.Bold+color.Blue+"Total Time:"+color.Reset, part2Time.Sub(startTime))
	fmt.Print("\n")
}

func determineDistanceTraveled(totalRaceTime int, timeHeld int) int {
	totalTimeTraveled := totalRaceTime - timeHeld
	speedOfTravel := timeHeld
	distanceTraveled := totalTimeTraveled * speedOfTravel
	return distanceTraveled
}

func determineOutcomes(races []Race) int {
	outcome := 1
	for _, race := range races {
		possibleOutcomes := 0
		for heldTime := 0; heldTime < race.RaceTime; heldTime++ {
			distanceTraveled := determineDistanceTraveled(race.RaceTime, heldTime)
			if distanceTraveled > race.Record {
				possibleOutcomes++
			}
		}
		outcome *= possibleOutcomes
	}

	return outcome
}
