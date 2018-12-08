package day8

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := util.TestInput(`2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`)

	require.Equal(t, 66, b(input))
}
