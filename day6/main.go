package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("part1:", part1())
	//fmt.Println("part2:", process(true))
}

func input() []byte {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return input
}

func part1() int {
	line := input()
	start := 0
	for i := 4; i < len(line)-1; i++ {
		if unique(line[start : i+1]) {
			return i
		}
		start++
	}
	return -1
}

func unique(list []byte) bool {
	seen := make(map[byte]struct{}, len(list))
	for i := 0; i < len(list)-1; i++ {
		if _, ok := seen[list[i]]; ok {
			return false
		}
		seen[list[i]] = struct{}{}
	}

	return true
}
