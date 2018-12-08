package day6

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := util.TestInput(`1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`)

	require.Equal(t, 16, b(input, 32))
}
