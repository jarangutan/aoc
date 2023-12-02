package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// "time"
)

type gameLine struct {
	blue  int
	red   int
	green int
}

type game struct {
	id   int
	game []string
}

func parseInput(input []byte) []game {
	games := []game{}
	lines := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")
	for i, line := range lines {
		gameLines := strings.Split(strings.Split(line, ": ")[1], "; ")

		games = append(games, game{
			id:   i + 1,
			game: gameLines,
		})
	}
	return games
}

var bag = map[string]int{
	"red":   12,
	"blue":  14,
	"green": 13,
}

func part1(input []byte) {
	games := parseInput(input)

	total := 0
	for id, g := range games {
		valid := true
		for _, round := range g.game {
			splits := strings.Split(round, ", ")
			for _, s := range splits {
				vals := strings.Split(s, " ")
				value, err := strconv.Atoi(vals[0])
				if err != nil {
					log.Fatal("We didn't get a number in the first part of the split!")
				}

				if vals[1] == "blue" {
					if value > bag["blue"] {
						valid = false
						break
					}
				}
				if vals[1] == "red" {
					if value > bag["red"] {
						valid = false
						break
					}
				}
				if vals[1] == "green" {
					if value > bag["green"] {
						valid = false
						break
					}
				}
			}
		}
		if valid == true {
			fmt.Println("id", id+1)
			total += id + 1
		}
	}

	fmt.Println("valid totals", total)
}

func part2(input []byte) {
	games := parseInput(input)
	success := 0
	for _, game := range games {
		r, g, b := 0, 0, 0
		for _, rounds := range game.game {
			parsedRounds := strings.Split(rounds, ", ")
			for _, item := range parsedRounds {
				values := strings.Split(item, " ")
				value, err := strconv.Atoi(values[0])
				if err != nil {
					log.Fatal("We didn't get a number in the first part of the split!")
				}
				if values[1] == "red" {
					if value > r {
						r = value
					}
				}
				if values[1] == "green" {
					if value > g {
						g = value
					}
				}
				if values[1] == "blue" {
					if value > b {
						b = value
					}
				}
			}
		}
		success += r * g * b
	}
	fmt.Println("value", success)
}

func main() {
	input, err := os.ReadFile("../inputs/input.txt")
	// input, err := os.ReadFile("../inputs/sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	// ---- Part 1 ---- //
	// start1 := time.Now()
	// for i := 0; i < 1000; i++ {
	// 	part1(input)
	// }
	// stop1 := time.Since(start1)
	// fmt.Println("P1 Took", stop1)

	// ---- Part 2 ---- //
	// start2 := time.Now()
	// for i := 0; i < 1000; i++ {
	part2(input)
	// }
	// stop2 := time.Since(start2)
	// fmt.Println("P2 Took", stop2)
}
