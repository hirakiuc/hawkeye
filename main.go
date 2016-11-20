package main

import (
	"os"
	"path/filepath"
)

func main() {
	opts, err := parseArgs()
	if err != nil {
		os.Exit(1)
	}

	walker, err := NewWalker(opts)
	if err != nil {
		os.Exit(1)
	}

	err = filepath.Walk(opts.Path, walker.Handler)
	if err != nil {
		showError("%v\n", err)
		os.Exit(1)
	}
}
