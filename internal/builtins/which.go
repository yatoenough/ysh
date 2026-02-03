package builtins

import (
	"fmt"
	"os/exec"
)

type Which struct {
	builtins map[string]Builtin
}

func NewWhich(builtins map[string]Builtin) *Which {
	return &Which{builtins: builtins}
}

func (t *Which) Execute(sh ShellContext, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("which: missing argument")
	}

	cmdName := args[1]

	if _, exists := t.builtins[cmdName]; exists {
		fmt.Printf("%s: shell built-in command\n", cmdName)
		return nil
	}

	if path, err := exec.LookPath(cmdName); err == nil {
		fmt.Printf("%s\n", path)
		return nil
	}

	fmt.Printf("%s not found\n", cmdName)
	return nil
}
