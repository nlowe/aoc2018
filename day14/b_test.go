package day14

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/require"
)

var bTests = []struct {
	limit    string
	expected int
}{
	{"51589", 9},
	{"01245", 5},
	{"92510", 18},
	{"59414", 2018},
}

func TestB(t *testing.T) {
	for _, tt := range bTests {
		t.Run(tt.limit, func(t *testing.T) {
			require.Equal(t, tt.expected, b(util.TestInput(tt.limit)))
		})
	}
}
