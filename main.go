package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/wrongheaven/maze-generator/mgen"
)

func main() {
	var err error
	var width, height int
	var maze *mgen.Maze

	if len(os.Args) == 1 {
		log.Fatal("missing width and height (ex. 16x9)")
	} else {
		s := strings.Split(os.Args[1], "x")

		if width, err = strconv.Atoi(s[0]); err != nil {
			log.Fatal(err)
		}
		if height, err = strconv.Atoi(s[1]); err != nil {
			log.Fatal(err)
		}
	}

	generator := mgen.NewGenerator()
	if maze, err = generator.Generate(width, height); err != nil {
		log.Fatal(err)
	}

	maze.PrintToConsole()
}
