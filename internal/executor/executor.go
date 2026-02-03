package executor

import (
	"fmt"
	"os"
	"os/exec"
)

type ShellContext interface {
	GetWorkingDir() string
}

type Executor struct {
	shell ShellContext
}

func New(shell ShellContext) *Executor {
	return &Executor{shell: shell}
}

func (e *Executor) Execute(cmdName string, args []string) error {
	path, err := exec.LookPath(cmdName)
	if err != nil {
		return fmt.Errorf("%s: command not found", cmdName)
	}

	cmd := exec.Command(path, args[1:]...)
	cmd.Args[0] = args[0]
	cmd.Dir = e.shell.GetWorkingDir()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmd.Run()

	return nil
}
