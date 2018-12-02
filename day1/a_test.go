package main

import (
	"github.com/nlowe/aoc2018/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExampleA(t *testing.T) {
	input := util.TestInput(`+1
+1
+1
`)

	assert.Equal(t, 3, a(input))
}

func TestExampleB(t *testing.T) {
	input := util.TestInput(`+1
+1
-2
`)

	assert.Equal(t, 0, a(input))
}

func TestExampleC(t *testing.T) {
	input := util.TestInput(`-1
-2
-3
`)

	assert.Equal(t, -6, a(input))
}