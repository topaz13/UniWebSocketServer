package model

import (
	"fmt"
)

type Lobby struct {
	Clients []Client
}

func NewLobby() *Lobby {
	var lobbt = Lobby{
		Clients: make([]Client, 0),
	}
	return &lobbt
}

func (l *Lobby) Enter(c *Client) {
	l.Clients = append(l.Clients, *c)
	fmt.Println("Client is login")
}

func (l *Lobby) Remove(i int) {
	l.Clients = append(l.Clients[:i], l.Clients[i+1:]...)
}

func (l *Lobby) Run() {
	for {
		for i := 0; i < len(l.Clients); i++ {
			m := <-l.Clients[i].ReceiveCh
			fmt.Println(m)
		}
	}
}
