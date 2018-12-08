package day6

import (
	"fmt"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "6b",
	Short: "Day 6, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(util.ReadInput(), 10000))
	},
}

func b(input *util.ChallengeInput, threshold int) int {
	var points []*point

	top, left, bottom, right := 0, 0, 0, 0

	for line := range input.Lines() {
		p := parsePoint(line)
		points = append(points, p)

		if len(points) == 1 {
			top, bottom = p.y, p.y
			left, right = p.x, p.x
		} else {
			if p.x < left {
				left = p.x
			} else if p.x > right {
				right = p.x
			}

			if p.y < top {
				top = p.y
			} else if p.y > bottom {
				bottom = p.y
			}
		}
	}

	size := 0
	for y := top + 1; y < bottom; y++ {
		for x := left + 1; x < right; x++ {
			sum := 0
			for _, p := range points {
				sum += p.manhattanDistanceFrom(x, y)
			}

			if sum < threshold {
				size++
			}
		}
	}

	return size
}
