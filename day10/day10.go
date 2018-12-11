package day10

import (
	"fmt"
	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

// We calculate the solution for 10b during 10a anyways. Since most of the code is re-used,
// they're just setup as aliases of eachOther
var A = &cobra.Command{
	Use:   "10a",
	Aliases: []string{"10b"},
	Short: "Day 10, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: \n\n%s\n", a(util.ReadInput(), SolutionLetterHeight))
	},
}

const TestLetterHeight = 8
const SolutionLetterHeight = 10

type point struct {
	x int
	y int
	vx int
	vy int
}

type bitmap struct {
	pixels map[int]map[int] bool
	minX int
	minY int
	maxX int
	maxY int
}

func parsePoint(line string) *point {
	components := strings.Split(line, "> velocity=<")
	p := strings.Split(strings.TrimRight(components[0][10:], ">"), ",")
	v := strings.Split(strings.TrimRight(components[1], ">"), ",")

	result := &point{}

	if x, err := strconv.Atoi(strings.TrimSpace(p[0])); err != nil {
		panic(err)
	} else {
		result.x = x
	}

	if y, err := strconv.Atoi(strings.TrimSpace(p[1])); err != nil {
		panic(err)
	} else {
		result.y = y
	}

	if vx, err := strconv.Atoi(strings.TrimSpace(v[0])); err != nil {
		panic(err)
	} else {
		result.vx = vx
	}

	if vy, err := strconv.Atoi(strings.TrimSpace(v[1])); err != nil {
		panic(err)
	} else {
		result.vy = vy
	}

	return result
}

// This is the best test I could come up with that works for tests and real inputs
// TODO: Figure out a better test
func isProbableSolution(height int, points []*point) bool {
	bitmap := mapIt(points)

	return bitmap.maxY - bitmap.minY == height - 1
}

func mapIt(points []*point) *bitmap {
	minX := points[0].x
	minY := points[0].y
	maxX := points[0].x
	maxY := points[0].y

	pixels := map[int]map[int]bool{}

	for _, p := range points {
		if _, found := pixels[p.x]; !found {
			pixels[p.x] = map[int]bool{}
		}

		pixels[p.x][p.y] = true

		if p.x < minX {
			minX = p.x
		}

		if p.x > maxX {
			maxX = p.x
		}

		if p.y < minY {
			minY = p.y
		}

		if p.y > maxY {
			maxY = p.y
		}
	}

	return &bitmap{pixels:pixels, minX: minX, minY: minY, maxX: maxX, maxY: maxY}
}

func prettyPrint(points []*point) string {
	bitmap := mapIt(points)


	builder := strings.Builder{}
	for y := bitmap.minY; y <= bitmap.maxY; y++ {
		line := strings.Builder{}

		for x := bitmap.minX; x <= bitmap.maxX; x++ {
			if _, found := bitmap.pixels[x]; found && bitmap.pixels[x][y] {
				line.WriteRune('#')
			} else {
				line.WriteRune('.')
			}
		}

		builder.WriteString(line.String())
		builder.WriteRune('\n')
	}

	return strings.TrimSpace(builder.String())
}

func a(input *util.ChallengeInput, targetHeight int) string {
	var points []*point
	for line := range input.Lines() {
		points = append(points, parsePoint(line))
	}

	steps := 0
	for !isProbableSolution(targetHeight, points) {
		steps += 1
		for _, p := range points {
			p.x += p.vx
			p.y += p.vy
		}
	}

	fmt.Printf("Solution found after %d steps\n", steps)

	return prettyPrint(points)
}
