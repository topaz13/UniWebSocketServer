package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/topaz13/UniWebSocketServer/model"
)

type Data struct {
	Message string `json:"message"`
}

type WebScoketHandler struct {
	lobby *model.Lobby
}

func NewWebSocketHandler(lobby *model.Lobby) *WebScoketHandler {
	return &WebScoketHandler{
		lobby: lobby,
	}
}

func (handler *WebScoketHandler) Handle(c echo.Context) error {

	fmt.Println("[UniWebSocketLog] New Cloent")
	upgrader := &websocket.Upgrader{}

	headers := c.Request().Header

	fmt.Println("[UniWebSocketLog] Request.Headers")
	for key, value := range headers {
		fmt.Println(key, value)
	}

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		log.Fatal(err)
		return err
	}

	client := model.NewClient(ws)
	fmt.Println("KOko")
	handler.lobby.Enter(client)

	go client.ReadLoop()
	go client.SendLoop()

	return nil
}

func ReadMessage(ws *websocket.Conn) {
	for {
		// jsonDataには一文字ずつ入っている。
		_, jsonData, err := ws.ReadMessage()
		if err != nil {
			println(err)
			println("EEEEEERRRRRR")
			break
		}
		println("READ DATA")
		println(jsonData)

		println(len(jsonData))

		var d Data
		if err := json.Unmarshal(jsonData, &d); err != nil {
			println(err)
			continue
		}

		fmt.Printf("%+v\n", d) //=> {name:まく age:14}
		fmt.Println(d.Message) //=> まく
		bytes, _ := json.Marshal(d)
		fmt.Print(string(bytes))
	}
}
