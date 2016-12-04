package main

import git "github.com/libgit2/git2go"

// IgnoreMatcher is an interface to check whether the target is match or not.
type IgnoreMatcher interface {
	Match(path string, isDir bool) bool
}

type Matcher struct {
	repo *git.Repository
}

// NewIgnoreMatcher return new Matcher instance.
func NewIgnoreMatcher(opts *Options) (IgnoreMatcher, error) {
	repo, err := git.OpenRepository(opts.Path)
	if err != nil {
		showError("%v\n", err)
		return nil, err
	}

	return &Matcher{
		repo,
	}, nil
}

func (matcher *Matcher) Match(path string, isDir bool) bool {
	if matcher.repo == nil {
		return false
	}

	ignored, err := matcher.repo.IsPathIgnored(path)
	if err != nil {
		showError("%v\n", err)
		return false
	}

	return ignored
}
