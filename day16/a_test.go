package day16

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := util.TestInput(`Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]`)

	require.Equal(t, 1, a(input))
}
