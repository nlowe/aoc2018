package day13

import (
	"fmt"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "13b",
	Short: "Day 13, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %s\n", b(util.ReadInput()))
	},
}

func b(input *util.ChallengeInput) string {
	carts, grid := initGrid(input)

	remainingCarts := len(carts)
	for remainingCarts > 1 {
		tick(carts, grid, func(cs cartCollection) bool {
			for _, c := range cs {
				for i, other := range carts {
					if c == other {
						carts[i] = nil
						remainingCarts--
					}
				}
			}

			return false
		})
	}

	var x, y int
	for _, c := range carts {
		if c != nil {
			x = c.x
			y = c.y
		}
	}

	return fmt.Sprintf("%d,%d", x, y)
}
