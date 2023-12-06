package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func part1(input []byte) {
	lines := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")

	parsedNums := []string{}
	for _, line := range lines {
		out := ""
		for _, char := range line {
			if unicode.IsNumber(char) {
				out = out + string(char)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			char := line[i]
			if unicode.IsNumber(rune(char)) {
				out = out + string(char)
				break
			}
		}
		parsedNums = append(parsedNums, out)
	}

	total := 0
	for _, s := range parsedNums {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("didn't get a number somehow %v", err)
		}
		total += num
	}
	// fmt.Printf("total %v\n", total)
}

var realNumbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func part2(input []byte) {
	lines := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")

	parsedNums := []string{}
	for _, line := range lines {
		out := ""

		fword := ""
		for _, char := range line {
			if unicode.IsNumber(char) {
				out = out + string(char)
				break
			} else {
				fword = fword + string(char)
				found := false
				for i, value := range realNumbers {
					if strings.Contains(fword, value) {
						out = out + strconv.Itoa(i+1)
						found = true
						break
					}
				}
				if found {
					break
				}
			}
		}

		bword := ""
		for i := len(line) - 1; i >= 0; i-- {
			char := line[i]
			if unicode.IsNumber(rune(char)) {
				out = out + string(char)
				break
			} else {
				bword = string(char) + bword
				found := false
				for i, value := range realNumbers {
					if strings.Contains(bword, value) {
						out = out + strconv.Itoa(i+1)
						found = true
						break
					}
				}
				if found {
					break
				}
			}
		}
		parsedNums = append(parsedNums, out)
	}

	total := 0
	for _, s := range parsedNums {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("didn't get a number somehow %v", err)
		}
		total += num
	}
	// fmt.Printf("total %v\n", total)
}

func main() {
	input, err := os.ReadFile("../inputs/input.txt")
	// input, err := os.ReadFile("../inputs/sample.txt")
	// input, err := os.ReadFile("../inputs/sample2.txt")
	if err != nil {
		log.Fatal(err)
	}

	start1 := time.Now()
	for i := 0; i < 1000; i++ {
		part1(input)
	}
	stop1 := time.Since(start1)
	fmt.Println("P1 Took", stop1)
	fmt.Println("P1 Took", stop1/1000)

	// ---- Part 2 ---- //
	start2 := time.Now()
	for i := 0; i < 1000; i++ {
		part2(input)
	}
	stop2 := time.Since(start2)
	fmt.Println("P2 Took", stop2)
	fmt.Println("P2 Took", stop2/1000)
}
