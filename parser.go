package main

import (
	"os"
	"strings"
)

func ParseLine(line string) []string {
	fields := strings.Fields(line)

	return fields
}

func isVarChar(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9') ||
		b == '_'
}

func ExpandVars(line string) string {
	var result strings.Builder
	i := 0

	for i < len(line) {
		if line[i] == '$' && i+1 < len(line) {
			i++
			start := i
			for i < len(line) && isVarChar(line[i]) {
				i++
			}
			name := line[start:i]
			result.WriteString(os.Getenv(name))
		} else {
			result.WriteByte(line[i])
			i++
		}
	}
	return result.String()
}

type Redirect struct {
	Stdout string // file for ">"
	Append string // file for ">>"
	Stdin  string // file for "<"
	Stderr string // file for "2>"
}

func ParseRedirects(tokens []string) ([]string, Redirect) {
	var args []string
	var redir Redirect
	i := 0

	for i < len(tokens) {
		switch tokens[i] {
		case ">":
			redir.Stdout = tokens[i+1]
			i += 2
		case ">>":
			redir.Append = tokens[i+1]
			i += 2
		case "<":
			redir.Stdin = tokens[i+1]
			i += 2
		case "2>":
			redir.Stderr = tokens[i+1]
			i += 2
		default:
			args = append(args, tokens[i])
			i++
		}
	}
	return args, redir
}

func SplitPipe(tokens []string) [][]string {
	var result [][]string
	var current []string

	for _, token := range tokens {
		if token == "|" {
			result = append(result, current)
			current = nil
		} else {
			current = append(current, token)

		}
	}
	result = append(result, current)
	return result
}
