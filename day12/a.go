package day12

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "12a",
	Short: "Day 12, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(util.ReadInput()))
	},
}

const (
	padding = "....."
)

func score(state string, start int) int {
	score := 0
	for i := 0; i < len(state); i++ {
		if state[i] == '#' {
			score += start + i
		}
	}

	return score
}

func simulate(input *util.ChallengeInput, rounds int) int {
	state := strings.TrimSpace((<-input.Lines())[len("initial state: "):])

	rules := map[string]rune{}
	for line := range input.Lines() {
		raw := strings.TrimSpace(line)
		if raw == "" || strings.HasSuffix(raw, ".") {
			continue
		}

		rules[raw[0:5]] = '#'
	}

	stable := 2
	lastScore := score(state, 0)
	lastDelta := 0
	offset := 0
	for i := 0; i < rounds; i++ {
		// Normalize ends
		state = strings.TrimRight(state, ".")
		if strings.HasPrefix(state, ".") {
			firstLiving := strings.IndexRune(state, '#')
			state = state[firstLiving:]
			offset += firstLiving
		}

		// Pad
		state = padding + state + padding
		offset -= len(padding)

		nextGeneration := strings.Builder{}
		nextGeneration.WriteString("..")

		for i := 2; i < len(state)-2; i++ {
			slice := state[(i - 2):(i + 3)]

			if _, found := rules[slice]; found {
				nextGeneration.WriteRune('#')
			} else {
				nextGeneration.WriteRune('.')
			}
		}

		state = nextGeneration.String()

		currentScore := score(state, offset)
		currentDelta := currentScore - lastScore
		if currentDelta == lastDelta {
			stable--
			if stable == 0 {
				fmt.Printf("Stabilized after %d iterations to +%d\n", i, currentDelta)

				return currentScore + ((rounds - i - 1) * currentDelta)
			}
		} else {
			lastDelta = currentDelta
			lastScore = currentScore
		}
	}

	return score(state, offset)
}

func a(input *util.ChallengeInput) int {
	return simulate(input, 20)
}
