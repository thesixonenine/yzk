package bot

import (
	"context"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/event"
	botLog "github.com/tencent-connect/botgo/log"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"log"
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
	userId := data.Author.ID
	content := message.ETLInput(data.Content)
	botLog.Debugf("接收到用户[%s]的@消息[%s]", userId, content)
	command, err := FindCommand(content)
	if err != nil {
		botLog.Warn(err.Error())
		return nil
	}
	content, err = command.Execute()
	if err != nil {
		botLog.Warn(err.Error())
		return nil
	}
	_, err = api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: content})
	if err != nil {
		botLog.Warn(err.Error())
		return nil
	}
	return nil
}

func QQBotStart() {
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
