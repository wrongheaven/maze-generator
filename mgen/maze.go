package mgen

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type Maze struct {
	Width     int
	Height    int
	Tiles     []*Tile
	StartTile *Tile
	EndTile   *Tile
}

func NewMaze(w int, h int) *Maze {
	var tiles []*Tile
	for row := range h {
		for col := range w {
			tiles = append(tiles, NewTile(col, row, w, h))
		}
	}

	return &Maze{
		Width:  w,
		Height: h,
		Tiles:  tiles,
	}
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

func (m *Maze) ForEachTile(fn func(tile *Tile)) {
	for i := range len(m.Tiles) {
		fn(m.Tiles[i])
	}
}

func (m *Maze) PrintToConsole() error {
	fmt.Printf("%s%s\n", strings.Repeat("  ", m.StartTile.X), V)
	for row := range m.Height {
		for col := range m.Width {
			tile, err := m.GetTile(col, row)
			if err != nil {
				return err
			}

			var bools []bool = []bool{
				tile.Walls[North],
				tile.Walls[East],
				tile.Walls[South],
				tile.Walls[West],
			}

			dec := boolSliceToDecimal(bools)
			fmt.Print(decimalToGlyph(dec))
		}
		fmt.Print("\n")
	}
	fmt.Printf("%s%s\n", strings.Repeat("  ", m.EndTile.X), V)

	return nil
}

func decimalToGlyph(dec int) Glyph {
	var glyph Glyph
	switch dec {
	case 0b0001:
		glyph = E + H
	case 0b0010:
		glyph = N + H
	case 0b0100:
		glyph = W + " "
	case 0b1000:
		glyph = S + H
	case 0b0011:
		glyph = NE + H
	case 0b0110:
		glyph = NW + " "
	case 0b1100:
		glyph = SW + " "
	case 0b1001:
		glyph = SE + H
	case 0b0111:
		glyph = N1 + " "
	case 0b1011:
		glyph = E1 + H
	case 0b1101:
		glyph = S1 + " "
	case 0b1110:
		glyph = W1 + " "
	case 0b0101:
		glyph = V + " "
	case 0b1010:
		glyph = H + H
	case 0b0000:
		glyph = P + H
	default:
		glyph = "?"
	}
	return glyph
}

func boolSliceToDecimal(bools []bool) int {
	var decimal int
	for i, b := range bools {
		if b {
			decimal |= 1 << (len(bools) - 1 - i)
		}
	}
	return decimal
}

type Glyph string

const (
	NE Glyph = "╚"
	NW Glyph = "╝"
	SE Glyph = "╔"
	SW Glyph = "╗"
	N  Glyph = "╩"
	S  Glyph = "╦"
	E  Glyph = "╠"
	W  Glyph = "╣"
	H  Glyph = "═"
	V  Glyph = "║"
	P  Glyph = "╬"
	N1 Glyph = "┴"
	E1 Glyph = "├"
	S1 Glyph = "┬"
	W1 Glyph = "┤"
)
