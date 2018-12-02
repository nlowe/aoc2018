package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2018/util"
)

var B = &cobra.Command{
	Use:   "1b",
	Short: "Day 1, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(util.ReadInput()))
	},
}

func b(input *util.ChallengeInput) int {
	frequency := 0
	tracker := map[int]bool{}
	tracker[0] = true

	var list []int
	for line := range input.Lines() {
		rawDelta := strings.TrimSpace(line)

		if rawDelta[0] == '+' {
			rawDelta = rawDelta[1:]
		}

		if delta, err := strconv.Atoi(rawDelta); err != nil {
			panic(err)
		} else {
			list = append(list, delta)
		}
	}

	for {
		for _, delta := range list {
			frequency += delta

			if tracker[frequency] {
				return frequency
			} else {
				tracker[frequency] = true
			}
		}
	}
}
