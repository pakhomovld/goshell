package main

import (
	"fmt"
	"os"
	"strings"
)

func RunBuiltin(args []string) (bool, bool) {
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			os.Chdir(os.Getenv("HOME"))
			return true, false
		} else {
			os.Chdir(args[1])
			return true, false
		}
	case "pwd":
		dir, _ := os.Getwd()
		fmt.Println(dir)
		return true, false
	case "exit":
		return true, true
	case "export":
		if len(args) < 2 {
			for _, e := range os.Environ() {
				fmt.Println(e)
			}
		} else {
			parts := strings.SplitN(args[1], "=", 2)
			os.Setenv(parts[0], parts[1])
		}
		return true, false
	case "unset":
		if len(args) < 2 {
			fmt.Println("unset: not enough arguments")
		} else {
			os.Unsetenv(args[1])
		}
		return true, false
	case "env":
		for _, e := range os.Environ() {
			fmt.Println(e)
		}
		return true, false
	default:
		return false, false
	}
}
