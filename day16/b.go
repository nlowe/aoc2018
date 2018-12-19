package day16

import (
	"fmt"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "16b",
	Short: "Day 16, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(util.ReadInput()))
	},
}

func b(input *util.ChallengeInput) int {
	c := cpu{}

	snapshots, program := readInput(input)
	c.train(snapshots)

	c.program = program
	c.run()

	return c.registers[0]
}
