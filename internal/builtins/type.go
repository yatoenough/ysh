package builtins

import (
	"fmt"
	"os/exec"
)

type Type struct {
	builtins map[string]Builtin
}

func NewType(builtins map[string]Builtin) *Type {
	return &Type{builtins: builtins}
}

func (t *Type) Execute(sh ShellContext, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("type: missing argument")
	}

	cmdName := args[1]

	if _, exists := t.builtins[cmdName]; exists {
		fmt.Printf("%s is a shell builtin\n", cmdName)
		return nil
	}

	if path, err := exec.LookPath(cmdName); err == nil {
		fmt.Printf("%s is %s\n", cmdName, path)
		return nil
	}

	fmt.Printf("%s: not found\n", cmdName)
	return nil
}
