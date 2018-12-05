package day4

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "4a",
	Short: "Day 4, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(util.ReadInput()))
	},
}

const (
	timestampLayout  = "2006-01-02 15:04"
	ActionBeginShift = "begins shift"
	ActionFallAsleep = "falls asleep"
	ActionWakeUp     = "wakes up"
)

type event struct {
	who  int
	what string
	when time.Time
}

func ParseEvent(line string) *event {
	result := &event{who: -1}
	rawTime, rest := line[1:(1+len(timestampLayout))], strings.TrimSpace(line[(2+len(timestampLayout)):])

	if t, err := time.Parse(timestampLayout, rawTime); err != nil {
		panic(err)
	} else {
		result.when = t
	}

	if rest == ActionFallAsleep || rest == ActionWakeUp {
		result.what = rest
	} else {
		result.what = ActionBeginShift

		idSlug := strings.Split(rest[7:], " ")[0]
		if id, err := strconv.Atoi(idSlug); err != nil {
			panic(err)
		} else {
			result.who = id
		}
	}

	return result
}

type sleepNumber struct {
	total   int
	minutes [60]int
}

func a(input *util.ChallengeInput) int {
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
	var candidate *sleepNumber

	for k, v := range patterns {
		if candidate == nil || v.total > candidate.total {
			candidateId = k
			candidate = v
		}
	}

	possibleMinute := 0
	possibleResult := candidate.minutes[0]
	for m, betterResult := range candidate.minutes[1:] {
		if betterResult > possibleResult {
			possibleMinute = m + 1
			possibleResult = betterResult
		}
	}

	return candidateId * possibleMinute
}
