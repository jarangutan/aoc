package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type card struct {
	id             int
	winningNumbers string
	numbers        string
}

func parseInput(input []byte) []card {
	lines := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")
	cards := []card{}
	for i, line := range lines {
		numbers := strings.Split(strings.Split(line, ": ")[1], " | ")
		cards = append(cards, card{
			id:             i + 1,
			winningNumbers: numbers[0],
			numbers:        numbers[1],
		})
	}
	return cards
}

// Naive approach due to time constraints <|:-)
// Ideally, I'd do recursion but rip time
func part1(cards []card) {
	totalPoints := 0
	for _, card := range cards {
		points := 0
		winningNumbers := strings.Fields(card.winningNumbers)
		numbers := strings.Fields(card.numbers)
		for _, wn := range winningNumbers {
			winNum, err := strconv.Atoi(wn)
			if err != nil {
				log.Fatal("Winning number was not a number some how D:")
			}

			for _, n := range numbers {
				num, err := strconv.Atoi(n)
				if err != nil {
					log.Fatal("Number was not a number some how D:")
				}

				if num == winNum {
					if points == 0 {
						points = 1
					} else {
						points = points * 2
					}
				}
			}
		}
		totalPoints += points
	}
	fmt.Println("totalPoints", totalPoints)
}

func main() {
	input, err := os.ReadFile("../inputs/input.txt")
	// input, err := os.ReadFile("../inputs/sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	cards := parseInput(input)
	// fmt.Printf("Id: %v\nWinNums: %v\nNums: %v", cards[0].id, cards[0].winningNumbers, cards[0].numbers)

	// ---- Part 1 ---- //
	// start1 := time.Now()
	// for i := 0; i < 1000; i++ {
	part1(cards)
	// }
	// stop1 := time.Since(start1)
	// fmt.Println("P1 Took", stop1)

	// ---- Part 2 ---- //
	// start2 := time.Now()
	// for i := 0; i < 1000; i++ {
	// part2(input)
	// }
	// stop2 := time.Since(start2)
	// fmt.Println("P2 Took", stop2)
}
