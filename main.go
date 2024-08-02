package main

import (
	"log"

	"github.com/wrongheaven/maze-generator/mgen"
)

func main() {
	generator, err := mgen.New(8, 8)
	check(err)
	_, err = generator.Generate()
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
