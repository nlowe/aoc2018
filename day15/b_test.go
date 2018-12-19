package day15

import (
	"fmt"
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

var bTests = []struct {
	input    string
	expected int
}{
	{`#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`, 4988},
	{`#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`, 31284},
	{`#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`, 3478},
	{`#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`, 6474},
	{`#########   
#G......#   
#.E.#...#
#..##..G#   
#...##..#   
#...#...#   
#.G...G.#   
#.....G.#   
#########`, 1140},
}

func TestB(t *testing.T) {
	for i, tt := range bTests {
		t.Run(fmt.Sprintf("Example %d", i), func(t *testing.T) {
			require.Equal(t, tt.expected, b(util.TestInput(tt.input)))
		})
	}
}
