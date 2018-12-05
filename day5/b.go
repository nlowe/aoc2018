package day5

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "5b",
	Short: "Day 5, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(util.ReadInput()))
	},
}

func b(input *util.ChallengeInput) int {
	polymer := <-input.Lines()

	smallest := -1
	for unit := 'A'; unit <= 'Z'; unit++ {
		candidate := len(
			reduce(
				strings.Replace(
					strings.Replace(polymer, string(unit), "", -1),
					string(rune(unit+delta)), "", -1,
				),
			),
		)

		if smallest == -1 || candidate < smallest {
			smallest = candidate
		}
	}

	return smallest
}
