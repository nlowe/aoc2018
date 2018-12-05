package day5

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := util.TestInput(`dabAcCaCBAcCcaDA`)

	require.Equal(t, 4, b(input))
}
