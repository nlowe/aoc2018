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
	rounds  = 20
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

func a(input *util.ChallengeInput) int {
	state := strings.TrimSpace((<-input.Lines())[len("initial state: "):])

	rules := map[string]rune{}
	for line := range input.Lines() {
		raw := strings.TrimSpace(line)
		if raw == "" || strings.HasSuffix(raw, ".") {
			continue
		}

		rules[raw[0:5]] = '#'
	}

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
	}

	return score(state, offset)
}
