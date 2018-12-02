package day2

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := util.TestInput(`abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`)

	require.Equal(t, "fgij", b(input))
}
