package builtins

import (
	"fmt"
	"strconv"

	"github.com/yatoenough/ysh/internal/history"
)

type History struct {
	hist *history.History
}

func NewHistory(hist *history.History) *History {
	return &History{
		hist,
	}
}

func (h *History) Execute(sh ShellContext, args []string) error {
	cmds := h.hist.Get()
	start := 0

	if len(args) > 1 {
		limit, err := strconv.Atoi(args[1])

		isValidLimit := err == nil && limit > 0 && limit < len(cmds)
		if isValidLimit {
			start = len(cmds) - limit
		}
	}

	for i := start; i < len(cmds); i++ {
		fmt.Printf("\t%d  %s\n", i+1, cmds[i])
	}

	return nil
}
