package mgen

import (
	"errors"
	"math/rand"
)

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
			tiles = append(tiles, NewTile(col, row, w, h))
		}
	}

	maze.Tiles = tiles

	return maze
}

func (m *Maze) GetTile(x int, y int) (*Tile, error) {
	for _, tile := range m.Tiles {
		if tile.X == x && tile.Y == y {
			return tile, nil
		}
	}
	return nil, errors.New("tile not found")
}

func (m *Maze) GetRandomTile() *Tile {
	i := rand.Intn(len(m.Tiles))
	return m.Tiles[i]
}

func (m *Maze) GetUnvisitedNeighbors(srcTile *Tile) ([]*Tile, []Dir) {
	nn, _ := m.GetTile(srcTile.X, srcTile.Y-1) // north
	ns, _ := m.GetTile(srcTile.X, srcTile.Y+1) // south
	ne, _ := m.GetTile(srcTile.X+1, srcTile.Y) // east
	nw, _ := m.GetTile(srcTile.X-1, srcTile.Y) // west

	var neighbors []*Tile
	var dirs []Dir
	if nn != nil && !nn.Visited {
		neighbors = append(neighbors, nn)
		dirs = append(dirs, North)
	}
	if ns != nil && !ns.Visited {
		neighbors = append(neighbors, ns)
		dirs = append(dirs, South)
	}
	if ne != nil && !ne.Visited {
		neighbors = append(neighbors, ne)
		dirs = append(dirs, East)
	}
	if nw != nil && !nw.Visited {
		neighbors = append(neighbors, nw)
		dirs = append(dirs, West)
	}

	return neighbors, dirs
}
