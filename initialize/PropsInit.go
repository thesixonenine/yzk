package properties

import (
	"github.com/tencent-connect/botgo/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type absHandler interface {
	handle()
}
type Props struct {
	AppID int    `yml:"appid"`
	Token string `yml:"token"`
}

func (p Props) check() {
	if p.AppID == 0 || p.Token == "" {
		log.Error("程序配置初始化失败")
		os.Exit(1)
	}
}

type envHandler struct {
	props *Props
}

func (handler envHandler) handle() {
	appIdEnvStr := "bot.appId"
	tokenEnvStr := "bot.token"
	appId := strings.TrimSpace(os.Getenv(appIdEnvStr))
	if appId == "" {
		log.Debugf("环境变量{%s}中没有appId", appIdEnvStr)
	} else {
		handler.props.AppID, _ = strconv.Atoi(appId)
	}
	token := strings.TrimSpace(os.Getenv(tokenEnvStr))
	if token == "" {
		log.Debugf("环境变量{%s}中没有token", tokenEnvStr)
	} else {
		handler.props.Token = token
	}

}

type ymlHandler struct {
	props *Props
}

func (handler ymlHandler) handle() {
	fileName := "config.yml"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Debug("读取配置文件出错, " + err.Error())
		return
	}
	err = yaml.Unmarshal(content, handler.props)
	if err != nil {
		log.Debug("解析配置文件出错, ", err)
	}
}

func Init() Props {
	var prop Props
	p := []absHandler{envHandler{&prop}, ymlHandler{&prop}}
	log.Debug("开始初始化配置")
	for _, handler := range p {
		handler.handle()
	}
	log.Debug("结束初始化配置")
	prop.check()
	return prop
}
