package day9

import (
	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input string
	expected int
}{
	{"9 players; last marble is worth 25 points", 32},
	{"10 players; last marble is worth 1618 points", 8317},
	{"13 players; last marble is worth 7999 points", 146373},
	{"17 players; last marble is worth 1104 points", 2764},
	{"21 players; last marble is worth 6111 points", 54718},
	{"30 players; last marble is worth 5807 points", 37305},
}

func TestA(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T){
			require.Equal(t, tt.expected, a(util.TestInput(tt.input)))
		})
	}
}
