package day15

import (
	"fmt"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "15b",
	Short: "Day 15, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(util.ReadInput()))
	},
}

func b(input *util.ChallengeInput) int {
	w, elfCount, _ := LoadWorld(input, 3)

	var remainingHP, survivors int
	var faction rune

	turns := 0
search:
	for {
		for w.AtWar() {
			if w.Tick() {
				turns++
			}
		}

		remainingHP, survivors, faction = w.Survivors()
		if faction == elf {
			if survivors == elfCount {
				break search
			} else {
				fmt.Printf("Elves won but lost %d comrads\n", elfCount-survivors)
			}
		} else {
			fmt.Printf("The elves were crushed\n")
		}

		w.ElfAttack++
		w.Reset()
		turns = 0
	}

	fmt.Printf("After %d round(s) with remaining health %d (ATK %d):\n", turns, remainingHP, w.ElfAttack)
	return turns * remainingHP
}
