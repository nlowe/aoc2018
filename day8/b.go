package day8

import (
	"fmt"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "8b",
	Short: "Day 8, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(util.ReadInput()))
	},
}

func (n *node) valueOf() int {
	sum := 0

	if len(n.children) == 0 {
		for _, v := range n.metadata {
			sum += v
		}

		return sum
	}

	for _, i := range n.metadata {
		normalized := i - 1
		if normalized >= 0 && normalized < len(n.children) {
			sum += n.children[normalized].valueOf()
		}
	}

	return sum
}

func b(input *util.ChallengeInput) int {
	head := makeTree(input)

	return head.valueOf()
}
