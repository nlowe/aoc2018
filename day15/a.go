package day15

import (
	"fmt"
	"sort"

	"github.com/beefsack/go-astar"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "15a",
	Short: "Day 15, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %d\n", a(util.ReadInput()))
	},
}

func a(input *util.ChallengeInput) int {
	w := LoadWorld(input)
	//fmt.Println("Initially:")
	//DumpWorld()
	//fmt.Println()

	turns := 0

war:
	for w.AtWar() {
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
						break war
					}
				}
			}
		}

		turns++
		//fmt.Printf("After %d round(s):\n", turns)
		//DumpWorld()
		//fmt.Println()
	}

	remainingHP := 0
	for _, u := range w.Units {
		if u.hp >= 0 {
			remainingHP += u.hp
		}
	}

	//fmt.Printf("After %d round(s) with remaining health %d:\n", turns, remainingHP)
	//w.Dump()
	//fmt.Println()
	return turns * remainingHP
}
