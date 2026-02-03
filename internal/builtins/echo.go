package builtins

import (
	"fmt"
	"os"
	"strings"
)

type Echo struct{}

func NewEcho() *Echo {
	return &Echo{}
}

func (e *Echo) Execute(sh ShellContext, args []string) error {
	if len(args) <= 1 {
		fmt.Println()
		return nil
	}

	var res strings.Builder

	for _, arg := range args[1:] {
		if len(arg) > 0 && arg[0] == '$' {
			res.WriteString(os.Getenv(arg[1:]))
		} else {
			res.WriteString(arg)
		}

		res.WriteByte(' ')
	}

	fmt.Println(strings.TrimSpace(res.String()))
	return nil
}
