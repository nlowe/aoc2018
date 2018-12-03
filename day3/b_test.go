package day3

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := util.TestInput(`#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`)

	require.Equal(t, 3, b(input))
}
