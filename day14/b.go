package day14

import (
	"fmt"
	"strconv"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "14b",
	Short: "Day 14, Problem b",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(util.ReadInput()))
	},
}

func b(input *util.ChallengeInput) int {
	limit := <-input.Lines()

	var limitParts []int
	for _, r := range limit {
		if digit, err := strconv.Atoi(string(r)); err != nil {
			panic(err)
		} else {
			limitParts = append(limitParts, digit)
		}
	}

	a := 0
	b := 1
	score := []int{3, 7}
	matchOffset := 1

	candidate := 0
	for {
		score, a, b = tick(score, a, b)

		for candidate+len(limitParts)-1 < len(score) {
			if score[candidate] == limitParts[0] {
				matchOffset = 0
				for j := candidate + 1; j < candidate+len(limitParts); j++ {
					matchOffset++
					if score[j] != limitParts[matchOffset] {
						break
					}
				}

				if matchOffset == len(limitParts)-1 && score[candidate+matchOffset] == limitParts[matchOffset] {
					return candidate
				}
			}
			candidate++
		}
	}
}
