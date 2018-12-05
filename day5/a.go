package day5

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "5a",
	Short: "Day 5, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(util.ReadInput()))
	},
}

const delta = 'a' - 'A'

func intAbs(a int32) int32 {
	return (a + (a >> 31)) ^ (a >> 31)
}

func reduce(polymer string) string {
	for {
		buff := strings.Builder{}
		reduced := false

		for i := 0; i < len(polymer)-1; i++ {
			if intAbs(rune(polymer[i])-rune(polymer[i+1])) == delta {
				i++
				reduced = true
			} else {
				buff.WriteRune(rune(polymer[i]))

				if i == len(polymer)-2 {
					buff.WriteRune(rune(polymer[i+1]))
				}
			}
		}

		polymer = buff.String()
		if !reduced {
			break
		}
	}

	return polymer
}

func a(input *util.ChallengeInput) int {
	polymer := <-input.Lines()

	return len(reduce(polymer))
}
