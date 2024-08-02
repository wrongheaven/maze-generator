package mgen

import "math/rand"

type Maze struct {
	Width  int
	Height int
	Tiles  []*Tile
}

func NewMaze(w int, h int) *Maze {
	maze := &Maze{
		Width:  w,
		Height: h,
	}

	var tiles []*Tile
	for row := range h {
		for col := range w {
			tiles = append(tiles, NewTile(col, row))
		}
	}

	maze.Tiles = tiles

	return maze
}

func (m *Maze) GetRandomTile() *Tile {
	walls := make(map[Dir]bool)
	walls[North] = false
	walls[East] = false
	walls[South] = false
	walls[West] = false

	return &Tile{
		X:     rand.Intn(m.Width),
		Y:     rand.Intn(m.Height),
		Walls: walls,
	}
}

type Tile struct {
	X, Y  int
	Walls map[Dir]bool
}

func NewTile(x int, y int) *Tile {
	walls := make(map[Dir]bool)
	walls[North] = false
	walls[East] = false
	walls[South] = false
	walls[West] = false

	return &Tile{
		X:     x,
		Y:     y,
		Walls: walls,
	}
}

type Dir int

const (
	North Dir = 1
	East  Dir = 2
	South Dir = 3
	West  Dir = 4
)
