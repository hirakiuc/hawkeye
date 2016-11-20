package main

import (
	"os"
	"path/filepath"
)

type Walkable interface {
	Handler(path string, info os.FileInfo, e error) error
}

type Walker struct {
	opts          *Options
	ignoreMatcher IgnoreMatcher
}

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

func (walker *Walker) Handler(path string, info os.FileInfo, e error) error {
	if info.IsDir() {
		if err := walker.dirFilter(path, info, e); err != nil {
			return err
		}
	}

	return walker.targetHandler(path, info, e)
}

func (walker *Walker) dirFilter(path string, info os.FileInfo, e error) error {
	IgnoreDirs := map[string]bool{
		".git": true,
		"..":   true,
	}

	_, ok := IgnoreDirs[info.Name()]
	if ok {
		return filepath.SkipDir
	} else {
		return nil
	}
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

	showMsg("%s\n", rel)
	return nil
}
