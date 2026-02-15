package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Execute(args []string, redir Redirect) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if redir.Stdout != "" {
		f, err := os.Create(redir.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "goshell: %s\n", err)
			return
		}
		defer f.Close()
		cmd.Stdout = f
	}

	if redir.Append != "" {
		f, err := os.OpenFile(redir.Append, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "goshell: %s\n", err)
			return
		}
		defer f.Close()
		cmd.Stdout = f
	}

	if redir.Stdin != "" {
		f, err := os.Open(redir.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "goshell: %s\n", err)
			return
		}
		defer f.Close()
		cmd.Stdin = f
	}

	if redir.Stderr != "" {
		f, err := os.Create(redir.Stderr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "goshell: %s\n", err)
			return
		}
		defer f.Close()
		cmd.Stderr = f
	}

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "goshell: %s\n", err)
	}
}

func ExecutePipeline(commands [][]string) {
	var cmds []*exec.Cmd
	var writers []*os.File

	for i, args := range commands {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stderr = os.Stderr

		if i == 0 {
			cmd.Stdin = os.Stdin
		}
		if i == len(commands)-1 {
			cmd.Stdout = os.Stdout
		}

		cmds = append(cmds, cmd)
	}

	for i := 0; i < len(cmds)-1; i++ {
		reader, writer, _ := os.Pipe()
		cmds[i].Stdout = writer
		cmds[i+1].Stdin = reader
		writers = append(writers, writer)
	}

	for _, cmd := range cmds {
		cmd.Start()
	}

	for _, w := range writers {
		w.Close()
	}

	for _, cmd := range cmds {
		cmd.Wait()
	}
}
