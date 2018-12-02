package main

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2018/util"
)

func main() {
	fmt.Printf("Answer: %d\n", a(util.ReadInput()))
}

func a(input *util.ChallengeInput) int {
	twos := 0
	threes := 0

	for line := range input.Lines() {
		counts := [26]int{}

		for _, c := range strings.TrimSpace(line) {
			counts[c-'a']++
		}

		two := false
		three := false
		for _, count := range counts {
			if !two && count == 2 {
				two = true
				twos++
			} else if !three && count == 3 {
				three = true
				threes++
			} else if two && three {
				break
			}
		}
	}

	return twos * threes
}
