package base

import (
	"context"
	"log"

	"github.com/traPtitech/go-traq"
	"github.com/traPtitech/traq-ws-bot/payload"
)

// Ping
func (b *base) Ping(p *payload.Ping) {
	log.Println("Pong::", p.EventTime.String())
}

// botの入退場を管理
func (b *base) BotJoiner(channelID string) {
	_, err := b.bot.API().BotApi.LetBotJoinChannel(context.Background(), b.botid).
		PostBotActionJoinRequest(*traq.NewPostBotActionJoinRequest(channelID)).Execute()
	if err != nil {
		log.Println(err)
	}
	channel, _, _ := b.bot.API().ChannelApi.GetChannel(context.Background(), channelID).Execute()
	log.Println("joined:" + channel.Name)
}

func (b *base) BotLeaver(channelID string) {
	_, err := b.bot.API().BotApi.LetBotLeaveChannel(context.Background(), b.botid).
		PostBotActionLeaveRequest(*traq.NewPostBotActionLeaveRequest(channelID)).Execute()
	if err != nil {
		log.Println(err)
	}
	channel, _, _ := b.bot.API().ChannelApi.GetChannel(context.Background(), channelID).Execute()
	log.Println("left:" + channel.Name)
}
