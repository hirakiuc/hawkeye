package main

import (
	"errors"
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Path string `short:"p" long:"path" description:"Target Path" default:"."`
}

func newOptions() (*Options, error) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil, err
	}

	return &Options{
		Path: pwd,
	}, nil
}

func validatePath(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func parseArgs() (*Options, error) {
	opts, err := newOptions()
	if err != nil {
		return nil, err
	}

	parser := flags.NewParser(opts, flags.PrintErrors)
	_, err = parser.Parse()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil, err
	}

	if !validatePath(opts.Path) {
		fmt.Fprintf(os.Stderr, "Invalid Path: %s\n", opts.Path)
		return nil, errors.New("Invalid Path")
	}

	return opts, nil
}
