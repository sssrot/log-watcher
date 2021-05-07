package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func TestGetBot(t *testing.T) {
	fmt.Println("hi")
	bot, err := botapi.NewBotAPI(os.Getenv("TOKEN"))
	pe(err)

	// bot.GetChat(botapi.ChatConfig{})

	updates, err := bot.GetUpdates(botapi.NewUpdate(0))
	pe(err)
	fmt.Println("len:", len(updates))
	for _, upd := range updates {
		fmt.Println("", upd.Message.Chat.ID, upd.Message.Text)
	}
}

func TestSendMsg(t *testing.T) {
	cid, err := strconv.ParseInt(os.Getenv("CHATID"), 10, 64)
	pe(err)

	bot, err := botapi.NewBotAPI(os.Getenv("TOKEN"))
	pe(err)

	msg, err := bot.Send(botapi.NewMessage(cid, "from bot"))
	pe(err)
	fmt.Println(msg)
}
