package command

var PaperAirplaneCommandMap = map[string]PaperAirplaneCommand{
	"/ping": Ping,
}

type PaperAirplaneCommand int

const (
	Unknown PaperAirplaneCommand = iota - 1
	Ping
)
