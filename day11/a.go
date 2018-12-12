package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "11a",
	Short: "Day 11, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %s\n", a(util.ReadInput()))
	},
}

const cellGridSize = 300

func genCells(serial int) [cellGridSize][cellGridSize]int {
	result := [cellGridSize][cellGridSize]int{}

	for y := 0; y < cellGridSize; y++ {
		for x := 0; x < cellGridSize; x++ {
			rackID := x + 10

			result[x][y] = ((rackID * y) + serial) * rackID
			result[x][y] /= 100
			result[x][y] %= 10
			result[x][y] -= 5
		}
	}

	return result
}

func a(input *util.ChallengeInput) string {
	serial, err := strconv.Atoi(strings.TrimSpace(<-input.Lines()))
	if err != nil {
		panic(err)
	}

	var max, maxX, maxY int

	grid := genCells(serial)
	for y := 0; y < (cellGridSize - 3); y++ {
		for x := 0; x < (cellGridSize - 3); x++ {
			sum := grid[x][y] + grid[x+1][y] + grid[x+2][y] +
				grid[x][y+1] + grid[x+1][y+1] + grid[x+2][y+1] +
				grid[x][y+2] + grid[x+1][y+2] + grid[x+2][y+2]

			if sum > max {
				max = sum
				maxX = x
				maxY = y
			}
		}
	}

	return fmt.Sprintf("%d,%d", maxX, maxY)
}
