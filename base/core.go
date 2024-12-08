package base

import (
	"log"
	"time"

	"github.com/traPtitech/traq-ws-bot/payload"
)

func (b *base) BotHandler() {
	log.Println(time.Now())
	b.bot.OnPing(b.Ping)
	b.bot.OnMessageCreated(func(p *payload.MessageCreated) { log.Println("Message created::", p.Message.Text) })
}
