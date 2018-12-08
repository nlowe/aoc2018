package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "6a",
	Short: "Day 6, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(util.ReadInput()))
	},
}

func intAbs(a int) int {
	return (a + (a >> 31)) ^ (a >> 31)
}

type point struct {
	closest  int
	x        int
	y        int
	infinite bool
}

func parsePoint(line string) *point {
	result := &point{}
	parts := strings.Split(line, ",")

	if x, err := strconv.Atoi(strings.TrimSpace(parts[0])); err != nil {
		panic(err)
	} else {
		result.x = x
	}

	if y, err := strconv.Atoi(strings.TrimSpace(parts[1])); err != nil {
		panic(err)
	} else {
		result.y = y
	}

	return result
}

func (p *point) manhattanDistanceFrom(x, y int) int {
	return intAbs(p.x-x) + intAbs(p.y-y)
}

func closestTo(points []*point, x, y int) (results []*point, closestDistance int) {
	first := true
	for _, p := range points {
		d := p.manhattanDistanceFrom(x, y)
		if first || d <= closestDistance {
			first = false
			if d == closestDistance {
				results = append(results, p)
			} else {
				results = []*point{p}
			}

			closestDistance = d
		}
	}

	return results, closestDistance
}

func a(input *util.ChallengeInput) int {
	var points []*point

	top, left, bottom, right := 0, 0, 0, 0

	for line := range input.Lines() {
		p := parsePoint(line)
		points = append(points, p)

		if len(points) == 1 {
			top, bottom = p.y, p.y
			left, right = p.x, p.x
		} else {
			if p.x < left {
				left = p.x
			} else if p.x > right {
				right = p.x
			}

			if p.y < top {
				top = p.y
			} else if p.y > bottom {
				bottom = p.y
			}
		}
	}

	var infinite []*point
	for x := left; x <= right; x++ {
		infinite, _ = closestTo(points, x, top)
		if len(infinite) == 1 {
			infinite[0].infinite = true
		}

		infinite, _ = closestTo(points, x, bottom)
		if len(infinite) == 1 {
			infinite[0].infinite = true
		}
	}

	for y := top; y <= bottom; y++ {
		infinite, _ = closestTo(points, left, y)
		if len(infinite) == 1 {
			infinite[0].infinite = true
		}

		infinite, _ = closestTo(points, right, y)
		if len(infinite) == 1 {
			infinite[0].infinite = true
		}
	}

	var max *point
	for y := top + 1; y < bottom; y++ {
		for x := left + 1; x < right; x++ {
			candidates, _ := closestTo(points, x, y)

			if len(candidates) == 1 && !candidates[0].infinite {
				candidates[0].closest++
				if max == nil || candidates[0].closest > max.closest {
					max = candidates[0]
				}
			}
		}
	}

	return max.closest
}
