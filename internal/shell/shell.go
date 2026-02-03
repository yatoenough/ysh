package shell

import (
	"fmt"
	"os"
	"strings"

	"github.com/yatoenough/ysh/internal/builtins"
	"github.com/yatoenough/ysh/internal/executor"
	"github.com/yatoenough/ysh/internal/history"
)

type Shell struct {
	workingDir string
	builtins   map[string]builtins.Builtin
	executor   *executor.Executor
	history    *history.History
}

func New() (*Shell, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	sh := &Shell{
		workingDir: wd,
		builtins:   make(map[string]builtins.Builtin),
		history:    history.NewHistory(),
	}

	sh.executor = executor.New(sh)

	sh.RegisterBuiltin("echo", builtins.NewEcho())
	sh.RegisterBuiltin("exit", builtins.NewExit())
	sh.RegisterBuiltin("type", builtins.NewType(sh.builtins))
	sh.RegisterBuiltin("pwd", builtins.NewPwd())
	sh.RegisterBuiltin("cd", builtins.NewCd())
	sh.RegisterBuiltin("which", builtins.NewWhich(sh.builtins))
	sh.RegisterBuiltin("history", builtins.NewHistory(sh.history))

	return sh, nil
}

func (sh *Shell) RegisterBuiltin(name string, builtin builtins.Builtin) {
	sh.builtins[name] = builtin
}

func (sh *Shell) Execute(cmdLine string) error {
	args, err := parseCmdLine(cmdLine)

	if err != nil {
		return err
	}

	if len(args) == 0 {
		return nil
	}

	cmdName := args[0]
	sh.history.Append(fmt.Sprintf("%s", strings.Join(args, " ")))

	if builtin, exists := sh.builtins[cmdName]; exists {
		return builtin.Execute(sh, args)
	}

	return sh.executor.Execute(cmdName, args)
}

func (sh *Shell) GetWorkingDir() string {
	return sh.workingDir
}

func (sh *Shell) SetWorkingDir(dir string) {
	sh.workingDir = dir
}

func (sh *Shell) GetBuiltinNames() []string {
	names := make([]string, 0, len(sh.builtins))
	for name := range sh.builtins {
		names = append(names, name)
	}
	return names
}

func (sh *Shell) GetPathExecutables() []string {
	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return nil
	}

	paths := strings.Split(pathEnv, ":")
	executables := make(map[string]bool)

	for _, dir := range paths {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			info, err := entry.Info()
			if err != nil {
				continue
			}

			if info.Mode()&0111 != 0 {
				executables[entry.Name()] = true
			}
		}
	}

	result := make([]string, 0, len(executables))
	for name := range executables {
		result = append(result, name)
	}
	return result
}

func parseCmdLine(cmdLine string) ([]string, error) {
	var args []string
	var current string
	inWord := false

	for _, ch := range cmdLine {
		switch ch {
		case ' ', '\t', '\n':
			if inWord {
				args = append(args, current)
				current = ""
				inWord = false
			}
		default:
			current += string(ch)
			inWord = true
		}
	}

	if inWord {
		args = append(args, current)
	}

	return args, nil
}
