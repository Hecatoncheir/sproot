package main

import (
	"fmt"
	socket "sproot/socket_engine"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"encoding/json"
)

func main() {
	connection, _, err := websocket.DefaultDialer.Dial("ws://0.0.0.0:8181", nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			inputMessage := socket.MessageEvent{}
			_, messageBytes, err := connection.ReadMessage()

			if err != nil {
				fmt.Fprintf(os.Stdout, "Can't receive message from hecatonhair: %v \n", err)
				fmt.Fprintf(os.Stdout, "Closed connection of hecatonhair \n")
				break
			}

			json.Unmarshal(messageBytes, &inputMessage)

			switch inputMessage.Message {
			case "Item from categories of company parsed":
				fmt.Print(inputMessage)
			}
		}
	}()

	socketEngine := socket.NewEngine("v1.0")
	socketEngine.Listen("0.0.0.0", 8182)
}
