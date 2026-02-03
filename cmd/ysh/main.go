package main

import (
	"fmt"
	"os"

	"github.com/yatoenough/ysh/internal/repl"
	"github.com/yatoenough/ysh/internal/shell"
)

func main() {
	sh, err := shell.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize shell: %v\n", err)
		os.Exit(1)
	}

	repl.Loop(sh)
}
