package day15

import (
	"fmt"
	"sort"
	"strings"

	"github.com/beefsack/go-astar"
	"github.com/nlowe/aoc2018/util"
)

type World struct {
	Layout [][]*Tile
	Units  []*Unit
	W      int
	H      int

	source []string
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
		dx := t.x + delta.x
		dy := t.y + delta.y

		if dx < 0 || dy < 0 || dx > len(t.w.Layout[0]) || dy > len(t.w.Layout) {
			continue
		}

		if neighbor := t.w.Layout[dx][dy]; neighbor.canOccupy() {
			result = append(result, neighbor)
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
	result := &World{}

	for line := range input.Lines() {
		result.source = append(result.source, line)
	}

	result.Reset()

	return result
}

func (w *World) Reset() {
	w.Layout = [][]*Tile{}

	for x := 0; x < len(w.source[0]); x++ {
		w.Layout = append(w.Layout, []*Tile{})
		for y := 0; y < len(w.source); y++ {
			w.Layout[x] = append(w.Layout[x], nil)
		}
	}

	for y, line := range w.source {
		w.Layout = append(w.Layout, []*Tile{})
		for x, r := range strings.TrimSpace(line) {
			tile := &Tile{x: x, y: y, w: w}

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

				w.Units = append(w.Units, unit)
				tile.occupyingUnit = unit
			}

			w.Layout[x][y] = tile
			if x > w.W {
				w.W = x
			}
		}

		if y > w.H {
			w.H = y
		}
	}
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

func (w *World) Tick() bool {
	sort.Slice(w.Units, w.UnitSortFunc)

	for _, u := range w.Units {
		if u.hp <= 0 {
			continue
		}

		var target *Unit

		// Check if we're within striking distance already
		var openPoints []*Tile
		for _, other := range u.Enemies(w) {
			target = u.pickTarget(target, other)

			if target == nil {
				openPoints = append(openPoints, other.OpenAdjacentTiles(w)...)
			}
		}

		if target == nil {
			var path astar.Pather
			var cost float64

			// If not, find the shortest path that puts us within striking distance
			for _, t := range openPoints {
				// Pick open starting tiles to help the a* out a bit
				for _, start := range u.OpenAdjacentTiles(w) {
					// The path is reversed, with the last element being the source tile
					p, c, found := astar.Path(start, t)
					if found && (path == nil || c < cost) {
						path = p[len(p)-1]
						cost = c
					}
				}
			}

			if path == nil {
				// This unit is blocked by walls and friendly units
				continue
			}

			// Take a step in that direction
			step := path.(*Tile)
			w.Layout[u.x][u.y].occupyingUnit = nil
			step.occupyingUnit = u
			u.x = step.x
			u.y = step.y

			// Check again if we're within striking distance
			for _, other := range u.Enemies(w) {
				target = u.pickTarget(target, other)
			}
		}

		// Do we have someone to attack?
		if target != nil {
			target.hp -= u.attack

			// Remove them from the world if they died
			if target.hp <= 0 {
				w.Layout[target.x][target.y].occupyingUnit = nil

				if !w.AtWar() {
					return false
				}
			}
		}
	}

	return true
}

func (w *World) Survivors() (remainingHP, count int, faction rune) {
	for _, u := range w.Units {
		if u.hp >= 0 {
			count++
			faction = u.faction
			remainingHP += u.hp
		}
	}

	return
}
