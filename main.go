package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {

	boatarde := &tb.Audio{File: tb.FromDisk("assets/boatarde.ogg")}

	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL: "https://api.telegram.org",

		Token:  "1861037011:AAHS_ofMQ_nOz92-HeXQSCKOKTrTten2hbk",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Chat, "Hello World!")
	})

	b.Handle("/sobre", func(m *tb.Message) {
		b.Send(m.Chat, "Bot open source em Go, visite: github.com/jeanlucaslima/tgbot/")
	})

	b.Handle("/boatarde", func(m *tb.Message) {
		b.Send(m.Chat, boatarde)
	})

	b.Start()
}
