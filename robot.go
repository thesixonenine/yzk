package main

import (
	"context"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"log"
	"strings"
	"time"
	config "yzk/initialize"
)

var botConfig config.Config
var api openapi.OpenAPI
var ctx context.Context

func init() {
	config.Init(&botConfig)
}

func atMessageEventHandler(event *dto.WSPayload, data *dto.WSATMessageData) error {
	if strings.HasSuffix(data.Content, "> hello") {
		_, err := api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "这里没有涩图可以看"})
		if err != nil {
			return nil
		}
	} else {
		_, err := api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "这是默认回复"})
		if err != nil {
			return nil
		}
	}
	return nil
}

func main() {
	botToken := token.BotToken(uint64(botConfig.BotAppId), botConfig.BotToken)
	api = botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)
	ctx = context.Background()
	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatalln("websocket错误， err = ", err)
	}

	var atMessage event.ATMessageEventHandler = atMessageEventHandler

	intent := websocket.RegisterHandlers(atMessage)
	err = botgo.NewSessionManager().Start(ws, botToken, &intent)
	if err != nil {
		return
	}
}
