package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

func main() {
	path := "Dockerfile"
	f, err := os.Open(filepath.ToSlash(path))
	if err != nil {
		panic(err)
	}

	defer f.Close()

	parsedDockerfile, err := parser.Parse(f)
	if err != nil {
		panic(err)
	}

	for _, child := range parsedDockerfile.AST.Children {
		v, err := instructions.ParseInstruction(child)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%#v\n", v)

	}
}
