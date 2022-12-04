package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	split := strings.Split(string(data), "\n")
	max := make([]int, 3)
	current := 0
	for _, item := range split {
		if item == "" {
			for i := range max {
				if current > max[i] {
					old := max[i]
					max[i] = current
					if i+1 < len(max) {
						max[i+1] = old
					}
					break
				}
			}
			current = 0
		} else {
			cur, _ := strconv.Atoi(item)
			current += cur
		}
	}

	total := 0
	for _, m := range max {
		total += m
	}
	fmt.Println(total)
}
