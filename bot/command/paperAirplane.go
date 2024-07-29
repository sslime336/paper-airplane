package command

var PaperAirplaneCommandMap = map[string]PaperAirplaneCommand{
	"/ping": Ping,
	"/chat": Chat,
}

type PaperAirplaneCommand int

const (
	Ping PaperAirplaneCommand = iota - 1
	Chat
)
