package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	crawler "hecatonhair/crawler"
	"log"
	"os"
	"sproot/data"
	socket "sproot/socket_engine"
)

func main() {
	// TODO: add log
	// TODO: add reconnection
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
				parsedItem := crawler.Item{}
				bytes, err := json.Marshal(inputMessage.Data.(map[string]interface{})["Item"])
				if err != nil {
					fmt.Println(err)
				}

				json.Unmarshal(bytes, &parsedItem)
				fmt.Print(parsedItem)

				item, err := data.GetItemByName(parsedItem.Name)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(item.Name)
			}
		}
	}()

	socketEngine := socket.NewEngine("v1.0")
	socketEngine.Listen("0.0.0.0", 8182)
}
