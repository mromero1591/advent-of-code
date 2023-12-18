package main

import (
	"fmt"
	"log"
	utils "mromero1591/advent-of-code/utilities"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

var (
	year = "2023"
	day  = "5"
)

func Part1(input string) int {
	output := 0

	steps := []string{"seeds:", "seed-to-soil map:", "soil-to-fertilizer map:", "fertilizer-to-water map:", "water-to-light map:", "light-to-temperature map:", "temperature-to-humidity map:", "humidity-to-location map:"}
	seeds, alamancsMap := convertInput(input, steps)

	locations := []int{}

	steps = steps[1:]
	for _, seed := range seeds {
		value := seed
		for _, step := range steps {
			alamanc := alamancsMap[step]
			for _, a := range alamanc {
				if value >= a.sourceStart && value <= a.sourceEnd {
					sourceDistance := value - a.sourceStart
					value = a.destinationStart + sourceDistance
					break
				}
			}
		}
		locations = append(locations, value)
	}
	output = findSmallest(locations)
	return output
}

func Part2(input string) int {
	output := 0
	steps := []string{"seeds:", "seed-to-soil map:", "soil-to-fertilizer map:", "fertilizer-to-water map:", "water-to-light map:", "light-to-temperature map:", "temperature-to-humidity map:", "humidity-to-location map:"}
	seeds, alamancsMap := convertInputPart2(input, steps)

	steps = steps[1:]
	location := -1

	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		end := seeds[i+1]
		rangeLen := (start + end) - 1
		for j := start; j <= rangeLen; j++ {
			value := j
			for _, step := range steps {
				alamanc := alamancsMap[step]
				for _, a := range alamanc {
					if value >= a.sourceStart && value <= a.sourceEnd {
						sourceDistance := value - a.sourceStart
						value = a.destinationStart + sourceDistance
						break
					}
				}
			}
			if location == -1 || value < location {
				location = value
			}
		}
	}

	output = location
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
	fmt.Printf(color.Bold+"Lines: %-*d\tParse: %v\n"+color.Reset, printWidth, len(input), puzzleParseTime.Sub(startTime))
	fmt.Println(color.Bold + color.Blue + "Part 1:" + color.Reset)
	fmt.Printf(color.Bold+"Result: %-*d\tTime: %v\n"+color.Reset, printWidth, result1, part1Time.Sub(puzzleParseTime))
	fmt.Println(color.Bold + color.Blue + "Part 2:" + color.Reset)
	fmt.Printf(color.Bold+"Result: %-*d\tTime: %v\n"+color.Reset, printWidth, result2, part2Time.Sub(part1Time))
	fmt.Println(color.Bold+color.Blue+"Total Time:"+color.Reset, part2Time.Sub(startTime))
	fmt.Print("\n")
}

type Almanac struct {
	sourceStart      int
	sourceEnd        int
	destinationStart int
	destinationEnd   int
	rangeLength      int
}

