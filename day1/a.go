package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2018/util"
)

var A = &cobra.Command{
	Use:   "1a",
	Short: "Day 1, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(util.ReadInput()))
	},
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
