package day13

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := util.TestInput(`/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`)

	require.Equal(t, "6,4", b(input))
}
