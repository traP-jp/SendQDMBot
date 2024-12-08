package base

import (
	"log"
	"os"

	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

type base struct {
	botid string
	bot   *traqwsbot.Bot
}

func Setup(botttoken string) {
	base := newBot(botttoken)
	log.Println("BotSetup")
	go base.BotHandler()
	base.bot.Start()
}

func newBot(bottoken string) *base {
	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: bottoken, // Required
	})
	if err != nil {
		log.Fatalf("Error initialising bot: %v", err)
	}
	botid := os.Getenv("TRAQ_BOT_ID")

	return &base{botid: botid, bot: bot}
}
