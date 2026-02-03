package builtins

type Builtin interface {
	Execute(sh ShellContext, args []string) error
}

type ShellContext interface {
	GetWorkingDir() string
	SetWorkingDir(dir string)
}
