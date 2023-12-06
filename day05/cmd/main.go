package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/* Part1 thoughts
Find if in range through source range + range length
Do destination range - source range to get an offset
Add offset to seed while going through each map
???
Get a result
*/

type Range struct {
	dst int
	src int
	rng int
}

func readUntilEmpty(input *bufio.Scanner) []Range {
	mapping := []Range{}

	for i := 0; input.Scan(); i++ {
		line := input.Text()

		if line == "" {
			input.Scan() // skip next label
			break
		}

		fields := strings.Fields(line)
		dst, _ := strconv.Atoi(fields[0])
		src, _ := strconv.Atoi(fields[1])
		rng, _ := strconv.Atoi(fields[2])

		mapping = append(mapping, Range{
			dst: dst,
			src: src,
			rng: rng,
		})
	}

	return mapping
}

// Thank you HFX on The Primeagen discord for hinting me in the right direction <3
func part1(input *bufio.Scanner) {
	input.Scan()
	in := input.Text()
	seedStr := strings.Fields(in[strings.Index(in, ":")+1:])
	seeds := []int{}

	for _, str := range seedStr {
		seed, _ := strconv.Atoi(str)
		seeds = append(seeds, seed)
	}
	fmt.Println(seeds)

	// skip first new line and label
	input.Scan()
	input.Scan()

	// Gather the mappings to their respective namespaces
	// That way we can break early on the seed to avoid re-hitting a range
	seedToSoil := readUntilEmpty(input)
	soilToFertilizer := readUntilEmpty(input)
	fertilizerToWater := readUntilEmpty(input)
	waterToLight := readUntilEmpty(input)
	lightToTemperature := readUntilEmpty(input)
	temperatureToHumidity := readUntilEmpty(input)
	humidityToLocation := readUntilEmpty(input)

	mappings := [][]Range{}
	mappings = append(mappings, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation)

	for i := range seeds {
		for _, mapping := range mappings {
			for _, rng := range mapping {
				if seeds[i] >= rng.src && seeds[i] <= rng.src+rng.rng {
					offset := rng.dst - rng.src
					seeds[i] += offset
					break
				}
			}
		}
	}

	sort.Ints(seeds)
	fmt.Println("Lowest destination", seeds[0])
}

func main() {
	// Got idea from https://github.com/erik-adelbert/aoc/blob/main/2023/day.mk#L51
	input := bufio.NewScanner(os.Stdin)

	// ---- Part 1 ---- //
	part1(input)

	// ---- part 2 ---- //
	// No part 2 due to time constraints :-)
}
