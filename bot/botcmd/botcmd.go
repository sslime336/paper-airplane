package botcmd

var CommandMap = map[string]Command{
	"/ping": Ping,
}

type Command int

const (
	Unknown Command = iota - 1
	Ping
)
