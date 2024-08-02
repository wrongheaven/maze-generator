package mgen

import (
	"errors"
	"fmt"
)

type MazeGenerator struct {
	Maze      *Maze
	Generated bool
}

func New(w int, h int) (*MazeGenerator, error) {
	if w < 4 || h < 4 {
		return nil, errors.New("width and height must be >= 4")
	}

	return &MazeGenerator{
		Maze:      NewMaze(w, h),
		Generated: false,
	}, nil
}

func (mg *MazeGenerator) Generate() (*Maze, error) {
	currentTile := mg.Maze.GetRandomTile()

	fmt.Println(currentTile)

	return mg.Maze, nil
}
