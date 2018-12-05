package day4

import (
	"fmt"
	"sort"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "4b",
	Short: "Day 4, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", b(util.ReadInput()))
	},
}

func b(input *util.ChallengeInput) int {
	var events []*event

	for line := range input.Lines() {
		events = append(events, ParseEvent(line))
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].when.Before(events[j].when)
	})

	patterns := map[int]*sleepNumber{}

	lastGuard := 0
	sleepingAt := 0
	for _, event := range events {
		if event.what == ActionBeginShift {
			lastGuard = event.who
		} else {
			var tracker *sleepNumber
			var found bool

			if tracker, found = patterns[lastGuard]; !found {
				tracker = &sleepNumber{total: 1, minutes: [60]int{}}
				patterns[lastGuard] = tracker
			}

			if event.what == ActionFallAsleep {
				sleepingAt = event.when.Minute()
			} else {
				for m := sleepingAt; m < event.when.Minute(); m++ {
					tracker.total++
					tracker.minutes[m]++
				}
			}
		}
	}

	var candidateId int
	minuteId := -1
	candidate := -1

	for k := range patterns {
		if candidate == -1 {
			candidateId = k
			minuteId = 0
			candidate = patterns[k].minutes[0]
		}

		for m := 0; m < 60; m++ {
			if patterns[k].minutes[m] > candidate {
				candidateId = k
				minuteId = m
				candidate = patterns[k].minutes[m]
			}
		}
	}

	return candidateId * minuteId
}
