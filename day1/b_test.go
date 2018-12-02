package day1

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/assert"
)

func TestB_ExampleA(t *testing.T) {
	input := util.TestInput(`+1
-1`)

	assert.Equal(t, 0, b(input))
}

func TestB_ExampleB(t *testing.T) {
	input := util.TestInput(`+3
+3
+4
-2
-4`)

	assert.Equal(t, 10, b(input))
}

func TestB_ExampleC(t *testing.T) {
	input := util.TestInput(`-6
+3
+8
+5
-6`)

	assert.Equal(t, 5, b(input))
}

func TestB_ExampleD(t *testing.T) {
	input := util.TestInput(`+7
+7
-2
-7
-4`)

	assert.Equal(t, 14, b(input))
}
