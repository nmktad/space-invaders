package tcp

type Command int

const (
	SHOOT Command = iota + 1
	MOVE
	NEWGAME
	ENDGAME
	CREATEPLAYER
)

type Message struct {
	Data    []byte
	Version float32
	Command Command
}

func NewMessage(data []byte, version float32, command Command) *Message {
	return &Message{
		Data:    data,
		Version: version,
		Command: command,
	}
}
