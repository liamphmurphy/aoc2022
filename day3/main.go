package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("part1:", part1())
	fmt.Println("part2:", part2())
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

func part2() int {
	sum := 0

	lines := strings.Split(string(input()), "\n")
	priorityRange := 52
	for x := 0; x < len(lines)-3; x += 3 {
		m1 := make(map[int]int, priorityRange)
		m2 := make(map[int]int, priorityRange)
		m3 := make(map[int]int, priorityRange)
		for _, k := range lines[x] {
			m1[priority(string(k))]++
		}

		for _, k := range lines[x+1] {
			m2[priority(string(k))]++
		}
		for _, k := range lines[x+2] {
			m3[priority(string(k))]++
		}

		for i := 0; i <= priorityRange; i++ {
			if _, ok := m2[i]; !ok {
				delete(m1, i)
			}

			if _, ok := m3[i]; !ok {
				delete(m1, i)
			}
		}

		for k := range m1 {
			sum += k
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
