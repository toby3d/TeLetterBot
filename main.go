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

	log.Println("Connected as ", "@"+bot.Self.Username)

	updates := bot.NewLongPollingChannel(nil)

	for update := range updates {
		switch {
		case update.Message != nil:
			log.Printf("%#v", update.Message)
		}
	}
}