func convertInput(str string, steps []string) ([]int, map[string][]Almanac) {
	splitInput := strings.Split(str, "\n")

	seeds := []int{}
	step := steps[0]
	alamancsMap := map[string][]Almanac{}
	for _, str := range splitInput {
		if step == "seeds:" {
			seeds = getSeeds(str)
			step = steps[1]
			continue
		}
		if step == steps[1] {
			if str == "" {
				continue
			}
			if str == steps[2] {
				step = steps[2]
				continue
			}
			if str == steps[1] {
				alamancsMap[steps[1]] = []Almanac{}
				continue
			}
			alamancsMap[steps[1]] = append(alamancsMap[steps[1]], getsourceToDestinationMap(str))
		}
		if step == steps[2] {
			if str == "" {
				continue
			}
			if str == steps[3] {
				step = steps[3]
				continue
			}
			if str == steps[2] {
				alamancsMap[steps[2]] = []Almanac{}
				continue
			}
			alamancsMap[steps[2]] = append(alamancsMap[steps[2]], getsourceToDestinationMap(str))
		}
		if step == steps[3] {
			if str == "" {
				continue
			}
			if str == steps[4] {
				step = steps[4]
				continue
			}
			if str == steps[3] {
				alamancsMap[steps[3]] = []Almanac{}
				continue
			}
			alamancsMap[steps[3]] = append(alamancsMap[steps[3]], getsourceToDestinationMap(str))
		}
		if step == steps[4] {
			if str == "" {
				continue
			}
			if str == steps[5] {
				step = steps[5]
				continue
			}
			if str == steps[4] {
				alamancsMap[steps[4]] = []Almanac{}
				continue
			}
			alamancsMap[steps[4]] = append(alamancsMap[steps[4]], getsourceToDestinationMap(str))
		}
		if step == steps[5] {
			if str == "" {
				continue
			}
			if str == steps[6] {
				step = steps[6]
				continue
			}
			if str == steps[5] {
				alamancsMap[steps[5]] = []Almanac{}
				continue
			}
			alamancsMap[steps[5]] = append(alamancsMap[steps[5]], getsourceToDestinationMap(str))
		}
		if step == steps[6] {
			if str == "" {
				continue
			}
			if str == steps[7] {
				step = steps[7]
				continue
			}
			if str == steps[6] {
				alamancsMap[steps[6]] = []Almanac{}
				continue
			}
			alamancsMap[steps[6]] = append(alamancsMap[steps[6]], getsourceToDestinationMap(str))
		}
		if step == steps[7] {
			if str == "" {
				continue
			}
			if str == steps[7] {
				alamancsMap[steps[7]] = []Almanac{}
				continue
			}
			alamancsMap[steps[7]] = append(alamancsMap[steps[7]], getsourceToDestinationMap(str))
		}
	}
	return seeds, alamancsMap
}
func convertInputPart2(str string, steps []string) ([]int, map[string][]Almanac) {
	splitInput := strings.Split(str, "\n")

	seeds := []int{}
	step := steps[0]
	alamancsMap := map[string][]Almanac{}
	for _, str := range splitInput {
		if step == "seeds:" {
			seeds = getSeedsPart2(str)
			step = steps[1]
			continue
		}
		if step == steps[1] {
			if str == "" {
				continue
			}
			if str == steps[2] {
				step = steps[2]
				continue
			}
			if str == steps[1] {
				alamancsMap[steps[1]] = []Almanac{}
				continue
			}
			alamancsMap[steps[1]] = append(alamancsMap[steps[1]], getsourceToDestinationMap(str))
		}
		if step == steps[2] {
			if str == "" {
				continue
			}
			if str == steps[3] {
				step = steps[3]
				continue
			}
			if str == steps[2] {
				alamancsMap[steps[2]] = []Almanac{}
				continue
			}
			alamancsMap[steps[2]] = append(alamancsMap[steps[2]], getsourceToDestinationMap(str))
		}
		if step == steps[3] {
			if str == "" {
				continue
			}
			if str == steps[4] {
				step = steps[4]
				continue
			}
			if str == steps[3] {
				alamancsMap[steps[3]] = []Almanac{}
				continue
			}
			alamancsMap[steps[3]] = append(alamancsMap[steps[3]], getsourceToDestinationMap(str))
		}
		if step == steps[4] {
			if str == "" {
				continue
			}
			if str == steps[5] {
				step = steps[5]
				continue
			}
			if str == steps[4] {
				alamancsMap[steps[4]] = []Almanac{}
				continue
			}
			alamancsMap[steps[4]] = append(alamancsMap[steps[4]], getsourceToDestinationMap(str))
		}
		if step == steps[5] {
			if str == "" {
				continue
			}
			if str == steps[6] {
				step = steps[6]
				continue
			}
			if str == steps[5] {
				alamancsMap[steps[5]] = []Almanac{}
				continue
			}
			alamancsMap[steps[5]] = append(alamancsMap[steps[5]], getsourceToDestinationMap(str))
		}
		if step == steps[6] {
			if str == "" {
				continue
			}
			if str == steps[7] {
				step = steps[7]
				continue
			}
			if str == steps[6] {
				alamancsMap[steps[6]] = []Almanac{}
				continue
			}
			alamancsMap[steps[6]] = append(alamancsMap[steps[6]], getsourceToDestinationMap(str))
		}
		if step == steps[7] {
			if str == "" {
				continue
			}
			if str == steps[7] {
				alamancsMap[steps[7]] = []Almanac{}
				continue
			}
			alamancsMap[steps[7]] = append(alamancsMap[steps[7]], getsourceToDestinationMap(str))
		}
	}
	return seeds, alamancsMap
}
func getSeeds(str string) []int {
	seeds := []int{}
	valueStr := strings.Trim(str, "")
	values := strings.Split(valueStr, " ")
	values = values[1:]
	for _, v := range values {
		value, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Failed to convert string: \t%s", v)
		}
		seeds = append(seeds, value)
	}
	return seeds
}
func getSeedsPart2(str string) []int {
	seedRanges := []int{}
	valueStr := strings.Trim(str, "")
	values := strings.Split(valueStr, " ")
	values = values[1:]
	for _, v := range values {
		value, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Failed to convert string: \t%s", v)
		}
		seedRanges = append(seedRanges, value)
	}
	// for i := 0; i < len(seedRanges); i += 2 {
	// 	start := seedRanges[i]
	// 	end := seedRanges[i+1]
	// 	rangeLen := (start + end) - 1
	// 	count := 0
	// 	for j := start; j <= rangeLen; j++ {
	// 		if count > 10000 {
	// 			count = 0
	// 			seeds = []int{}
	// 		}
	// 		seeds = append(seeds, j)
	// 		count++
	// 	}
	// }
	// fmt.Printf("seeds created: \t%v", len(seeds))
	return seedRanges
}

func getsourceToDestinationMap(str string) Almanac {
	almanac := Almanac{}
	values := strings.Split(str, " ")
	// fmt.Printf("\nconverting alamanc: \t%v", values)
	destinationStart, _ := strconv.Atoi(values[0])
	sourceStart, _ := strconv.Atoi(values[1])
	rangeLength, _ := strconv.Atoi(values[2])
	almanac.rangeLength = rangeLength
	almanac.sourceStart = sourceStart
	almanac.destinationStart = destinationStart
	almanac.sourceEnd = (sourceStart + almanac.rangeLength) - 1
	almanac.destinationEnd = (destinationStart + almanac.rangeLength) - 1
	return almanac
}

func mapSourceToDesination(sourceValue int, almanac Almanac) int {
	return 0
}

func findSmallest(numbers []int) int {
	if len(numbers) == 0 {
		panic("Cannot find the smallest number in an empty slice")
	}

	smallest := numbers[0]

	for _, num := range numbers {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}
