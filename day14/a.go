package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "14a",
	Short: "Day 14, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %s\n", a(util.ReadInput()))
	},
}

func tick(score []int, a int, b int) ([]int, int, int) {
	combination := score[a] + score[b]
	newA := combination / 10
	newB := combination % 10
	if newA > 0 {
		score = append(score, newA)
	}
	score = append(score, newB)
	return score, (a + 1 + score[a]) % len(score), (b + 1 + score[b]) % len(score)
}

func a(input *util.ChallengeInput) string {
	var limit int
	var err error
	if limit, err = strconv.Atoi(<-input.Lines()); err != nil {
		panic(err)
	}

	a := 0
	b := 1
	score := []int{3, 7}
	for i := 2; i < limit+10; i++ {
		score, a, b = tick(score, a, b)
	}

	result := strings.Builder{}
	for i := limit; i < limit+10; i++ {
		result.WriteString(strconv.Itoa(score[i]))
	}

	return result.String()
}
