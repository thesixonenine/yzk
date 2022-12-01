package config

import (
	"github.com/tencent-connect/botgo/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type absHandler interface {
	handle()
}
type Config struct {
	BotAppId int    `yaml:"appid"`
	BotToken string `yaml:"token"`
}

func (c Config) done() bool {
	if c.BotAppId == 0 || c.BotToken == "" {
		return false
	}
	return true
}

func (c Config) check() {
	if !c.done() {
		log.Error("程序配置初始化失败")
		os.Exit(1)
	}
}

type envHandler struct {
	config *Config
}

func (handler envHandler) handle() {
	appIdEnvStr := "bot.appId"
	tokenEnvStr := "bot.token"
	appId := strings.TrimSpace(os.Getenv(appIdEnvStr))
	if appId == "" {
		log.Debugf("环境变量{%s}中没有appId", appIdEnvStr)
	} else {
		handler.config.BotAppId, _ = strconv.Atoi(appId)
	}
	token := strings.TrimSpace(os.Getenv(tokenEnvStr))
	if token == "" {
		log.Debugf("环境变量{%s}中没有token", tokenEnvStr)
	} else {
		handler.config.BotToken = token
	}

}

type ymlHandler struct {
	config *Config
}

func (handler ymlHandler) handle() {
	fileName := "config.yml"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Debug("读取配置文件出错, " + err.Error())
		return
	}
	err = yaml.Unmarshal(content, handler.config)
	if err != nil {
		log.Debug("解析配置文件出错, ", err)
	}
}

func Init(config *Config) {
	p := []absHandler{envHandler{config}, ymlHandler{config}}
	log.Debug("开始初始化配置")
	for _, handler := range p {
		handler.handle()
		if config.done() {
			break
		}
	}
	log.Debug("结束初始化配置")
	config.check()
}
