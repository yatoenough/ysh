package builtins

import "os"

type Exit struct{}

func NewExit() *Exit {
	return &Exit{}
}

func (e *Exit) Execute(sh ShellContext, args []string) error {
	os.Exit(0)
	return nil
}
