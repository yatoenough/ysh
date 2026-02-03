package builtins

import "fmt"

type Pwd struct{}

func NewPwd() *Pwd {
	return &Pwd{}
}

func (p *Pwd) Execute(sh ShellContext, args []string) error {
	fmt.Println(sh.GetWorkingDir())
	return nil
}
