package main

import (
	"log"

	"github.com/wrongheaven/maze-generator/mgen"
)

func main() {
	width, height := 32, 32
	generator := mgen.NewGenerator()
	maze, err := generator.Generate(width, height)
	check(err)

	maze.PrintToConsole()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
