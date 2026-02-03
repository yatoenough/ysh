package builtins

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Cd struct{}

func NewCd() *Cd {
	return &Cd{}
}

func (c *Cd) Execute(sh ShellContext, args []string) error {
	var targetDir string

	if len(args) == 1 {
		home := os.Getenv("HOME")
		if home == "" {
			return fmt.Errorf("cd: HOME not set")
		}
		targetDir = home
	} else {
		targetDir = args[1]
	}

	if targetDir == "~" || strings.HasPrefix(targetDir, "~/") {
		home := os.Getenv("HOME")
		if home == "" {
			return fmt.Errorf("cd: HOME not set")
		}
		if targetDir == "~" {
			targetDir = home
		} else {
			targetDir = filepath.Join(home, targetDir[2:])
		}
	}

	if !filepath.IsAbs(targetDir) {
		targetDir = filepath.Join(sh.GetWorkingDir(), targetDir)
	}

	targetDir = filepath.Clean(targetDir)

	info, err := os.Stat(targetDir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("cd: %s: No such file or directory", args[1])
		}
		return fmt.Errorf("cd: %s: %w", args[1], err)
	}

	if !info.IsDir() {
		return fmt.Errorf("cd: %s: Not a directory", args[1])
	}

	sh.SetWorkingDir(targetDir)
	return nil
}
