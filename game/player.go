package game

import (
	"fmt"
	"net"
)

type Player struct {
	Conn         net.Conn
	Name         string
	HighestScore int
}

func NewPlayer(name string, conn net.Conn) *Player {
	return &Player{
		Name:         name,
		HighestScore: 0,
		Conn:         conn,
	}
}

func (p *Player) Close() {
	p.Conn.Close()
}

func (p *Player) Send(data []byte) {
	p.Conn.Write(data)
}

func (p *Player) Receive() []byte {
	data := make([]byte, 1024)

	_, err := p.Conn.Read(data)
	if err != nil {
		fmt.Println("Error reading", err.Error())
	}

	return data
}
