package day7

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := util.TestInput(`Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`)

	require.Equal(t, 15, b(input, 2, 0))
}
