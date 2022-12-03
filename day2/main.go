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

func part2() int {
	score := 0
	for _, line := range strings.Split(string(input()), "\n") {
		if line == "" {
			continue
		}
		sp := strings.Split(line, " ")
		opp, outcome := sp[0], sp[1]
		switch outcome {
		case "Y":
			score += playScore(opp) + 3
		case "X":
			score += playScore(howToLose(opp))
		case "Z":
			score += playScore(howToWin(opp)) + 6
		}
	}

	return score
}

func part1() int {
	score := 0
	for _, line := range strings.Split(string(input()), "\n") {
		if line == "" {
			continue
		}
		sp := strings.Split(line, " ")
		opp, me := sp[0], sp[1]
		if draw(opp, me) {
			score += playScore(me) + 3
		} else {
			score += playScore(me)
			if me == howToWin(opp) {
				score += 6
			}
		}
	}

	return score
}

func draw(opp, me string) bool {
	return opp == "A" && me == "X" || opp == "B" && me == "Y" || opp == "C" && me == "Z"
}

func howToWin(opp string) string {
	switch opp {
	case "A":
		return "Y"
	case "B":
		return "Z"
	case "C":
		return "X"
	}

	return ""
}

func howToLose(opp string) string {
	switch opp {
	case "A":
		return "Z"
	case "B":
		return "X"
	case "C":
		return "Y"
	}

	return ""
}

func playScore(play string) int {
	switch play {
	case "X", "A":
		return 1
	case "Y", "B":
		return 2
	case "Z", "C":
		return 3
	}

	return 0
}
