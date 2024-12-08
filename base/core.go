package base

import (
	"log"
	"strings"
	"time"

	"github.com/traPtitech/traq-ws-bot/payload"
)

func (b *base) BotHandler() {
	log.Println(time.Now())
	b.bot.OnPing(b.Ping)
	b.bot.OnDirectMessageCreated(DMCreated)
}

func DMCreated(p *payload.DirectMessageCreated) {
	quoted := false
	// if there are quotation,ignore space
	sep := strings.FieldsFunc(p.Message.Text, func(r rune) bool {
		if r == '"' {
			quoted = !quoted
		}
		return !quoted && r == ' '
	})

	log.Println(sep) // Foo, bar, random, "letters lol", stuff
}
