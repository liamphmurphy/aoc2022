package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("part1:", part1())
	//fmt.Println("part2:", part2())
}

func input() []byte {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return input
}

func part1() int {
	sum := 0
	for _, line := range strings.Split(string(input()), "\n") {
		if line == "" {
			continue
		}

		shared := map[string]struct{}{}
		pivot := len(line) / 2
		split := strings.Split(line, "")
		for i := 0; i < pivot; i++ {
			if _, ok := shared[split[i]]; !ok {
				shared[split[i]] = struct{}{}
			}
		}

		for i := pivot; i < len(split); i++ {
			if _, ok := shared[split[i]]; ok {
				sum += priority(split[i])
				delete(shared, split[i])
			}
		}
	}

	return sum
}

// priority uses unicode values to get 1-26 for a-z, 27-52 for A-Z
func priority(s string) int {
	r := []rune(s)
	upper := false
	// if less than 'a' (unicode of 97), likely an uppercase char
	if r[0] < 'a' {
		upper = true
		r = []rune(strings.ToLower(s))
	}
	modifier := 1
	if upper {
		modifier = 27
	}
	return int(r[0]-'a') + modifier
}
