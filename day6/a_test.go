package day6

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := util.TestInput(`1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`)

	require.Equal(t, 17, a(input))
}
