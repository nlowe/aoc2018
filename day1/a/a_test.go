package main

import (
	"testing"

	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/assert"
)

func TestA_ExampleA(t *testing.T) {
	input := util.TestInput(`+1
+1
+1
`)

	assert.Equal(t, 3, a(input))
}

func TestA_ExampleB(t *testing.T) {
	input := util.TestInput(`+1
+1
-2
`)

	assert.Equal(t, 0, a(input))
}

func TestA_ExampleC(t *testing.T) {
	input := util.TestInput(`-1
-2
-3
`)

	assert.Equal(t, -6, a(input))
}
