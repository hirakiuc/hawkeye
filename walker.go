package main

import (
	"os"
	"path/filepath"
)

// Walkable interface provide the api to pass as filepath.Walk handler.
type Walkable interface {
	Handler(path string, info os.FileInfo, e error) error
}

// Walker is a filepath.Walk handler.
type Walker struct {
	opts          *Options
	ignoreMatcher IgnoreMatcher
}

// NewWalker create a new Walker instance.
func NewWalker(opts *Options) (Walkable, error) {
	matcher, err := NewIgnoreMatcher(opts)
	if err != nil {
		return nil, err
	}

	return &Walker{
		opts:          opts,
		ignoreMatcher: matcher,
	}, nil
}

// Handler is a function for filepath.Walk handler.
func (walker *Walker) Handler(path string, info os.FileInfo, e error) error {
	if info.IsDir() {
		if err := walker.dirFilter(path, info, e); err != nil {
			return err
		}
	}

	return walker.targetHandler(path, info, e)
}

func (walker *Walker) dirFilter(path string, info os.FileInfo, e error) error {
	if info.Name() == ".git" {
		return filepath.SkipDir
	}
	return nil
}

func (walker *Walker) isIgnoreTarget(rel string, info os.FileInfo) bool {
	if walker.ignoreMatcher == nil {
		return false
	}

	return walker.ignoreMatcher.Match(rel, info.IsDir())
}

func (walker *Walker) targetHandler(path string, info os.FileInfo, e error) error {
	rel, err := filepath.Rel(walker.opts.Path, path)
	if err != nil {
		return err
	}

	if walker.isIgnoreTarget(rel, info) {
		if info.IsDir() {
			return filepath.SkipDir
		}
	}

	if rel == "." {
		return nil
	}

	showMsg("%s\n", rel)
	return nil
}
