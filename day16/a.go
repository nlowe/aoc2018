package day16

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "16a",
	Short: "Day 16, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(util.ReadInput()))
	},
}

func readInput(input *util.ChallengeInput) (snapshots []snapshot, instructions []instruction) {
	c := input.Lines()

	for line := range c {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Before: ") {
			line = line[9 : len(line)-1]
			beforeParts := strings.Split(line, ", ")

			line = strings.TrimSpace(<-c)
			parts := strings.Split(line, " ")

			line = strings.TrimSpace(<-c)
			line = line[9 : len(line)-1]
			afterParts := strings.Split(line, ", ")

			snapshots = append(snapshots, snapshot{
				before: [4]int{util.AtoiOrPanic(beforeParts[0]), util.AtoiOrPanic(beforeParts[1]), util.AtoiOrPanic(beforeParts[2]), util.AtoiOrPanic(beforeParts[3])},
				i: instruction{
					opcode: util.AtoiOrPanic(parts[0]),
					a:      util.AtoiOrPanic(parts[1]),
					b:      util.AtoiOrPanic(parts[2]),
					c:      util.AtoiOrPanic(parts[3]),
				},
				after: [4]int{util.AtoiOrPanic(afterParts[0]), util.AtoiOrPanic(afterParts[1]), util.AtoiOrPanic(afterParts[2]), util.AtoiOrPanic(afterParts[3])},
			})
		} else {
			parts := strings.Split(line, " ")

			instructions = append(instructions, instruction{
				opcode: util.AtoiOrPanic(parts[0]),
				a:      util.AtoiOrPanic(parts[1]),
				b:      util.AtoiOrPanic(parts[2]),
				c:      util.AtoiOrPanic(parts[3]),
			})
		}
	}

	return
}

func a(input *util.ChallengeInput) int {
	c := cpu{}

	snapshots, _ := readInput(input)
	return c.train(snapshots)
}
