package main

import (
	"fmt"

	"./bot"
	"./config"
)

func main() {

	// Reads the configuration
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Give this bot some life!
	bot.Start()

	// Wait for future input
	<-make(chan struct{})
	return
}
