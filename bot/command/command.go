package command

import "strings"

type Command string

const (
	Ping                Command = "/ping"
	Chat                Command = "/chat"
)

func Contains(content string, cmd Command) bool {
	return strings.HasPrefix(content, string(cmd))
}
