package day9

import (
	"fmt"
	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "9b",
	Short: "Day 9, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(util.ReadInput()))
	},
}

func b(input *util.ChallengeInput) int {
	r := parseRules(<-input.Lines())
	r.lastMarble *= 100

	return playGame(r)
}
