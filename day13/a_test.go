package day13

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := util.TestInput(`/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/ `)

	require.Equal(t, "7,3", a(input))
}
