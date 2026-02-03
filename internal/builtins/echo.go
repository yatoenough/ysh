package builtins

import (
	"fmt"
	"strings"
)

type Echo struct{}

func NewEcho() *Echo {
	return &Echo{}
}

func (e *Echo) Execute(sh ShellContext, args []string) error {
	msg := strings.Join(args[1:], " ")
	fmt.Println(msg)
	return nil
}
