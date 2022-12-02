package genshin

import (
	"log"
)

type Api interface {
	Get(uid string) Account
}

func Get(uid string) Account {
	err := checkUID(uid)
	if err != nil {
		log.Println(err.Error())
	}
	return Account{}
}
