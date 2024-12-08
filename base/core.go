package base

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/traPtitech/traq-ws-bot/payload"
)

func (b *base) BotHandler() {
	log.Println(time.Now())
	b.bot.OnPing(b.Ping)
	b.bot.OnDirectMessageCreated(b.DMCreated)
}

func (b *base) DMCreated(p *payload.DirectMessageCreated) {
	quoted := false
	if b.OnGroupExists(p.Message.User.ID, "ba6552f8-cd46-4123-803b-89440da06860") {
		log.Println("会計じゃないのはダメ!!")
	}
	// 会計でない人からのDMの場合無視
	// if(b.bot.API().GroupApi)
	// if there are quotation,ignore space
	sep := strings.FieldsFunc(p.Message.Text, func(r rune) bool {
		if r == '"' {
			quoted = !quoted
		}
		return !quoted && (r == ' ' || r == '\n')
	})

	if sep[0] != "/sendTo" {
		return
	}
	sendlist := strings.FieldsFunc(sep[1], func(r rune) bool { return r == ',' })
	// 全員が存在するIDでなければDM送信しない

	log.Println(sep)      // Foo, bar, random, "letters lol", stuff
	log.Println(sendlist) // Foo, bar, random, "letters lol", stuff
}

func (b *base) OnGroupExists(userID string, groupID string) bool {
	group, _, _ := b.bot.API().GroupApi.GetUserGroup(context.Background(), groupID).Execute()
	for _, m := range group.Members {
		if userID == m.Id {
			return true
		}
	}
	return false
}
