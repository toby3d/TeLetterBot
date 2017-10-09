package main

import (
	"log"

	"github.com/toby3d/go-telegram"
)

var bot *telegram.Bot

func main() {
	var err error
	bot, err = telegram.NewBot(*flagToken)
	checkError(err)

	log.Printf("Connected as @%s", bot.Self.Username)
}
