package repl

import (
	"errors"
	"fmt"
	"golite/frontend"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

type typeMetaCommand int

const (
	metaCommandUnrecognized typeMetaCommand = iota
	metaCommandSuccess
)

var completer = readline.NewPrefixCompleter()

func Run() {
	reader, err := readline.NewEx(&readline.Config{
		Prompt:            "\033[34mgolite\u001B[0m \033[32m»\033[0m ",
		HistoryFile:       "/tmp/.golite_repl.tmp",
		HistorySearchFold: true,
		AutoComplete:      completer,
	})
	if err != nil {
		panic(err)
	}

	defer reader.Close()

	for {
		line, err := reader.Readline()
		if err != nil {
			if !errors.Is(err, readline.ErrInterrupt) {
				fmt.Println(err)
			}

			break
		}

		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if line[0] == '.' {
			switch doMetaCommand(line) {
			case metaCommandSuccess:
				break
			case metaCommandUnrecognized:
				fmt.Printf("Unrecognized command '%s'.\n", line)

				continue
			}
		}

		statement := frontend.PrepareStatement(line)
		if statement == nil {
			fmt.Println("Unrecognized keyword at start of " + line)

			continue
		}

		frontend.ExecuteStatement(statement)
		fmt.Println("Executed.")
	}
}

func doMetaCommand(line string) typeMetaCommand {
	switch line {
	case ".exit":
		os.Exit(0)
	}

	return metaCommandUnrecognized
}
