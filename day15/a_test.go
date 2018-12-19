package day15

import (
	"fmt"
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

var aTests = []struct {
	input    string
	expected int
}{
	{`#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`, 36334},
	{`#######   
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#   
#######`, 39514},
	{`#######
#E.G#.#
#.#G..#
#G.#.G#   
#G..#.#
#...E.#
#######`, 27755},
	{`#######   
#.E...#   
#.#..G#
#.###.#   
#E#G#G#   
#...#G#
#######`, 28944},
	{`#########   
#G......#
#.E.#...#
#..##..G#
#...##..#   
#...#...#
#.G...G.#   
#.....G.#   
#########`, 18740},
}

func TestA(t *testing.T) {
	for i, tt := range aTests {
		t.Run(fmt.Sprintf("Example %d", i), func(t *testing.T) {
			require.Equal(t, tt.expected, a(util.TestInput(tt.input)))
		})
	}
}
