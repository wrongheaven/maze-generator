package mgen

import (
	"fmt"
	"math/rand"
)

type MazeGenerator struct {
	Maze      *Maze
	Generated bool
}

func New(w int, h int) (*MazeGenerator, error) {
	return &MazeGenerator{
		Maze:      NewMaze(w, h),
		Generated: false,
	}, nil
}

func (mg *MazeGenerator) Generate() (*Maze, error) {
	tileStack := []*Tile{}

	startTile := mg.Maze.GetRandomTile()
	fmt.Println(startTile)

	currentTile := startTile
	currentTile.Visited = true
	tileStack = append(tileStack, currentTile)

	for {
		unvisitedNeighbors, dirs := mg.Maze.GetUnvisitedNeighbors(currentTile)
		if len(unvisitedNeighbors) == 0 {
			if currentTile == startTile {
				break
			}

			// backtrack
			currentTile = tileStack[len(tileStack)-1]
			tileStack = tileStack[:len(tileStack)-1]
			continue
		}

		// Pick random neighbor
		r := rand.Intn(len(unvisitedNeighbors))
		otherTile := unvisitedNeighbors[r]

		currentTile.Walls[dirs[r]] = false
		otherTile.Walls[OppDir(dirs[r])] = false

		currentTile = otherTile
		currentTile.Visited = true
		tileStack = append(tileStack, currentTile)
	}

	return mg.Maze, nil
}
