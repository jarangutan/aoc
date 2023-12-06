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

	for i := 0; input.Scan(); i++ {
		line := input.Text()

		// Skip new line and label
		if line == "" {
			input.Scan()
			input.Scan()
			continue
		}

		fields := strings.Fields(line)
		dst, _ := strconv.Atoi(fields[0])
		src, _ := strconv.Atoi(fields[1])
		rng, _ := strconv.Atoi(fields[2])

		// fmt.Println("fields", dst, src, rng)
		// fmt.Println("offset", dst-src)

		for i, seed := range seeds {
			if seed >= src && seed <= src+rng {
				offset := dst - src
				fmt.Println(offset)
				seeds[i] += offset
				// fmt.Println("Seed after offset", seeds[i])
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
}
