package day9

import (
	"fmt"
	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var A = &cobra.Command{
	Use:   "9a",
	Short: "Day 9, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(util.ReadInput()))
	},
}

type rules struct {
	players int
	lastMarble int
}

func parseRules(line string) *rules {
	result := &rules{}
	parts := strings.Split(line, " ")

	if players, err := strconv.Atoi(parts[0]); err != nil {
		panic(err)
	} else {
		result.players = players
	}

	if lastMarble, err := strconv.Atoi(parts[6]); err != nil {
		panic(err)
	} else {
		result.lastMarble = lastMarble
	}

	return result
}

type node struct {
	value int
	previous *node
	next *node
}

func a(input *util.ChallengeInput) int {
	r := parseRules(<-input.Lines())
	score := make([]int, r.players)

	head := &node{value: 0}
	head.previous = head
	head.next = head

	highestScore := 0
	currentPlayer := 0
	for m := 1; m <= r.lastMarble; m++ {
		if m % 23 == 0 {
			score[currentPlayer] += m
			for i := 0; i < 7; i++ {
				head = head.previous
			}

			score[currentPlayer] += head.value

			left := head.previous
			right := head.next

			left.next = right
			right.previous = left

			head = right

			if score[currentPlayer] > highestScore {
				highestScore = score[currentPlayer]
			}
		} else {
			head = head.next
			ins := &node{
				value: m,
				previous: head,
				next: head.next,
			}

			head.next.previous = ins
			head.next = ins
			head = ins
		}


		currentPlayer++
		currentPlayer %= r.players
	}

	return highestScore
}
