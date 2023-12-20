package main

import (
	"fmt"
	utils "mromero1591/advent-of-code/utilities"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

var (
	year = "2023"
	day  = "7"
)

type Bid struct {
	Hand   Hand
	Amount int
}

type Hand struct {
	Cards    []Card
	Strength int
}

type Card struct {
	Face     string
	Strength int
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

func Part1(input string) int {
	output := 0
	defaultCardsMap := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}
	inputLines := strings.Split(input, "\n")
	bids := parseInput(inputLines, defaultCardsMap)

	//determine each bids hand strength
	for i, bid := range bids {
		strength := determineHandStrength(bid.Hand)
		bids[i].Hand.Strength = strength
	}

	sortBids(bids)
	for i, bid := range bids {
		output += bid.Amount * (i + 1)
	}
	return output
}

func Part2(input string) int {
	output := 0
	defaultCardsMap := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}

	inputLines := strings.Split(input, "\n")
	bids := parseInput(inputLines, defaultCardsMap)

	//determine each bids hand strength
	for i, bid := range bids {
		strength := determineHandStrength2(bid.Hand)
		bids[i].Hand.Strength = strength
	}

	sortBids(bids)
	for i, bid := range bids {
		output += bid.Amount * (i + 1)
	}

	return output
}

func parseInput(inputLines []string, defaultCardsMap map[string]int) []Bid {
	bids := []Bid{}
	for _, inputLine := range inputLines {
		parts := strings.Fields(inputLine)
		amount, _ := strconv.Atoi(parts[1])
		bid := Bid{
			Amount: amount,
			Hand: Hand{
				Cards: []Card{},
			},
		}

		cards := strings.Split(parts[0], "")
		for _, card := range cards {
			cardValue := defaultCardsMap[card]
			bid.Hand.Cards = append(bid.Hand.Cards, Card{Face: card, Strength: cardValue})
		}

		bids = append(bids, bid)
	}
	return bids
}

func determineHandStrength(hand Hand) int {
	matchedValues := map[string]int{}

	for _, card := range hand.Cards {
		matchedValues[string(card.Face)]++
	}

	pairs, threes, fours, fives := 0, 0, 0, 0
	for _, v := range matchedValues {
		switch v {
		case 2:
			pairs++
		case 3:
			threes++
		case 4:
			fours++
		case 5:
			fives++
		}
	}

	power := 0
	switch {
	case pairs == 1 && threes == 0:
		power = 2
	case pairs == 2:
		power = 3
	case threes == 1 && pairs == 0:
		power = 4
	case threes == 1 && pairs == 1:
		power = 5
	case fours == 1:
		power = 6
	case fives == 1:
		power = 7
	default:
		power = 1
	}

	return power
}

func determineHandStrength2(hand Hand) int {
	matchedValues := map[string]int{}
	jsAvailable := 0
	for _, card := range hand.Cards {
		if card.Strength == 1 {
			jsAvailable++
			continue
		}
		matchedValues[string(card.Face)]++
	}

	pairs, threes, fours, fives := 0, 0, 0, 0
	for _, v := range matchedValues {
		switch v {
		case 2:
			pairs++
		case 3:
			threes++
		case 4:
			fours++
		case 5:
			fives++
		}
	}

	power := 0
	switch {
	case fives == 1,
		fours == 1 && jsAvailable == 1,
		threes == 1 && jsAvailable == 2,
		pairs == 1 && jsAvailable == 3,
		jsAvailable == 5,
		jsAvailable == 4:
		power = 7
	case fours == 1,
		threes == 1 && jsAvailable == 1,
		pairs == 1 && jsAvailable == 2,
		jsAvailable == 3:
		power = 6
	case threes == 1 && pairs == 1,
		pairs == 2 && jsAvailable == 1:
		power = 5
	case threes == 1 && pairs == 0,
		pairs == 1 && jsAvailable == 1,
		pairs == 0 && jsAvailable == 2:
		power = 4
	case pairs == 2:
		power = 3
	case pairs == 1,
		pairs == 0 && jsAvailable == 1:
		power = 2
	default:
		power = 1
	}

	return power
}

func sortBids(bids []Bid) {
	sort.Slice(bids, func(i, j int) bool {
		if bids[i].Hand.Strength != bids[j].Hand.Strength {
			return bids[i].Hand.Strength < bids[j].Hand.Strength
		}

		// If Hand strength is equal, compare Cards
		for k := 0; k < len(bids[i].Hand.Cards); k++ {
			if bids[i].Hand.Cards[k].Strength != bids[j].Hand.Cards[k].Strength {
				return bids[i].Hand.Cards[k].Strength < bids[j].Hand.Cards[k].Strength
			}
		}
		return false
	})
}
