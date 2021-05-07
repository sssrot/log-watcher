package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hpcloud/tail"
)

func main() {
	var confFile string
	var telegramBotToken string
	flag.StringVar(&confFile, "conf", "./conf.json", "-conf=xxx.json")
	flag.Parse()
	telegramBotToken = os.Getenv("TG_TOKEN")

	conf := parseConf(confFile)

	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	pe(err)

	ctx := context.Background()
	for _, fw := range conf.Files {
		f := fw
		go func() {
			f.start(ctx, &conf, bot)
		}()
	}

	c := make(chan struct{})
	<-c
}

func parseConf(path string) Conf {
	b, err := ioutil.ReadFile(path)
	pe(err)
	var conf Conf
	err = json.Unmarshal(b, &conf)
	pe(err)
	return conf
}

type Conf struct {
	ChatID int64          `json:"chat_id,omitempty"`
	Files  []*FileWatcher `json:"files,omitempty"`
}

func (c Conf) SendMessage(bot *tgbotapi.BotAPI, text string) {
	_, err := bot.Send(tgbotapi.NewMessage(c.ChatID, text))
	if err != nil {
		fmt.Println("[ERR] send bot msg err", err)
	}
}

type FileWatcher struct {
	File    string `json:"file,omitempty"`
	Content string `json:"content,omitempty"` //content to watch in line
}

func pe(err error) {
	if err != nil {
		panic(err)
	}
}

func (fw *FileWatcher) start(ctx context.Context, conf *Conf, bot *tgbotapi.BotAPI) {
	t, err := tail.TailFile(fw.File, tail.Config{Follow: true, ReOpen: true})
	pe(err)
	for line := range t.Lines {
		if line.Err != nil {
			continue
		}
		if strings.Contains(line.Text, fw.Content) {
			conf.SendMessage(bot, fmt.Sprintf("find %s in %s, line: %s", fw.Content, fw.File, line.Text))
		}
	}
}
