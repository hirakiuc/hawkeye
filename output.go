package main

import (
	"fmt"
	"os"
)

func showMsg(prefix string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, prefix, a...)
}

func showError(prefix string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, prefix, a...)
}
