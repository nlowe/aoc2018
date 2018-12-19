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
	w := LoadWorld(input)
	//fmt.Println("Initially:")
	//DumpWorld()
	//fmt.Println()

	turns := 0

	for w.AtWar() {
		if w.Tick() {
			turns++
		}
		//fmt.Printf("After %d round(s):\n", turns)
		//DumpWorld()
		//fmt.Println()
	}

	remainingHP, _, _ := w.Survivors()

	//fmt.Printf("After %d round(s) with remaining health %d:\n", turns, remainingHP)
	//w.Dump()
	//fmt.Println()
	return turns * remainingHP
}
