package util

import (
	"os"
	"strings"
)

func BodyFrom(args []string) string {
	if len(args) < 3 || os.Args[2] == "" {
		return "hello"
	}

	return strings.Join(args[2:], " ")
}

func SeverityFrom(args []string) string {
	if len(args) < 2 || os.Args[1] == "" {
		return "info"
	}

	return os.Args[1]
}
