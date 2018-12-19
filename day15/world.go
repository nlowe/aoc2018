package day15

import (
	"fmt"
	"strings"

	"github.com/beefsack/go-astar"
	"github.com/nlowe/aoc2018/util"
)

type World struct {
	Layout map[int]map[int]*Tile
	Units  []*Unit
	W      int
	H      int
}

const (
	wall   = '#'
	elf    = 'E'
	goblin = 'G'
)

func intAbs(a int) int {
	return (a + (a >> 31)) ^ (a >> 31)
}

type Tile struct {
	x int
	y int

	isWall        bool
	occupyingUnit *Unit

	w *World
}

func (t *Tile) manhattanDistanceFrom(x, y int) int {
	return intAbs(t.x-x) + intAbs(t.y-y)
}

func (t *Tile) canOccupy() bool {
	return !t.isWall && t.occupyingUnit == nil
}

func (t *Tile) PathNeighbors() (result []astar.Pather) {
	for _, delta := range []struct {
		x int
		y int
	}{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		if _, found := t.w.Layout[t.x+delta.x]; found {
			if neighbor, found := t.w.Layout[t.x+delta.x][t.y+delta.y]; found && neighbor.canOccupy() {
				result = append(result, neighbor)
			}
		}
	}

	return
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	tile := to.(*Tile)

	if !tile.canOccupy() {
		return 0.0
	}

	return 1.0
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	tile := to.(*Tile)
	return float64(t.manhattanDistanceFrom(tile.x, tile.y))
}

func LoadWorld(input *util.ChallengeInput) *World {
	result := &World{Layout: map[int]map[int]*Tile{}}

	y := 0
	for line := range input.Lines() {
		for x, r := range strings.TrimSpace(line) {
			if _, found := result.Layout[x]; !found {
				result.Layout[x] = map[int]*Tile{}
			}

			tile := &Tile{x: x, y: y, w: result}

			if r == wall {
				tile.isWall = true
			} else if r == goblin || r == elf {
				unit := &Unit{
					x: x,
					y: y,

					hp:     200,
					attack: 3,

					faction: r,
				}

				result.Units = append(result.Units, unit)
				tile.occupyingUnit = unit
			}

			result.Layout[x][y] = tile
			if x > result.W {
				result.W = x
			}
		}

		if y > result.H {
			result.H = y
		}

		y++
	}

	return result
}

func (w *World) Dump() {
	for y := 0; y <= w.H; y++ {
		line := strings.Builder{}

		for x := 0; x <= w.W; x++ {
			t := w.Layout[x][y]

			if t.isWall {
				line.WriteRune(wall)
			} else if t.occupyingUnit != nil {
				line.WriteRune(t.occupyingUnit.faction)
			} else {
				line.WriteRune('.')
			}
		}

		fmt.Printf("%s\n", line.String())
	}
}

func (w *World) AtWar() bool {
	elvesRemaining := false
	goblinsRemaining := false

	for _, u := range w.Units {
		if u.hp > 0 {
			if u.faction == elf {
				elvesRemaining = true
			} else {
				goblinsRemaining = true
			}
		}
	}

	return elvesRemaining && goblinsRemaining
}
