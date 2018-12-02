package day2

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2018/util"
)

func TestA(t *testing.T) {
	input := util.TestInput(`abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`)

	require.Equal(t, 12, a(input))
}
