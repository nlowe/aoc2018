package day14

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

var aTests = []struct {
	limit    string
	expected string
}{
	{"9", "5158916779"},
	{"5", "0124515891"},
	{"18", "9251071085"},
	{"2018", "5941429882"},
}

func TestA(t *testing.T) {
	for _, tt := range aTests {
		t.Run(tt.limit, func(t *testing.T) {
			require.Equal(t, tt.expected, a(util.TestInput(tt.limit)))
		})
	}
}
