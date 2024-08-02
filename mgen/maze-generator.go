package mgen

import (
	"math/rand"
)

type MazeGenerator struct {
	Maze      *Maze
	Generated bool
}

func NewGenerator() *MazeGenerator {
	return &MazeGenerator{
		Generated: false,
	}
}

func (mg *MazeGenerator) Generate(w int, h int) (*Maze, error) {
	mg.Maze = NewMaze(w, h)
	startX, endX := rand.Intn(w), rand.Intn(w)
	mg.Maze.StartTile = NewTile(startX, -1, w, h)
	mg.Maze.EndTile = NewTile(endX, h, w, h)
	mg.Maze.Tiles = append(mg.Maze.Tiles, mg.Maze.StartTile, mg.Maze.EndTile)

	tileStack := []*Tile{}

	startTile := mg.Maze.GetRandomTile()

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
			currentTile = tileStack[len(tileStack)-2]
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
