package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var B = &cobra.Command{
	Use:   "11b",
	Short: "Day 11, Problem B",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %s\n", b(util.ReadInput()))
	},
}

func b(input *util.ChallengeInput) string {
	serial, err := strconv.Atoi(strings.TrimSpace(<-input.Lines()))
	if err != nil {
		panic(err)
	}

	var max, maxSize, maxX, maxY int

	grid := genCells(serial)
	for size := 1; size <= cellGridSize; size++ {
		allNegative := true
		for y := 0; y < (cellGridSize - (size - 1)); y++ {
			for x := 0; x < (cellGridSize - (size - 1)); x++ {
				sum := grid[x][y]

				for yy := 1; yy < size; yy++ {
					for xx := 1; xx < size; xx++ {
						sum += grid[x+xx][y+yy]
					}
				}

				if sum > 0 {
					allNegative = false
				}

				if sum > max {
					max = sum
					maxSize = size
					maxX = x
					maxY = y
				}
			}
		}

		// TODO: Optimize this better later
		// After a certain point, the sum stays negative for a while which means we
		// can ***probably*** bail early
		if allNegative {
			break
		}
	}

	return fmt.Sprintf("%d,%d,%d", maxX+1, maxY+1, maxSize-1)
}
