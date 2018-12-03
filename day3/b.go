package day3

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2018/util"
)

var B = &cobra.Command{
	Use:   "3b",
	Short: "Day 3, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(util.ReadInput()))
	},
}

func b(input *util.ChallengeInput) int {
	var claims []*claim

	for line := range input.Lines() {
		claims = append(claims, parseClaim(line))
	}

search:
	for i := 0; i < len(claims); i++ {
		for j := 0; j < len(claims); j++ {
			if i == j {
				continue
			}

			if claims[i].intersection(claims[j]) != nil {
				continue search
			}
		}

		return claims[i].id
	}

	panic("All claims overlap with at least one other claim")
}
