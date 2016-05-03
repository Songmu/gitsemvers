package main

import (
	"os"

	"github.com/Songmu/gitsemvers"
)

func main() {
	os.Exit((&gitsemvers.CLI{ErrStream: os.Stderr, OutStream: os.Stdout}).Run(os.Args[1:]))
}
