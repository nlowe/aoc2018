package day13

import (
	"fmt"
	"sort"

	"github.com/nlowe/aoc2018/util"
	"github.com/spf13/cobra"
)

var A = &cobra.Command{
	Use:   "13a",
	Short: "Day 13, Problem A",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Answer: %s\n", a(util.ReadInput()))
	},
}

type direction int
type action int

const (
	leftRightTrackPiece    = '-'
	upDownTrackPiece       = '|'
	intersectionTrackPiece = '+'
	aTurnTrackPiece        = '/'
	bTurnTrackPiece        = '\\'

	leftMovingCart      = '<'
	rightMovingCart     = '>'
	upwardsMovingCart   = '^'
	downwardsMovingCart = 'v'
)

const (
	UP direction = iota
	DOWN
	LEFT
	RIGHT
)

const (
	goLeft action = iota
	goStraight
	// implicit: goRight
)

type cart struct {
	x          int
	y          int
	d          direction
	nextAction action
}

func (c *cart) goStraight() {
	switch c.d {
	case UP:
		c.y--
		break
	case RIGHT:
		c.x++
		break
	case DOWN:
		c.y++
		break
	case LEFT:
		c.x--
		break
	}
}

func (c *cart) turnLeft() {
	switch c.d {
	case UP:
		c.x--
		c.d = LEFT
		break
	case RIGHT:
		c.y--
		c.d = UP
		break
	case DOWN:
		c.x++
		c.d = RIGHT
		break
	case LEFT:
		c.y++
		c.d = DOWN
		break
	}
}

func (c *cart) turnRight() {
	switch c.d {
	case UP:
		c.x++
		c.d = RIGHT
		break
	case RIGHT:
		c.y++
		c.d = DOWN
		break
	case DOWN:
		c.x--
		c.d = LEFT
		break
	case LEFT:
		c.y--
		c.d = UP
		break
	}
}

type cartCollection []*cart

func (cs cartCollection) sortFunc(i, j int) bool {
	c := cs[i]
	other := cs[j]

	if c == nil {
		return false
	}

	if other == nil {
		return true
	}

	return c.y < other.y || (c.y == other.y && c.x < other.x)
}

func (cs cartCollection) cartAt(x, y int) *cart {
	for _, c := range cs {
		if c.x == x && c.y == y {
			return c
		}
	}

	return nil
}

func (cs cartCollection) getCollidingCarts() []*cart {
	var result []*cart

	for i := 0; i < len(cs)-1; i++ {
		if cs[i] == nil {
			continue
		}

		for j := i + 1; j < len(cs); j++ {
			if cs[j] == nil {
				continue
			}

			if cs[i].x == cs[j].x && cs[i].y == cs[j].y {
				result = append(result, cs[i], cs[j])
			}
		}
	}

	return result
}

func a(input *util.ChallengeInput) string {
	carts, grid := initGrid(input)

	for {
		cx := -1
		cy := -1
		isFirstCrash := true
		tick(carts, grid, func(collidingCarts cartCollection) bool {
			if isFirstCrash {
				isFirstCrash = false
				cx = collidingCarts[0].x
				cy = collidingCarts[0].y
			}

			return true
		})
		if cx > 0 && cy > 0 {
			return fmt.Sprintf("%d,%d", cx, cy)
		}
	}
}

func initGrid(input *util.ChallengeInput) (cartCollection, [][]rune) {
	var lines []string
	w := 0
	for line := range input.Lines() {
		lines = append(lines, line)
		if len(line) > w {
			w = len(line)
		}
	}
	var carts cartCollection
	grid := make([][]rune, w)
	for i := 0; i < w; i++ {
		grid[i] = make([]rune, len(lines))
	}
	for y, line := range lines {
		for x, r := range line {
			switch r {
			case leftMovingCart:
				grid[x][y] = leftRightTrackPiece
				carts = append(carts, &cart{x: x, y: y, d: LEFT})
				break
			case rightMovingCart:
				grid[x][y] = leftRightTrackPiece
				carts = append(carts, &cart{x: x, y: y, d: RIGHT})
				break
			case upwardsMovingCart:
				grid[x][y] = upDownTrackPiece
				carts = append(carts, &cart{x: x, y: y, d: UP})
				break
			case downwardsMovingCart:
				grid[x][y] = upDownTrackPiece
				carts = append(carts, &cart{x: x, y: y, d: DOWN})
				break
			default:
				grid[x][y] = r
			}
		}
	}
	return carts, grid
}

func tick(carts cartCollection, grid [][]rune, onCollisionCallback func(cartCollection) bool) {
	// sort carts
	sort.Slice(carts, carts.sortFunc)

	for _, cart := range carts {
		if cart == nil {
			continue
		}
		tile := grid[cart.x][cart.y]

		if tile == leftRightTrackPiece || tile == upDownTrackPiece {
			cart.goStraight()
		} else if tile == aTurnTrackPiece {
			// /
			if cart.d == UP {
				cart.turnRight()
			} else if cart.d == RIGHT {
				cart.turnLeft()
			} else if cart.d == DOWN {
				cart.turnRight()
			} else {
				cart.turnLeft()
			}
		} else if tile == bTurnTrackPiece {
			// \
			if cart.d == UP {
				cart.turnLeft()
			} else if cart.d == RIGHT {
				cart.turnRight()
			} else if cart.d == DOWN {
				cart.turnLeft()
			} else {
				cart.turnRight()
			}
		} else if tile == intersectionTrackPiece {
			if cart.nextAction == goLeft {
				cart.turnLeft()
			} else if cart.nextAction == goStraight {
				cart.goStraight()
			} else {
				cart.turnRight()
			}

			cart.nextAction++
			cart.nextAction %= 3
		}

		if onCollisionCallback != nil {
			if cs := carts.getCollidingCarts(); len(cs) > 0 {
				if onCollisionCallback(cs) {
					return
				}
			}
		}
	}
}
