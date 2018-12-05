package day5

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := util.TestInput(`dabAcCaCBAcCcaDA`)

	require.Equal(t, 10, a(input))
}
