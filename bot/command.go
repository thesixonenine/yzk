package bot

import (
	"errors"
	"strings"
)

// var (
//
//	bind_uid = Command{Name: "绑定uid"}
//
// )
var Commands = []Command{
	{Name: "绑定uid"}, {Name: "查询"}, {Name: "体力"},
}

type CommandMethod interface {
	Equals(s string) bool
}

type Command struct {
	Name  string
	Param string
}

func (c Command) Equals(s string) bool {
	return strings.Contains(c.Name, s)
}

func (c Command) Execute() (string, error) {
	// 执行命令
	return "这是默认回复", nil
}

func FindCommand(s string) (*Command, error) {
	for _, command := range Commands {
		if command.Equals(s) {
			var com = &Command{Name: command.Name}
			p := strings.TrimPrefix(s, command.Name)
			com.Param = strings.TrimSpace(p)
			return com, nil
		}
	}
	return nil, errors.New("未查询到相关命令")
}
