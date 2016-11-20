package main

import (
	"os"
	"path/filepath"

	"github.com/monochromegane/go-gitignore"
)

// IgnoreMatcher is an interface to check whether the target is match or not.
type IgnoreMatcher interface {
	Match(path string, isDir bool) bool
}

// NewIgnoreMatcher return new Matcher instance.
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

	matcher, err := gitignore.NewGitIgnore(gitIgnoreFile)
	if err != nil {
		showError("%v\n", err)
		return nil, err
	}
	return matcher, nil
}
