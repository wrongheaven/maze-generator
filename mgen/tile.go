package mgen

import (
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
	// walls = 0b1111

	return &Tile{
		X:     x,
		Y:     y,
		Walls: walls,
		MazeW: mw,
		MazeH: mh,
	}
}

// Converts
func (t *Tile) GetGlyph() Glyph {
	var bits int
	for i := range 4 {
		if t.Walls[Dir(i+1)] {
			bits |= 1 << (3 - i)
		}
	}

	var glyph Glyph
	switch bits {
	case 0b0001:
		glyph = GlyphE
	case 0b0010:
		glyph = GlyphN
	case 0b0100:
		glyph = GlyphW
	case 0b1000:
		glyph = GlyphS
	case 0b0011:
		glyph = GlyphNE
	case 0b0110:
		glyph = GlyphNW
	case 0b1100:
		glyph = GlyphSW
	case 0b1001:
		glyph = GlyphSE
	case 0b0111:
		glyph = GlyphN1
	case 0b1011:
		glyph = GlyphE1
	case 0b1101:
		glyph = GlyphS1
	case 0b1110:
		glyph = GlyphW1
	case 0b0101:
		glyph = GlyphV
	case 0b1010:
		glyph = GlyphH
	case 0b0000:
		glyph = GlyphP
	default:
		glyph = "?"
	}
	return glyph
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

type Glyph string

const (
	GlyphNE Glyph = "╚═"
	GlyphNW Glyph = "╝ "
	GlyphSE Glyph = "╔═"
	GlyphSW Glyph = "╗ "
	GlyphN  Glyph = "╩═"
	GlyphS  Glyph = "╦═"
	GlyphE  Glyph = "╠═"
	GlyphW  Glyph = "╣ "
	GlyphH  Glyph = "══"
	GlyphV  Glyph = "║ "
	GlyphP  Glyph = "╬═"
	GlyphN1 Glyph = "┴ "
	GlyphE1 Glyph = "├═"
	GlyphS1 Glyph = "┬ "
	GlyphW1 Glyph = "┤ "
)
