package day11

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

var tests = []struct {
	input    string
	expected string
}{
	{"18", "33,45"},
	{"42", "21,61"},
}

func TestA(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			require.Equal(t, tt.expected, a(util.TestInput(tt.input)))
		})
	}
}
