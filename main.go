package main

import (
	"fmt"
	"log"

	"github.com/wrongheaven/maze-generator/mgen"
)

func main() {
	generator, err := mgen.New(4, 4)
	check(err)
	maze, err := generator.Generate()
	check(err)

	fmt.Println()
	for _, tile := range maze.Tiles {
		fmt.Println(tile.Draw())
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
