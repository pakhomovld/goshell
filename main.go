package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("goshell> ")
		scan := scanner.Scan()
		if !scan {
			break
		}
		text := scanner.Text()

		text = ExpandVars(text)
		tokens := ParseLine(text)
		if len(tokens) == 0 {
			continue
		}
		handled, shouldExit := RunBuiltin(tokens)
		if shouldExit {
			break
		}
		if handled {
			continue
		}
		commands := SplitPipe(tokens)
		if len(commands) == 1 {
			args, redir := ParseRedirects(commands[0])
			Execute(args, redir)
		} else {
			ExecutePipeline(commands)
		}
	}
}
