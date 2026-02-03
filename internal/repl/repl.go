package repl

import (
	"fmt"
	"io"
	"os"

	"github.com/chzyer/readline"
	"github.com/yatoenough/ysh/internal/shell"
)

const colorNone = "\033[0m"

func createCompleter(sh *shell.Shell) *readline.PrefixCompleter {
	builtinNames := sh.GetBuiltinNames()
	pathExecutables := sh.GetPathExecutables()

	items := make([]readline.PrefixCompleterInterface, 0, len(builtinNames)+len(pathExecutables))

	for _, name := range builtinNames {
		items = append(items, readline.PcItem(name))
	}

	for _, name := range pathExecutables {
		items = append(items, readline.PcItem(name))
	}

	return readline.NewPrefixCompleter(items...)
}

func Loop(sh *shell.Shell) {
	completer := createCompleter(sh)

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          colorNone + "$ ",
		HistoryFile:     "",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		AutoComplete:    completer,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create readline: %v\n", err)
		return
	}

	defer rl.Close()

	for {
		cmdLine, err := rl.Readline()

		if err != nil {
			if err == readline.ErrInterrupt || err == io.EOF {
				fmt.Println()
				return
			}
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			continue
		}

		if cmdLine != "" {
			rl.SaveHistory(cmdLine)
		}

		if err := sh.Execute(cmdLine); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}
}
