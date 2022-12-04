package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

// Not the most efficient to use a struct here, but using a struct to easily represent a tuple (and the meaning of the elements) is clearer to me.
type shift struct {
	start int
	end   int
}

func part1() int {
	sum := 0
	for _, line := range strings.Split(string(input()), "\n") {
		if line == "" {
			continue
		}
		shifts := strings.Split(line, ",")
		if contains(processShift(shifts[0]), processShift(shifts[1])) {
			sum++
		}
	}
	return sum
}

func part2() int {
	sum := 0
	for _, line := range strings.Split(string(input()), "\n") {
		if line == "" {
			continue
		}
		shifts := strings.Split(line, ",")
		if overlaps(processShift(shifts[0]), processShift(shifts[1])) {
			sum++
		}
	}
	return sum
}

func contains(one, two shift) bool {
	if one.start <= two.start && one.end >= two.end {
		return true
	} else if two.start <= one.start && two.end >= one.end {
		return true
	}
	return false
}

func overlaps(one, two shift) bool {
	if two.start <= one.end && one.start <= two.start {
		return true
	} else if one.start <= two.end && two.start <= one.start {
		return true
	}
	return false
}

func processShift(s string) shift {
	shifts := strings.Split(s, "-")
	first, _ := strconv.Atoi(shifts[0])
	second, _ := strconv.Atoi(shifts[1])
	return shift{
		start: first,
		end:   second,
	}
}
