package main

import (
	"context"
	"log"
	"strings"
	"time"
	"yzk/initialize"

	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
)

//Config 定义了配置文件的结构
//type Config struct {
//	AppID uint64 `yml:"appid"` //机器人的appid
//	Token string `yml:"token"` //机器人的token
//}

var config initialize.Props
var api openapi.OpenAPI
var ctx context.Context

// 获取机器人的配置信息，即机器人的appid和token
func init() {
	props := initialize.Init()
	log.Println(props)
}

// 处理 @机器人 的消息
func atMessageEventHandler(event *dto.WSPayload, data *dto.WSATMessageData) error {
	if strings.HasSuffix(data.Content, "> hello") { // 如果@机器人并输入 hello 则回复 你好。
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
	//第二步：生成token，用于校验机器人的身份信息
	botToken := token.BotToken(uint64(config.AppID), config.Token)
	//第三步：获取操作机器人的API对象
	api = botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)
	//获取context
	ctx = context.Background()
	//第四步：获取websocket
	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatalln("websocket错误， err = ", err)
	}

	var atMessage event.ATMessageEventHandler = atMessageEventHandler

	// 注册socket消息处理
	intent := websocket.RegisterHandlers(atMessage)
	err = botgo.NewSessionManager().Start(ws, botToken, &intent)
	if err != nil {
		return
	} // 启动socket监听
}
