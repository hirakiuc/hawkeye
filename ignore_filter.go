package main

import (
	"os"
	"path/filepath"

	"github.com/monochromegane/go-gitignore"
)

type IgnoreMatcher interface {
	Match(path string, isDir bool) bool
}

func NewIgnoreMatcher(opts *Options) (IgnoreMatcher, error) {
	return getIgnoreMatcher(opts.Path)
}

func getIgnoreMatcher(path string) (gitignore.IgnoreMatcher, error) {
	gitIgnoreFile, err := filepath.Rel(path, ".gitignore")
	if err != nil {
		showError("%v\n", err)
		return nil, err
	}

	if _, err := os.Stat(gitIgnoreFile); err != nil {
		return nil, nil
	}

	if matcher, err := gitignore.NewGitIgnore(gitIgnoreFile); err != nil {
		showError("%v\n", err)
		return nil, err
	} else {
		return matcher, nil
	}
}
