package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/topaz13/UniWebSocketServer/handler"
	"github.com/topaz13/UniWebSocketServer/model"
)

func main() {

	fmt.Println("############### START SERVER ###############")

	e := echo.New()

	lobby := model.NewLobby()
	go lobby.Run()

	e.GET("/ws", handler.NewWebSocketHandler(lobby).Handle)

	e.GET("/ping", func(c echo.Context) error {
		fmt.Println()
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from server")
	})

	e.Start(":8080")
}
