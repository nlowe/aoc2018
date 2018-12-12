package day11

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

var bTests = []struct {
	input    string
	expected string
}{
	{"18", "90,269,16"},
	{"42", "232,251,12"},
}

func TestB(t *testing.T) {
	for _, tt := range bTests {
		t.Run(tt.input, func(t *testing.T) {
			require.Equal(t, tt.expected, b(util.TestInput(tt.input)))
		})
	}
}
