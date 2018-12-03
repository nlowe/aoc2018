package day3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "3a",
	Short: "Day 3, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(util.ReadInput()))
	},
}

func intMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func intMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

type claim struct {
	top    int
	left   int
	width  int
	height int
}

//ID   L,T: WxH
//#1 @ 1,3: 4x4
func parseClaim(line string) *claim {
	parts := strings.Split(line, " ")

	tl := strings.Split(strings.Trim(parts[2], ":"), ",")
	wh := strings.Split(parts[3], "x")

	result := &claim{}

	if l, err := strconv.Atoi(tl[0]); err != nil {
		panic(err)
	} else {
		result.left = l
	}

	if t, err := strconv.Atoi(tl[1]); err != nil {
		panic(err)
	} else {
		result.top = t
	}

	if w, err := strconv.Atoi(wh[0]); err != nil {
		panic(err)
	} else {
		result.width = w
	}

	if h, err := strconv.Atoi(wh[1]); err != nil {
		panic(err)
	} else {
		result.height = h
	}

	return result
}

func (c *claim) area() int {
	return c.width * c.height
}

func (c *claim) intersection(other *claim) *claim {
	t := intMax(c.top, other.top)
	l := intMax(c.left, other.left)

	candidate := &claim{
		top:    t,
		left:   l,
		width:  intMin(c.left+c.width, other.left+other.width) - l,
		height: intMin(c.top+c.height, other.top+other.height) - t,
	}

	if candidate.width > 0 && candidate.height > 0 {
		return candidate
	}

	return nil
}

func a(input *util.ChallengeInput) int {
	var claims []*claim

	for line := range input.Lines() {
		claims = append(claims, parseClaim(line))
	}

	fabric := map[int]map[int]bool{}

	for {
		first := claims[0]
		claims = claims[1:]
		if len(claims) == 1 {
			break
		}

		for _, other := range claims {
			overlap := first.intersection(other)
			if overlap != nil {
				for y := overlap.top; y < overlap.top+overlap.height; y++ {
					for x := overlap.left; x < overlap.left+overlap.width; x++ {
						if _, found := fabric[y]; !found {
							fabric[y] = map[int]bool{}
						}
						fabric[y][x] = true
					}
				}
			}
		}
	}

	overlappingArea := 0
	for _, row := range fabric {
		for range row {
			overlappingArea++
		}
	}

	return overlappingArea
}
