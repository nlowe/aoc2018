package day15

// in awe of the size of this lad
type Unit struct {
	x int
	y int

	hp     int
	attack int

	faction rune
}

func (w *World) UnitSortFunc(i, j int) bool {
	u := w.Units[i]
	other := w.Units[j]

	return u.y < other.y || (u.y == other.y && u.x < other.x)
}

func (u *Unit) Enemies(w *World) (result []*Unit) {
	for _, other := range w.Units {
		if other.hp > 0 && u.faction != other.faction {
			result = append(result, other)
		}
	}

	return
}

func (u *Unit) withinStrikingDistance(other *Unit) bool {
	return intAbs(u.x-other.x)+intAbs(u.y-other.y) == 1
}

func (u *Unit) pickTarget(current, other *Unit) *Unit {
	// If we can't even hit the other one don't do anything
	if !u.withinStrikingDistance(other) {
		return current
	}

	// No target so far
	if current == nil {
		return other
	}

	if other.hp < current.hp {
		return other
	}

	if other.hp == current.hp && (other.y < current.y || (other.y == current.y && other.x < current.x)) {
		return other
	}

	return current
}

func (u *Unit) OpenAdjacentTiles(w *World) (result []*Tile) {
	for _, delta := range []struct {
		x int
		y int
	}{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
		if t := w.Layout[u.x+delta.x][u.y+delta.y]; t.canOccupy() {
			result = append(result, t)
		}
	}

	return result
}
