package mgen

import (
	"fmt"
	"math/rand"
)

type Tile struct {
	X, Y    int
	Walls   map[Dir]bool
	Visited bool
	MazeW   int
	MazeH   int
}

func NewTile(x int, y int, mw int, mh int) *Tile {
	walls := make(map[Dir]bool)
	walls[North] = true
	walls[East] = true
	walls[South] = true
	walls[West] = true

	return &Tile{
		X:     x,
		Y:     y,
		Walls: walls,
		MazeW: mw,
		MazeH: mh,
	}
}

func (t *Tile) IsWall(dir Dir) bool {
	if !t.Walls[dir] {
		return false
	}

	// If the wall is at the border, it is not traversable
	if dir == North && t.Y == 0 ||
		dir == East && t.X == t.MazeW-1 ||
		dir == South && t.Y == t.MazeH-1 ||
		dir == West && t.X == 0 {
		return false
	}

	return true
}

func (t *Tile) Draw() string {
	n := "#"
	e := "#"
	s := "#"
	w := "#"
	if !t.Walls[North] {
		n = " "
	}
	if !t.Walls[East] {
		e = " "
	}
	if !t.Walls[South] {
		s = " "
	}
	if !t.Walls[West] {
		w = " "
	}
	str := fmt.Sprintf("#%s#\n", n)
	str += fmt.Sprintf("%s %s\n", w, e)
	str += fmt.Sprintf("#%s#\n", s)
	return str
}

type Dir int

const (
	North Dir = 1
	East  Dir = 2
	South Dir = 3
	West  Dir = 4
)

func RandDir() Dir {
	return Dir(rand.Intn(4) + 1)
}

// Returns opposite direction.
// North <-> South, East <-> West
func OppDir(dir Dir) Dir {
	return Dir((dir+1)%4 + 1)
}
