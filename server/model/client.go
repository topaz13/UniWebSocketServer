package model

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	ws        *websocket.Conn
	SendCh    chan []byte
	ReceiveCh chan []byte
}

func NewClient(ws *websocket.Conn) *Client {
	c := &Client{
		ws:        ws,
		SendCh:    make(chan []byte, 256),
		ReceiveCh: make(chan []byte, 256),
	}
	c.SendCh <- []byte("Connencted With Server")
	return c
}

func (c *Client) SendLoop() {

	defer func() {
		fmt.Println("Client.SendMessages() goroutine stopping")
		c.Cleanup()
	}()

	for {
		message := <-c.SendCh

		w, err := c.ws.NextWriter(websocket.TextMessage)
		if err != nil {
			fmt.Println("ERROR IN SEND LOOP")
			fmt.Println(err.Error())
			return
		}
		w.Write(message)

		if err := w.Close(); err != nil {
			return
		}
	}
}

func (c *Client) ReadLoop() {
	defer func() {
		fmt.Println("Client.ReadMessage() goroutine stopping")
	}()

	for {
		_, data, err := c.ws.ReadMessage()
		if err != nil {
			println("ERROR IN READ LOOP")
			println(err.Error())
			break
		}

		fmt.Println("READ MESSAGE")
		fmt.Println(data)
		c.ReceiveCh <- data
	}
}

func (cl *Client) Cleanup() {
	close(cl.SendCh)
}
