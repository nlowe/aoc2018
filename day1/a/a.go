package main

import (
	"fmt"
	"github.com/nlowe/aoc2018/util"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf(	"Answer: %d", a(util.ReadInput()))
}

func a(input *util.ChallengeInput) int {
	frequency := 0

	for line := range input.Lines() {
		rawDelta := strings.TrimSpace(line)

		if rawDelta[0] == '+' {
			rawDelta = rawDelta[1:]
		}

		if delta, err := strconv.Atoi(rawDelta); err != nil {
			panic(err)
		} else {
			frequency += delta
		}
	}

	return frequency
}