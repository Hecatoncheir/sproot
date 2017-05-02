package main

import (
	"fmt"
	sproot "sproot/socket_engine"
)

func main() {
	socketEngine := sproot.NewEngine("v1.0")
	fmt.Println(socketEngine.APIVersion)
}
