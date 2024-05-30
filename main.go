package main

import (
	"flag"
	"log"
	"main/clients/telegram"
)

const (
	tgBotHost="api.telegam.org"
)

func main() {
	tgClient = telegram.New(tgBotHost, mustToken())
}

func mustToken() string {
	token := flag.String(
		"token-bot-token", 
		"",
		"token for access to telegramm bot",
	)

	flag.Parse()

	if *token == ""{
		log.Fatal("token is not specified")
	}

	return *token
}