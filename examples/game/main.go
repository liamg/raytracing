package main

import (
	"fmt"
	"os"
)

func main() {
	g := Game{}
	if err := g.Init(); err != nil {
		fmt.Printf("Failed to initialise: %s\n", err)
		os.Exit(1)
	}
	defer g.Close()

	level := NewEmptyLevel()

	if err := g.PlayLevel(level); err != nil {
		fmt.Printf("Game crashed: %s\n", err)
		os.Exit(1)
	}
}
