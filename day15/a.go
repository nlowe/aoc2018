package day15

import (
	"fmt"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "15a",
	Short: "Day 15, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(util.ReadInput()))
	},
}

func a(input *util.ChallengeInput) int {
	w, _, _ := LoadWorld(input, 3)
	turns := 0

	for w.AtWar() {
		if w.Tick() {
			turns++
		}
	}

	remainingHP, _, _ := w.Survivors()

	return turns * remainingHP
}
