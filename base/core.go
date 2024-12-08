package base

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/traPtitech/go-traq"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func (b *base) BotHandler() {
	log.Println(time.Now())
	b.bot.OnPing(b.Ping)
	b.bot.OnDirectMessageCreated(b.DMCreated)
}

func (b *base) DMCreated(p *payload.DirectMessageCreated) {

	// 会計でない人からのDMの場合無視
	if !b.OnGroupExists(p.Message.User.ID, "ba6552f8-cd46-4123-803b-89440da06860") {
		log.Println("会計じゃないのはダメ!!")
		b.BotDM(p.Message.User.ID, "BOT利用権限がありません")
		return
	}

	// ""を無視した空白くぎり
	quoted := false
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
	sendUUID := b.BotGetUsersUUID(sendlist)
	if sep[2] != "/message" {
		return
	}

	for _, u := range sendUUID {
		send := strings.Trim(sep[3], "\"")
		b.BotDM(u, send)
	}

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

func (b *base) BotGetUsersUUID(userNames []string) (useruuids []string) {
	res := []string{}
	users, httpres, err := b.bot.API().UserApi.GetUsers(context.Background()).Execute()
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Printf("Full HTTP response: %v\n", httpres)
	}
	// 探索高速化のための全ユーザーmap登録
	userNamemap := map[string]traq.User{}
	for _, u := range users {
		userNamemap[u.Name] = u
	}

	for _, s := range userNames {
		if val, ok := userNamemap[s]; !ok {
			log.Fatal("Not found such user")
			return []string{}
		} else {
			res = append(res, val.Id)
		}
	}
	log.Print(res)
	return res

}

func (b *base) BotDM(userid string, content string) {
	_, r, err := b.bot.API().
		MessageApi.
		PostDirectMessage(context.Background(), userid).
		PostMessageRequest(traq.PostMessageRequest{
			Content: content,
		}).Execute()
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Printf("Full HTTP response: %v\n", r)
	}
}
