package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part1:", process(false))
	fmt.Println("part2:", process(true))
}

func input() []byte {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return input
}

func process(maintainOrder bool) string {
	// If I had more time I would write a little parser to dynamically get the stacks, but we're only going to have these 9...
	// let's call it "premature optimization" by writing the parser ;).
	var s = [][]string{
		{"C", "S", "G", "B"},
		{"G", "V", "N", "J", "H", "W", "M", "T"},
		{"S", "Q", "M"},
		{"M", "N", "W", "T", "L", "S", "B"},
		{"P", "W", "G", "V", "T", "F", "Z", "J"},
		{"S", "H", "Q", "G", "B", "T", "C"},
		{"W", "B", "P", "J", "T"},
		{"M", "Q", "T", "F", "Z", "C", "D", "G"},
		{"F", "P", "B", "H", "S", "N"},
	}

	for _, line := range strings.Split(string(input()), "\n") {
		split := strings.Split(line, " ")
		if split[0] == "move" {
			amt, _ := strconv.Atoi(split[1])
			from, _ := strconv.Atoi(split[3])
			to, _ := strconv.Atoi(split[5])

			move(&s[from-1], &s[to-1], amt, maintainOrder)
		}
	}

	top := make([]string, len(s))
	for i, stack := range s {
		top[i] = stack[0]
	}
	return strings.Join(top, "")
}

func move(src, dst *[]string, amt int, backwards bool) {
	for i := 0; i < amt; i++ {
		*dst = append([]string{pop(src)}, *dst...)
	}

	if backwards {
		*dst = reverse(*dst, 0, amt-1)
	}
}

func reverse(dst []string, start, end int) []string {
	for start < end {
		dst[start], dst[end] = dst[end], dst[start]
		start++
		end--
	}
	return dst
}

func pop(s *[]string) string {
	val := (*s)[0]
	*s = (*s)[1:]
	return val
}
